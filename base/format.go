package base

import (
	"bytes"
	"fmt"
)

type MapString map[string]string

func (m MapString) String() string {
	if len(m) == 0 {
		return "{}"
	}

	buf := new(bytes.Buffer)

	buf.WriteString("{")
	for k, v := range m {
		if buf.Len() <= 1 {
			buf.WriteString(fmt.Sprintf("%q: %q", k, v))
			continue
		}

		buf.WriteString(fmt.Sprintf(", %q: %q", k, v))
	}
	buf.WriteString("}")

	return buf.String()
}
