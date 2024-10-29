package pub

import (
	"reflect"
)

func IsEmpty(obj interface{}) bool {
	if obj == nil {
		return true
	}

	objVal := reflect.ValueOf(obj)
	switch objVal.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return objVal.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return objVal.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return objVal.Float() == 0
	case reflect.Bool:
		return !objVal.Bool()
	case reflect.String:
		return objVal.String() == ""
	case reflect.Slice, reflect.Array, reflect.Map:
		return objVal.Len() == 0
	case reflect.Ptr:
		return objVal.IsNil()
	}

	return false
}

func StructToMap(obj interface{}) map[string]interface{} {
	objVal := reflect.ValueOf(obj).Elem()
	objType := objVal.Type()

	m := make(map[string]interface{})
	for i := 0; i < objVal.NumField(); i++ {
		field := objType.Field(i)
		val := objVal.Field(i).Interface()
		if !IsEmpty(val) {
			m[field.Name] = val
		}
	}

	return m
}
