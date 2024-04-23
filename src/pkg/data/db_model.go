package data

import (
	"reflect"

	"gorm.io/gorm"
)

type DbModel struct {
	modelType reflect.Type
	db        *gorm.DB
}

func NewDbModel(modelTypeMapped reflect.Type, db *gorm.DB) *DbModel {
	return &DbModel{
		modelType: modelTypeMapped,
		db:        db,
	}
}

func (m *DbModel) ListObjects() []interface{} {
	objectsType := reflect.SliceOf(m.modelType)
	concreteObjects := reflect.New(objectsType).Interface()

	concrete := reflect.New(m.modelType).Interface()
	query := m.db.Model(concrete)
	query.Find(concreteObjects)

	concreteSliceValue := reflect.ValueOf(concreteObjects).Elem()
	resultSlice := make([]interface{}, concreteSliceValue.Len())

	for i := 0; i < concreteSliceValue.Len(); i++ {
		resultSlice[i] = concreteSliceValue.Index(i).Interface()
	}

	return resultSlice
}

// func (m *DbModel) ListObjects() []DbObjectInstance {
// 	objectsType := reflect.SliceOf(m.modelType)
// 	concreteDbObjects := reflect.New(objectsType).Interface()

// 	concreteStruct := reflect.New(m.modelType).Interface()
// 	query := m.db.Model(concreteStruct)
// 	query.Find(concreteDbObjects)

// 	concreteDbObjectsValue := reflect.ValueOf(concreteDbObjects).Elem()
// 	// resultSlice := make([]interface{}, concreteDbObjectsValue.Len()) // To delete

// 	objects := make([]DbObjectInstance, concreteDbObjectsValue.Len())

// 	for i := 0; i < concreteDbObjectsValue.Len(); i++ {
// 		dbObject := concreteDbObjectsValue.Index(i)
// 		object := DbObjectInstance{}
// 		var objectProperties []reflect.StructField
// 		var objectPropertiesValues []reflect.Value

// 		for i := 0; i < dbObject.NumField(); i++ {
// 			propertyValue := dbObject.Field(i)
// 			property := dbObject.Type().Field(i)

// 			objectPropertiesValues = append(objectPropertiesValues, propertyValue)
// 			objectProperties = append(objectProperties, property)

// 			fmt.Println("---------------------------")
// 			fmt.Printf("%v :\n", dbObject.Type().Field(i).Name)
// 			fmt.Printf("%v :\n", dbObject.Type().Field(i).Type)
// 			fmt.Printf("%v :\n", dbObject.Type().Field(i).Tag)
// 			fmt.Printf("%v :\n", dbObject.Type().Field(i))
// 			fmt.Println(dbObject.Field(i))
// 		}

// 		object.Fields = objectProperties
// 		object.FieldsValues = objectPropertiesValues
// 		objects = append(objects, object)
// 		// resultSlice[i] = concreteDbObjectsValue.Index(i).Interface() // To delete
// 	}

// 	return objects
// }

func (m *DbModel) GetModelFields() []reflect.StructField {
	var fields []reflect.StructField
	for i := 0; i < m.modelType.NumField(); i++ {
		fields = append(fields, m.modelType.Field(i))
	}
	return fields
}
