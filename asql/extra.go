package asql

import (
	"bytes"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

var bs32 = base32.NewEncoding("123456789abcdfghjkmnopqrstuvwxyz").WithPadding('0')

func GenerateId() string {
	buf := &bytes.Buffer{}

	// 为便于排序等，增加时间戳前缀
	var prefix [4]byte
	binary.BigEndian.PutUint32(prefix[:], uint32(time.Now().Unix()))
	buf.Write(prefix[:])

	// UUID
	id := uuid.New()
	buf.Write(id[:])

	// 编码为base
	return bs32.EncodeToString(buf.Bytes())
}

func GenerateOrderId() int64 {
	return time.Now().UnixNano()
}

func GetNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Arguments(args ...interface{}) string {
	ss := make([]string, 0, len(args))

	for _, arg := range args {
		switch arg.(type) {
		case string:
			ss = append(ss, fmt.Sprintf("%q", arg))
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			ss = append(ss, fmt.Sprintf("%d", arg))
		case bool:
			ss = append(ss, fmt.Sprintf("%t", arg))
		case float32, float64:
			ss = append(ss, fmt.Sprintf("%f", arg))
		default:
			ss = append(ss, fmt.Sprintf("%#v<%T>", arg, arg))
		}
	}

	return strings.Join(ss, " , ")
}
