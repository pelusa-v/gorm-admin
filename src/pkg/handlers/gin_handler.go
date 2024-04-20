package handlers

import (
	"fmt"
)

type GinHandler struct {
	BaseHandler
	App *string
}

func (handler *GinHandler) Register() {
	fmt.Println("Registering admin in Gin app")
}

func (handler *GinHandler) RegisterModel(model string) {

}
