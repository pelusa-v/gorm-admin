package first

import (
	"reflect"

	"gorm.io/gorm"
)

type OperationType string

const (
	Create OperationType = "Create"
	Read   OperationType = "Read"
	Update OperationType = "Update"
	Delete OperationType = "Delete"
)

type CrudHandler struct {
	modelType reflect.Type
	db        *gorm.DB
}

func NewCrudHandler(modelTypeMapped reflect.Type, db *gorm.DB) *CrudHandler {
	return &CrudHandler{
		modelType: modelTypeMapped,
		db:        db,
	}
}

func (h *CrudHandler) ListOjects() []interface{} {
	objectsType := reflect.SliceOf(h.modelType)
	concreteObjects := reflect.New(objectsType).Interface()

	concrete := reflect.New(h.modelType).Interface()
	query := h.db.Model(concrete)
	query.Find(concreteObjects)

	concreteSliceValue := reflect.ValueOf(concreteObjects).Elem()
	resultSlice := make([]interface{}, concreteSliceValue.Len())
	for i := 0; i < concreteSliceValue.Len(); i++ {
		resultSlice[i] = concreteSliceValue.Index(i).Interface()
	}

	return resultSlice
}
