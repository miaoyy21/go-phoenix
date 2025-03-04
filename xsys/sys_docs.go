package xsys

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
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

func (o *SysDocs) GetDownload(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	docId := ctx.FormValue("id")

	return o.get(tx, ctx, docId)
}

func (o *SysDocs) GetSigner(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	user := ctx.FormValue("user")

	var docId string
	if err := asql.SelectRow(tx, "SELECT signer_ FROM sys_user WHERE id = ?", user).Scan(&docId); err != nil {
		return nil, err
	}

	return o.get(tx, ctx, docId)
}

func (o *SysDocs) get(tx *sql.Tx, ctx *handle.Context, docId string) (interface{}, error) {

	var dir, mime string

	if err := asql.SelectRow(tx, "SELECT dir_, mime_ FROM sys_doc WHERE id = ?", docId).Scan(&dir, &mime); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("无效的文档ID %q", docId)
		}

		return nil, err
	}

	// 根据操作系统类型，进行文件路径分割符进行替换
	if strings.EqualFold(runtime.GOOS, "windows") {
		dir = strings.ReplaceAll(dir, "/", "\\")
	} else {
		dir = strings.ReplaceAll(dir, "\\", "/")
	}

	// 获取文件
	rFile, err := os.Open(filepath.Join(base.Config.Dir(), dir, docId))
	if err != nil {
		return nil, err
	}
	defer rFile.Close()

	// 写入文件
	ctx.Writer.Header().Set("Content-Type", mime)
	if _, err := io.Copy(ctx.Writer, rFile); err != nil {
		return nil, err
	}

	return map[string]string{"status": "success"}, nil
}

// PostImport Excel导入
func (o *SysDocs) PostImport(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	// 保存Excel导入文件的目录
	dir := filepath.Join("store", "import", time.Now().Format("0601"))

	return o.save(tx, ctx, dir)
}

// PostUpload 文档上传
func (o *SysDocs) PostUpload(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	// 随机生成2个16进制目录
	k4 := make([]byte, 2)
	if _, err := io.ReadFull(base.Config.Rand(), k4); err != nil {
		return nil, err
	}

	// 上传文档目录
	p1, p2 := hex.EncodeToString(k4[:1]), hex.EncodeToString(k4[1:])
	dir := filepath.Join("store", "upload", time.Now().Format("0601"), p1, p2)

	return o.save(tx, ctx, dir)
}

// PostSigner 数字签名
func (o *SysDocs) PostSigner(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	// 上传签名目录
	dir := filepath.Join("store", "signer")

	return o.save(tx, ctx, dir)
}

// PostUpload 文档上传
func (o *SysDocs) save(tx *sql.Tx, ctx *handle.Context, dir string) (interface{}, error) {
	if err := ctx.ParseMultipartForm(1 << 30); err != nil {
		return nil, err
	}

	rFile, head, err := ctx.FormFile("upload")
	if err != nil {
		return nil, err
	}

	defer rFile.Close()

	mime := head.Header.Get("Content-Type")
	size := head.Size
	name := head.Filename

	if _, err := os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}

		if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
			return nil, err
		}
	}

	// 创建文件
	docId := asql.GenerateId()
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
		ctx.UserId(), ctx.UserCode(), ctx.UserName(), ctx.DepartId(), ctx.DepartCode(), ctx.DepartName(),
	}
	if err := asql.Insert(tx, query, args...); err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "server", "id": docId, "value": docId}, nil
}
