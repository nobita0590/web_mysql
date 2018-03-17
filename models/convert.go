package models

import (
	"reflect"
	"fmt"
	"errors"
)

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	fmt.Println(structFieldType)
	fmt.Println(val.Type())
	if structFieldType != val.Type() {
		invalidTypeError := errors.New("Provided value type didn't match obj field type in field: " + name)
		return invalidTypeError
	}

	structFieldValue.Set(val)
	return nil
}

func FillStruct(m map[string]interface{},s interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}