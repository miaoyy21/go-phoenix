package base

import (
	"net/url"
	"strings"
)

func CompareMap(latest map[string]string, present map[string]string) (map[string]string, map[string]string, map[string]string) {
	added, changed, removed := make(map[string]string), make(map[string]string), make(map[string]string)

	// 增加
	for key, value := range present {
		xvalue, ok := latest[key]
		if !ok {
			added[key] = value
		} else if xvalue != value {
			// 更新
			changed[key] = value
		}
	}

	// 移除
	for key, value := range latest {
		if _, ok := present[key]; !ok {
			removed[key] = value
		}
	}

	return added, changed, removed
}

func GetURLValues(values url.Values) map[string]string {
	ms := make(map[string]string)
	if len(values) < 1 {
		return make(map[string]string)
	}

	for key, value := range values {
		if len(value) > 0 {
			ms[key] = value[len(value)-1]
		}
	}

	return ms
}

func CompareMapChanged(latest map[string]string, present map[string]string) map[string]string {
	changed := make(map[string]string)

	// 增加
	for key, value := range present {
		if v, ok := latest[key]; !ok || !strings.EqualFold(value, v) {
			changed[key] = value
		}
	}

	return changed
}
