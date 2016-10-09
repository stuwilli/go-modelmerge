package modelmerge

import (
	"fmt"
	"reflect"
)

func sliceContains(ss []string, val string) bool {
	fmt.Println(val)
	for _, s := range ss {

		if s == val {
			return true
		}
	}
	return false
}

//Merge ...
func Merge(original interface{}, updates interface{}, fields []string) {

	oVal := reflect.ValueOf(original)

	if oVal.Kind() == reflect.Ptr {
		oVal = oVal.Elem()
	}

	uVal := reflect.ValueOf(updates)

	if uVal.Kind() == reflect.Ptr {
		uVal = uVal.Elem()
	}

	for i := 0; i < oVal.NumField(); i++ {

		if sliceContains(fields, oVal.Type().Field(i).Name) {

			oVal.Field(i).Set(uVal.Field(i))
		}

	}
}
