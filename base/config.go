package base

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

var Config config

type config struct {
	Level  string `json:"level"`
	AesKey string `json:"aesKey"`
	AesIv  string `json:"aesIv"`

	Host string `json:"host"`
	Port string `json:"port"`

	DBDriver     string `json:"dbDriver"`
	DBDataSource string `json:"dbDataSource"`

	dir string

	aesIv    []byte
	aesBlock cipher.Block

	rand *rand.Rand
}

func (cfg config) Dir() string {
	return cfg.dir
}

func (cfg config) Rand() *rand.Rand {
	return cfg.rand
}

func (cfg config) AesStream(src []byte) []byte {
	aesStream := cipher.NewCTR(cfg.aesBlock, cfg.aesIv)

	dst := make([]byte, len(src))
	aesStream.XORKeyStream(dst, src)
	return dst
}

func (cfg config) AesEncodeString(src string) string {
	return hex.EncodeToString(cfg.AesStream([]byte(src)))
}

func (cfg config) AesDecodeString(dst string) (src string, err error) {
	eSrc, err := hex.DecodeString(dst)
	if err != nil {
		return src, err
	}

	return string(Config.AesStream(eSrc)), nil
}

func InitConfig(dir string) error {

	// 获取配置文件
	bytes, err := os.ReadFile(filepath.Join(dir, "config.json"))
	if err != nil {
		return err
	}

	// 解析
	if err := json.Unmarshal(bytes, &Config); err != nil {
		return err
	}

	// 设置日志级别
	level, err := logrus.ParseLevel(Config.Level)
	if err != nil {
		return err
	}

	Config.dir = dir

	// AES key
	aesKey, err := hex.DecodeString(Config.AesKey)
	if err != nil {
		return err
	}

	// AES cipher
	aesBlock, err := aes.NewCipher(aesKey)
	if err != nil {
		return err
	}

	// AES iv
	aesIv, err := hex.DecodeString(Config.AesIv)
	if err != nil {
		return err
	}
	Config.aesIv = aesIv

	Config.aesBlock = aesBlock

	logrus.SetLevel(level)
	logrus.Infof("Aes key is %q", Config.AesKey)
	logrus.Infof("Aes Iv is %q", Config.AesIv)
	logrus.Infof("Work Dir is %q", Config.dir)
	logrus.Infof("Log Level is %q", Config.Level)
	logrus.Infof("DB Driver is %q", Config.DBDriver)

	seed := time.Now().UnixNano()
	Config.rand = rand.New(rand.NewSource(seed))

	return nil
}
