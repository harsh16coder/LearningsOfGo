package reflection

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		fieldval := val.Field(i)
		switch fieldval.Kind() {
		case reflect.String:
			fn(fieldval.String())
		case reflect.Struct:
			walk(fieldval.Interface(), fn)
		}
	}
}
