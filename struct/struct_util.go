package _struct

import "reflect"

func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	value := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		key := typ.Field(i).Name
		result[key] = field.Interface()
	}

	return result
}
