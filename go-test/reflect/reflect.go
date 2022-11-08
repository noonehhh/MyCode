package main
/*

*/

type TestSt struct {
	age string  
	name string
}

func main () {
	ts := &TestSt{}
	tsType := reflect.TypeOf(*ts)
	tsValue := reflect.ValueOf(*ts)

	for i := 0; i < tsType.NumField(); i++ {
		field := tsType..Field(i)
		if filed.Tag.Get("json") == "age" {
			res = cast.ToString(tsValue.FieldByName(filed.Name))
			break
		}
	}
}