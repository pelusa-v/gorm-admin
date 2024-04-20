package handlers

import (
	"embed"
	"fmt"
)

type GinHandler struct {
	App         *string
	TemplatesFs embed.FS
}

func (handler *GinHandler) Register() {
	fmt.Println("Registering admin in Gin app")
}
