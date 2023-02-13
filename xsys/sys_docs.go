package xsys

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type SysDocs struct {
}

func (o *SysDocs) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	docs := strings.Split(ctx.FormValue("id"), ",")
	if len(docs) < 1 {
		return make([]interface{}, 0), nil
	}

	query := fmt.Sprintf("SELECT id, name_, size_ FROM sys_doc WHERE id IN (?%s)", strings.Repeat(", ?", len(docs)-1))
	args := make([]interface{}, 0, len(docs))
	for _, doc := range docs {
		args = append(args, doc)
	}

	return asql.Select(tx, query, args...)
}

// PostUpload 文档上传
func (o *SysDocs) PostUpload(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	if err := ctx.ParseMultipartForm(1 << 30); err != nil {
		return nil, err
	}

	logrus.Debugf("upload_fullpath is %q", ctx.PostFormValue("upload_fullpath"))

	rFile, head, err := ctx.FormFile("upload")
	if err != nil {
		return nil, err
	}

	defer rFile.Close()

	mime := head.Header.Get("Content-Type")
	size := head.Size
	name := head.Filename
	docId := asql.GenerateId()

	// 随机生成2个16进制目录
	k4 := make([]byte, 2)
	if _, err := io.ReadFull(base.Config.Rand(), k4); err != nil {
		return nil, err
	}

	// 创建目录
	p1, p2 := hex.EncodeToString(k4[:1]), hex.EncodeToString(k4[1:])
	dir := filepath.Join("store", "upload", time.Now().Format("0601"), p1, p2)
	if _, err := os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}

		if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
			return nil, err
		}
	}

	// 创建文件
	wFile, err := os.Create(filepath.Join(dir, docId))
	if err != nil {
		return nil, err
	}
	defer wFile.Close()

	// 文件拷贝
	if _, err := io.Copy(wFile, rFile); err != nil {
		return nil, err
	}

	// 记录文件信息
	query := `
		INSERT INTO sys_doc(
			id, name_, size_, mime_, dir_, order_, create_at_, 
			user_id_, user_code_, user_name_, depart_id_, depart_code_, depart_name_
		)
		VALUES (?,?,?,?,?,?,?, ?,?,?,?,?,?)
	`
	args := []interface{}{
		docId, name, size, mime, dir, asql.GenerateOrderId(), asql.GetNow(),
		ctx.GetUserId(), ctx.GetUserCode(), ctx.GetUserName(), ctx.GetDepartId(), ctx.GetDepartCode(), ctx.GetDepartName(),
	}
	if err := asql.Insert(tx, query, args...); err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "server", "id": docId}, nil
}

func (o *SysDocs) PostDownload(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	docId := ctx.PostFormValue("id")

	var dir string

	if err := asql.SelectRow(tx, "SELECT dir_ FROM sys_doc WHERE id = ?", docId).Scan(&dir); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("无效的文档ID %q", docId)
		}

		return nil, err
	}

	// 获取文件
	rFile, err := os.Open(filepath.Join(base.Config.Dir(), dir, docId))
	if err != nil {
		return nil, err
	}
	defer rFile.Close()

	// 写入文件
	if _, err := io.Copy(ctx.Writer, rFile); err != nil {
		return nil, err
	}

	return map[string]string{"status": "success"}, nil
}
