package xsys

import (
	"go-phoenix/base"
	"log"
	"testing"
)

func TestSysUsers_GetPassword(t *testing.T) {
	if err := base.InitConfig("/Users/miaojingyi/Documents/dev/go/src/go-phoenix"); err != nil {
		log.Fatal(err.Error())
	}

	src, err := base.Config.AesDecodeString("ccf96c9be3fbccb498982b63")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(string(src))
}
