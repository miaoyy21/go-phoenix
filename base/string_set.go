package base

type StringSet struct {
	values []string
	exists map[string]struct{}
}

func NewStringSet(slice []string) *StringSet {
	exists := make(map[string]struct{})

	values := make([]string, 0, len(slice))
	for _, value := range slice {
		if _, ok := exists[value]; ok {
			continue
		}

		exists[value] = struct{}{}
		values = append(values, value)
	}

	return &StringSet{
		values: values,
		exists: exists,
	}
}

func (set *StringSet) Append(value string) {
	if _, ok := set.exists[value]; ok {
		return
	}

	set.exists[value] = struct{}{}
	set.values = append(set.values, value)
}

func (set *StringSet) Remove(value string) {
	if _, ok := set.exists[value]; !ok {
		return
	}

	delete(set.exists, value)

	newValues := make([]string, 0, len(set.values)-1)
	for _, old := range set.values {
		if old == value {
			continue
		}

		newValues = append(newValues, old)
	}

	set.values = newValues
}

func (set *StringSet) Reset() {
	set.values = make([]string, 0)
	set.exists = make(map[string]struct{})
}

func (set *StringSet) Values() []string {
	return set.values
}
