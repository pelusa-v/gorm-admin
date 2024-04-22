package data

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type DbModel struct {
	modelType reflect.Type
	db        *gorm.DB
}

type DbObjectInstance struct {
	// PropertyNames []string
	Properties []DbObjectInstanceProperty
	DetailURL  string
}
type DbObjectInstanceProperty struct {
	FieldData reflect.StructField
	Value     reflect.Value
}

func NewDbModel(modelTypeMapped reflect.Type, db *gorm.DB) *DbModel {
	return &DbModel{
		modelType: modelTypeMapped,
		db:        db,
	}
}

// func (m *DbModel) ListObjects() []interface{} {
// 	objectsType := reflect.SliceOf(m.modelType)
// 	concreteObjects := reflect.New(objectsType).Interface()

// 	concrete := reflect.New(m.modelType).Interface()
// 	query := m.db.Model(concrete)
// 	query.Find(concreteObjects)

// 	concreteSliceValue := reflect.ValueOf(concreteObjects).Elem()
// 	resultSlice := make([]interface{}, concreteSliceValue.Len())

// 	for i := 0; i < concreteSliceValue.Len(); i++ {
// 		resultSlice[i] = concreteSliceValue.Index(i).Interface()
// 	}

// 	return resultSlice
// }

func (m *DbModel) ListObjects() []DbObjectInstance {
	objectsType := reflect.SliceOf(m.modelType)
	concreteDbObjects := reflect.New(objectsType).Interface()

	concreteStruct := reflect.New(m.modelType).Interface()
	query := m.db.Model(concreteStruct)
	query.Find(concreteDbObjects)

	concreteDbObjectsValue := reflect.ValueOf(concreteDbObjects).Elem()
	// resultSlice := make([]interface{}, concreteDbObjectsValue.Len()) // To delete

	objects := make([]DbObjectInstance, concreteDbObjectsValue.Len())

	for i := 0; i < concreteDbObjectsValue.Len(); i++ {
		dbObject := concreteDbObjectsValue.Index(i)
		object := DbObjectInstance{}
		var objectProperties []DbObjectInstanceProperty

		for i := 0; i < dbObject.NumField(); i++ {
			property := DbObjectInstanceProperty{
				FieldData: dbObject.Type().Field(i),
				Value:     dbObject.Field(i),
			}
			objectProperties = append(objectProperties, property)

			fmt.Println("---------------------------")
			fmt.Printf("%v :\n", dbObject.Type().Field(i).Name)
			fmt.Printf("%v :\n", dbObject.Type().Field(i).Type)
			fmt.Printf("%v :\n", dbObject.Type().Field(i).Tag)
			fmt.Printf("%v :\n", dbObject.Type().Field(i))
			fmt.Println(dbObject.Field(i))
		}

		object.Properties = objectProperties

		objects = append(objects, object)
		// resultSlice[i] = concreteDbObjectsValue.Index(i).Interface() // To delete
	}

	return objects
}
