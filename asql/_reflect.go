package asql


//type Student struct {
//	ID    string    `name:"id" required:"true"`
//	Name  string    `name:"name" required:"true"`
//	Age   int       `name:"age"`
//	Score float64   `name:"score"`
//	Birth time.Time `name:"birth"`
//
//	None string
//}
//
//func (s *Student) Execute(id string, name string) (int, error) {
//	logrus.Infof("EXECUTE >>>>>>")
//
//	return 1234321, errors.New("TEST")
//}
//
//func tstReflect1() {
//	form := map[string]string{
//		"id":    "ssss",
//		"name":  "dddd",
//		"age":   "23",
//		"score": "12.23",
//		"birth": "2022-01-03 14:23:11",
//	}
//
//	student := &Student{}
//
//	ele := reflect.TypeOf(student).Elem()
//	for i := 0; i < ele.NumField(); i++ {
//		//logrus.Debugf("%#v", ele.Field(i))
//
//		tag := ele.Field(i).Tag
//		logrus.Debugf("name:%s required:%t", tag.Get("name"), tag.Get("required") == "true")
//
//		name, ok := tag.Lookup("name")
//		if !ok {
//			logrus.Warnf("Miss StructTag Field %s", ele.Field(i).Name)
//			continue
//		}
//
//		required := tag.Get("required") == "true"
//
//		src, ok1 := form[name]
//		if !ok1 && required {
//			logrus.Errorf("Missing Field %s", tag.Get("name"))
//			continue
//		}
//
//		if err := setFieldValue(src, reflect.ValueOf(student).Elem().Field(i)); err != nil {
//			logrus.Errorf("Set Field Value Failure %s", err.Error())
//			continue
//		}
//	}
//
//	logrus.Infof("Parse Value After %#v", student)
//}
//
//func tstReflect2() {
//
//	student := &Student{}
//
//	typ := reflect.TypeOf(student)
//	for i := 0; i < typ.NumMethod(); i++ {
//		mtyp := typ.Method(i).Type
//
//		logrus.Debugf("Function Num In %d", mtyp.NumIn())
//		logrus.Debugf("Function Num Out %d", mtyp.NumOut())
//
//		in := make([]reflect.Value, 0)
//		in = append(in, reflect.New(mtyp).Elem())
//		for p := 1; p < mtyp.NumIn(); p++ {
//			logrus.Debugf("Argument Value %s", mtyp.In(p))
//			inValue := reflect.New(mtyp.In(p))
//
//			inValue.Elem().SetString("1111")
//			logrus.Debugf("Create Argument Value %#v", inValue)
//
//			in = append(in, reflect.ValueOf(inValue))
//		}
//
//		logrus.Debugf("Create In Values %#v", in[1].Elem().Interface())
//		logrus.Debugf("Create Function Num In %d", len(in))
//		logrus.Debugf("Create Function Num In %#v", reflect.ValueOf(student).Method(i))
//		reflect.ValueOf(student).Method(i).Call(in)
//	}
//}
//
//func setFieldValue(src string, value reflect.Value) error {
//	kind := value.Kind()
//	switch kind {
//	case reflect.String:
//		value.SetString(src)
//	case reflect.Int:
//		num, err := strconv.ParseInt(src, 10, 64)
//		if err != nil {
//			return err
//		}
//		value.SetInt(num)
//	case reflect.Float64:
//		num, err := strconv.ParseFloat(src, 64)
//		if err != nil {
//			return err
//		}
//		value.SetFloat(num)
//	case reflect.Struct:
//		_, ok := value.Interface().(time.Time)
//		if !ok {
//			logrus.Errorf("UNDEFINED TYPE %#v", kind.String())
//			return nil
//		}
//
//		dt, err := time.Parse("2006-01-02 15:04:05", src)
//		if err != nil {
//			return err
//		}
//
//		value.Set(reflect.ValueOf(dt))
//	default:
//		logrus.Errorf("UNDEFINED TYPE %#v", kind.String())
//	}
//
//	return nil
//}