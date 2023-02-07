package base

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func ResAsMapSlice(res []map[string]string, sensitive bool, k string, v string) (map[string]string, []string) {
	ms, ss := make(map[string]string), make([]string, 0, len(res))

	for _, row := range res {
		key, ok := row[k]
		if !ok {
			logrus.Panicf("arguments ResAsMapSlice's Map Must has key field %q .", k)
		}

		value, ok := row[v]
		if !ok {
			logrus.Panicf("arguments ResAsMapSlice's Map Must has Value field %q .", v)
		}

		if sensitive {
			ms[key] = value
			ss = append(ss, key)
		} else {
			ms[strings.ToLower(key)] = strings.ToLower(value)
			ss = append(ss, strings.ToLower(key))
		}
	}

	return ms, ss
}

func ResAsMap2(res []map[string]string, k string) map[string]map[string]string {
	ms := make(map[string]map[string]string)

	for _, row := range res {
		key, ok := row[k]
		if !ok {
			logrus.Panicf("arguments ResAsMap2's Map Must has field %q .", k)
		}

		ms[key] = row
	}

	return ms
}

func ResAsSliceString(res []map[string]string, k string) []string {
	ss := make([]string, 0, len(res))

	for _, row := range res {
		value, ok := row[k]
		if !ok {
			logrus.Panicf("arguments ResAsSliceString's Map Must has field %q .", k)
		}

		ss = append(ss, value)
	}

	return ss
}

func ResAsMapInt(res []map[string]string, k string, v string) map[int]string {
	ms := make(map[int]string)

	for _, row := range res {
		value, ok := row[k]
		if !ok {
			logrus.Panicf("arguments ResAsMapInt's Map Must has field %q .", k)
		}

		// 如果数据库定义为INT类型，不存在转换失败的问题
		i64, err := strconv.Atoi(value)
		if err != nil {
			logrus.Panicf("value of %q convert to int failure %s.", value, err.Error())
		}

		ms[i64] = row[v]
	}

	return ms
}

func SliceStringAsSet(res []string) []string {
	ss := make([]string, 0, len(res))
	ms := make(map[string]struct{})

	for _, val := range res {
		if _, ok := ms[val]; ok {
			continue
		}

		ss = append(ss, val)
		ms[val] = struct{}{}
	}

	return ss
}

func SliceIntAsSet(res []int) []int {
	ss := make([]int, 0, len(res))
	ms := make(map[int]struct{})

	for _, val := range res {
		if _, ok := ms[val]; ok {
			continue
		}

		ss = append(ss, val)
		ms[val] = struct{}{}
	}

	return ss
}

func StringSliceInt(olds string, values []int) string {
	oldValues := make([]int, 0)

	if len(olds) > 2 {
		if err := json.Unmarshal([]byte(olds), &oldValues); err != nil {
			logrus.Panicf("arguments SliceAppendInt's Unmarshal failure %s", err.Error())
		}
	}

	newValues := append(oldValues, values...)
	news, err := json.Marshal(SliceIntAsSet(newValues))
	if err != nil {
		logrus.Panicf("arguments SliceAppendInt's Unmarshal failure %s", err.Error())
	}

	return string(news)
}
