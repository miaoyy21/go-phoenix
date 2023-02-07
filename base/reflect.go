package base

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"time"
)

func ReflectToStruct(source map[string]string, target interface{}) error {
	ele := reflect.TypeOf(target).Elem()

	for i := 0; i < ele.NumField(); i++ {
		tag := ele.Field(i).Tag
		name, ok := tag.Lookup("name")
		if !ok {
			continue
		}

		value := reflect.ValueOf(target).Elem().Field(i)
		if err := setFieldValue(source[name], value); err != nil {
			return err
		}
	}

	return nil
}

func setFieldValue(src string, value reflect.Value) error {
	kind := value.Kind()
	switch kind {
	case reflect.String:
		value.SetString(src)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		num, err := strconv.ParseInt(src, 10, 64)
		if err != nil {
			return err
		}
		value.SetInt(num)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		num, err := strconv.ParseUint(src, 10, 64)
		if err != nil {
			return err
		}
		value.SetUint(num)
	case reflect.Float32, reflect.Float64:
		num, err := strconv.ParseFloat(src, 64)
		if err != nil {
			return err
		}
		value.SetFloat(num)
	case reflect.Struct:
		_, ok := value.Interface().(time.Time)
		if !ok {
			logrus.Errorf("UNDEFINED TYPE %#v", kind.String())
			return nil
		}

		dt, err := time.Parse("2006-01-02 15:04:05", src)
		if err != nil {
			return err
		}

		value.Set(reflect.ValueOf(dt))
	default:
		return fmt.Errorf("UnRecogniza Field Type %#v", kind.String())
	}

	return nil
}
