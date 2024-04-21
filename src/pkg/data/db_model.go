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

func (m *DbModel) ListOjects() []interface{} {
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
