package base

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GenerateToken(userId, userCode, userName, departId, departCode, departName, password, userAgent string, iExpire int) string {

	src := make([]string, 0, 3)

	// 0 用户ID
	src = append(src, userId)

	// 1 附加信息
	ext := fmt.Sprintf("%s_%s_%s_%s_%s_%s_%s", userAgent, password, userCode, userName, departId, departCode, departName)
	md5Ext := md5.Sum([]byte(ext))
	src = append(src, base64.StdEncoding.EncodeToString(md5Ext[:]))

	// 2 失效时间
	expire := strconv.FormatInt(time.Now().Add(time.Duration(iExpire)*time.Second).Unix(), 10)
	src = append(src, expire)

	bytes := Config.AesStream([]byte(strings.Join(src, ",")))

	return base64.StdEncoding.EncodeToString(bytes)
}
