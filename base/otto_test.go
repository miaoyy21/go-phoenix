package base

import (
	"github.com/robertkrimen/otto"
	"testing"
)

func TestOtto(t *testing.T) {
	vm := otto.New()
	if err := vm.Set("$values", map[string]int{
		"abc": 123,
		"def": 456,
	}); err != nil {
		t.Fatalf("jsvm Set Failure :: %s", err.Error())
	}
	value, err := vm.Run(`
		//var onStart = function(){
		//	// TODO Something
		//	console.log($values["def"]);
		//	return ["111","222","333"];
		//};
		//
		//onStart();

["111","222","333"]
	`)
	if err != nil {
		t.Fatalf("jsvm Run Failure :: %s", err.Error())
	}

	t.Logf("Out value is %#v", value.String())

	val, err := value.Export()
	if err != nil {
		t.Fatalf("jsvm Export Failure :: %s", err.Error())
	}

	t.Logf("Out export is %t, %#v", value.IsBoolean(), val)

	ss, ok := val.([]string)
	t.Logf("Out Convert is %t, %#v", ok, ss)

	t.Logf("%t", false)

	//out,err:=value.ToInteger()
	//if err!=nil{
	//	t.Fatalf("value ToInteger Failure :: %s",err.Error())
	//}
	//
	//t.Logf("Out is %d",out)

	//obj := value.Object()
	//out, err := obj.Get("s")
	//if err != nil {
	//	t.Fatalf("value ToInteger Failure :: %s", err.Error())
	//}

	//val,err:=obj.Value().Export()
	//if err != nil {
	//	t.Fatalf("jsvm Export Failure :: %s", err.Error())
	//}
	//
	//t.Logf("Out is %#v", val)

	//out, err := value.Call(otto.NullValue())
	//if err != nil {
	//	t.Fatalf("value Call Failure :: %s", err.Error())
	//}
	//
	//t.Logf("Out is %#v", out)
}
