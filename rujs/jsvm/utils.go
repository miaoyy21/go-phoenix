package jsvm

import (
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type Utils struct {
}

func NewUtils() *Utils {
	return &Utils{}
}

func (u *Utils) NewDecimal(value interface{}) decimal.Decimal {

	switch value.(type) {
	case nil:
		return decimal.NewFromInt(0)
	case int, int8, int16, int32, int64:
		return decimal.NewFromInt(value.(int64))
	case uint, uint8, uint16, uint32, uint64:
		return decimal.NewFromUint64(value.(uint64))
	case float32, float64:
		return decimal.NewFromFloat(value.(float64))
	case string:
		dec, err := decimal.NewFromString(value.(string))
		if err != nil {
			logrus.Panicf("%#v 无法转换为decimal.Decimal类型：%s", value, err.Error())
		}

		return dec
	default:
		logrus.Panicf("%#v 无法转换为decimal.Decimal类型", value)
	}

	return decimal.NewFromInt(0)
}
