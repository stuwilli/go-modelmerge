package modelmerge

import (
	"fmt"
	"reflect"
)

//Merge ...
func Merge(original interface{}, updates interface{}) {

	oVal := reflect.ValueOf(original)

	if oVal.Kind() == reflect.Ptr {
		oVal = oVal.Elem()
	}

	uVal := reflect.ValueOf(updates)

	if uVal.Kind() == reflect.Ptr {
		uVal = uVal.Elem()
	}

	typ := oVal.Type()

	for i := 0; i < oVal.NumField(); i++ {
		fmt.Println(typ.Field(i).Name)
	}
}
