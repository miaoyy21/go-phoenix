package jsvm

import (
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type Utils struct {
}

func NewUtils() *Utils {
	return &Utils{}
}

func (u *Utils) ToInt(value interface{}) int64 {
	if value == nil {
		return 0
	}

	switch value.(type) {
	case int, int8, int16, int32, int64:
	case uint, uint8, uint16, uint32, uint64:
	case float32, float64:
		return value.(int64)
	case string:
		newValue := strings.TrimSpace(value.(string))
		if len(newValue) < 1 {
			return 0
		}

		num, err := strconv.ParseInt(newValue, 10, 64)
		if err != nil {
			logrus.Panicf("%q 无法转换为int64类型：%s", value, err.Error())
		}

		return num
	default:
		logrus.Panicf("%#v 无法转换为int64类型", value)
	}

	return 0
}

func (u *Utils) ToFloat(value interface{}) float64 {
	if value == nil {
		return 0
	}

	switch value.(type) {
	case int, int8, int16, int32, int64:
	case uint, uint8, uint16, uint32, uint64:
	case float32, float64:
		return value.(float64)
	case string:
		newValue := strings.TrimSpace(value.(string))
		if len(newValue) < 1 {
			return 0
		}

		num, err := strconv.ParseFloat(newValue, 64)
		if err != nil {
			logrus.Panicf("%q 无法转换为float64类型：%s", newValue, err.Error())
		}

		return num
	default:
		logrus.Panicf("%#v 无法转换为float64类型", value)
	}

	return 0
}
