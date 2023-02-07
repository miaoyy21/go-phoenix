package flow

import (
	"encoding/json"
	"github.com/robertkrimen/otto"
)

func flowReg(node Flowable, data string) func(vm *otto.Otto) error {
	return func(vm *otto.Otto) error {
		// Node
		if err := vm.Set("$node", node); err != nil {
			return err
		}

		// Convert Data as map<string>interface{}
		values := make(map[string]interface{})
		if err := json.Unmarshal([]byte(data), &values); err != nil {
			return err
		}

		// Values
		if err := vm.Set("$values", values); err != nil {
			return err
		}

		return nil
	}
}
