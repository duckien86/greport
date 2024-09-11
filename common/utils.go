package common

import (
	"reflect"
)

func StructToMap(item interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	itemValue := reflect.ValueOf(item)
	itemType := reflect.TypeOf(item)

	for i := 0; i < itemValue.NumField(); i++ {
		if itemType.Field(i).PkgPath != "" {
			// Skip unexported fields
			continue
		}

		result[itemType.Field(i).Name] = itemValue.Field(i).Interface()
	}

	return result
}
