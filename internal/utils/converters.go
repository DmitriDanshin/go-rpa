package utils

import "reflect"

func StructToMap(node interface{}) map[string]any {
	result := make(map[string]any)
	v := reflect.ValueOf(node)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Tag.Get("json")
		if fieldName == "" {
			fieldName = t.Field(i).Name
		}
		result[fieldName] = field.Interface()
	}
	return result
}
