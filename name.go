package merge

import "reflect"

//DeRef ...
func DeRef(x interface{}) reflect.Value {

	z := reflect.ValueOf(x)

	if z.Kind() == reflect.Ptr {
		z = z.Elem()
	}

	return z
}

//FieldsByName ...
func FieldsByName(dest interface{}, src interface{}) error {

	destVal := DeRef(dest)
	srcVal := DeRef(src)

	for i := 0; i < srcVal.NumField(); i++ {

		name := srcVal.Type().Field(i).Name
		kind := srcVal.Field(i).Kind()

		val := destVal.FieldByName(name)

		if val.Kind() != kind {
			continue
		}

		val.Set(srcVal.Field(i))

	}
	return nil
}
