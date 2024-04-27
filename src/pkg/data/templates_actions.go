package data

import (
	"fmt"
	"reflect"
)

func (manager *TemplateManager) CreateModelObjectAction(objectType reflect.Type, objectData map[string]interface{}) {
	// Check if the type is a struct or a pointer to a struct
	if objectType.Kind() != reflect.Struct {
		if objectType.Kind() == reflect.Ptr && objectType.Elem().Kind() == reflect.Struct {
			objectType = objectType.Elem()
		} else {
			panic("provided type is not a struct and not a pointer to a struct")
		}
	}

	// Create a new struct of the provided type
	newObject := reflect.New(objectType).Elem()

	// Iterate through the map and set field values
	for fieldName, value := range objectData {
		fieldVal := newObject.FieldByName(fieldName)
		if !fieldVal.IsValid() {
			fmt.Printf("No such field: %s in obj\n", fieldName)
			continue
		}
		if !fieldVal.CanSet() {
			fmt.Printf("Cannot set %s field value\n", fieldName)
			continue
		}

		fieldValue := reflect.ValueOf(value)
		if fieldVal.Type() != fieldValue.Type() {
			// Types do not match, attempt to convert types
			if fieldValue.Type().ConvertibleTo(fieldVal.Type()) {
				fieldValue = fieldValue.Convert(fieldVal.Type())
			} else {
				fmt.Printf("Provided value type didn't match obj field type and could not be converted\n")
				continue
			}
		}

		fieldVal.Set(fieldValue)
	}
}
