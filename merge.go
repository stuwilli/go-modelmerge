package merge

import (
	"errors"
	"fmt"
	"reflect"
)

func inFields(fields []string, fieldName string) bool {

	for _, field := range fields {

		if field == fieldName {
			return true
		}
	}
	return false
}

func setValue(i int, fields []string, oVal reflect.Value, uVal reflect.Value) error {

	switch oVal.Field(i).Kind() {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		oVal.Field(i).SetInt(uVal.Field(i).Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:

		oVal.Field(i).SetUint(uVal.Field(i).Uint())

	case reflect.Float32, reflect.Float64:

		oVal.Field(i).SetFloat(uVal.Field(i).Float())

	case reflect.String:

		oVal.Field(i).SetString(uVal.Field(i).String())

	case reflect.Bool:

		oVal.Field(i).SetBool(uVal.Field(i).Bool())

	case reflect.Slice:

		return errors.New("Slice support has not been implimented")

	case reflect.Struct:

		SelectedFields(oVal.Field(i).Interface(), uVal.Field(i).Interface(), fields)

	case reflect.Map:

		return errors.New("Slice support has not been implimented")

	default:

		return errors.New("Unkown type")
	}

	return nil
}

//SelectedFields ...
func SelectedFields(original interface{}, updates interface{}, fields []string) error {

	oVal := reflect.ValueOf(original)

	if oVal.Kind() == reflect.Ptr {
		oVal = oVal.Elem()
	}

	uVal := reflect.ValueOf(updates)

	if uVal.Kind() == reflect.Ptr {
		uVal = uVal.Elem()
	}

	for i := 0; i < oVal.NumField(); i++ {

		if !inFields(fields, oVal.Type().Field(i).Name) {
			fmt.Println("Skipping " + oVal.Type().Field(i).Name)
			continue
		}

		if oVal.Field(i).CanSet() {

			if err := setValue(i, fields, oVal, uVal); err != nil {
				return err
			}

		}

	}

	return nil
}
