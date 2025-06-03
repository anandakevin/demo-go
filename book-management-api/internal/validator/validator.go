package validator

import (
	"reflect"
	"strconv"
)

func SetDefaults(i interface{}) {
	t := reflect.TypeOf(i).Elem()
	v := reflect.ValueOf(i).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if defaultVal := field.Tag.Get("default"); defaultVal != "" && value.IsZero() {
			switch value.Kind() {
			case reflect.String:
				value.SetString(defaultVal)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if intVal, err := strconv.ParseInt(defaultVal, 10, 64); err == nil {
					value.SetInt(intVal)
				}
			}
		}
	}
}
