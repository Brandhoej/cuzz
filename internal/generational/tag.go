package generational

import "reflect"

func Fill[T any]() (T, error) {
	var data T
	elem := reflect.ValueOf(data).Elem()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		tag := elem.Type().Field(i).Tag.Get("spec")

		if tag != "" {
			field.Set(reflect.ValueOf(1))
		}
	}

	return data, nil
}