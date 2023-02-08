package base

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type IntSet struct {
	values []int
	exists map[int]struct{}
}

func NewIntSet(slice []int) *IntSet {
	exists := make(map[int]struct{})

	values := make([]int, 0, len(slice))
	for _, value := range slice {
		if _, ok := exists[value]; ok {
			continue
		}

		exists[value] = struct{}{}
		values = append(values, value)
	}

	return &IntSet{
		values: values,
		exists: exists,
	}
}

func NewIntSetFromString(src string) *IntSet {
	slice := make([]int, 0)
	if len(src) > 2 {
		if err := json.Unmarshal([]byte(src), &slice); err != nil {
			logrus.Panicf("NewIntSetFromString(%s) Unmarshal() failure %s", src, err.Error())
		}
	}

	return NewIntSet(slice)
}

func (set *IntSet) Append(value int) {
	if _, ok := set.exists[value]; ok {
		return
	}

	set.exists[value] = struct{}{}
	set.values = append(set.values, value)
}

func (set *IntSet) Remove(value int) {
	if _, ok := set.exists[value]; !ok {
		return
	}

	delete(set.exists, value)

	newValues := make([]int, 0, len(set.values)-1)
	for _, old := range set.values {
		if old == value {
			continue
		}

		newValues = append(newValues, old)
	}

	set.values = newValues
}

func (set *IntSet) Reset() {
	set.values = make([]int, 0)
	set.exists = make(map[int]struct{})
}

func (set *IntSet) String() string {
	bs, err := json.Marshal(set.values)
	if err != nil {
		logrus.Panicf("IntSet() String() Marshal() failure %s", err.Error())
	}

	return string(bs)
}
