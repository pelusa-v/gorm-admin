package handlers

import (
	"fmt"
	"reflect"
)

type BuiltInHandler struct {
	BaseHandler
}

func (handler *BuiltInHandler) Register() {
	fmt.Println("Registering admin in BuiltIn http app")
}

func (handler *BuiltInHandler) RegisterModel(model reflect.Type) {

}
