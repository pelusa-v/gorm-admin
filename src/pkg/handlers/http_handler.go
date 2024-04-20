package handlers

import (
	"embed"
	"fmt"
)

type BuiltInHandler struct {
	TemplatesFs embed.FS
}

func (handler *BuiltInHandler) Register() {
	fmt.Println("Registering admin in BuiltIn http app")
}
