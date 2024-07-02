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

func (m *DbModel) GetObject(pk string) interface{} {
	concreteObject := reflect.New(m.modelType).Interface()
	m.db.First(&concreteObject, pk)
	return concreteObject
}

// func (m *DbModel) CreateObject(newObject interface{}) error {
// 	res := m.db.Create(newObject)
// 	return res.Error
// }

func (m *DbModel) CreateObject(newObject interface{}, objectType reflect.Type) error {
	res := m.db.Model(reflect.New(objectType).Interface()).Create(newObject)
	return res.Error
}

func (m *DbModel) UpdateObject(newObject interface{}) error {
	res := m.db.Save(newObject)
	return res.Error
}

func (m *DbModel) DeleteObject(pk interface{}) error {
	concreteObject := reflect.New(m.modelType).Interface()
	res := m.db.Delete(concreteObject, pk)
	return res.Error
}
