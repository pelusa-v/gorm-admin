package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type FiberHandler struct {
	BaseHandler
	App        *fiber.App
	TestModels []reflect.Type
}

func (handler *FiberHandler) RegisterModel(model reflect.Type) {
	handler.TestModels = append(handler.TestModels, model)
	handler.RegisterModelDetailPage(model)
}

func (handler *FiberHandler) Register() {
	fmt.Println("Registering admin in Fiber app")

	handler.RegisterHomePage()
}

func (handler *FiberHandler) RegisterHomePage() {
	tmpl, err := template.ParseFS(handler.TemplatesFs, "templates/home.html")
	if err != nil {
		panic(err)
	}

	handler.App.Get("/admin", func(c *fiber.Ctx) error {

		// first := handler.TestModels[0].Name()
		// Send the HTML content as the response
		var tmplOutput bytes.Buffer
		err := tmpl.Execute(&tmplOutput, handler.TestModels)
		if err != nil {
			return err
		}

		// Send the rendered HTML as the response
		c.Set("Content-Type", "text/html")
		return c.SendString(tmplOutput.String())
	})
}

type ModelDataDetail struct {
	Model    string
	Elements []interface{}
}

func (handler *FiberHandler) RegisterModelDetailPage(model reflect.Type) {
	tmpl, err := template.ParseFS(handler.TemplatesFs, "templates/ModelDetail.html")
	if err != nil {
		panic(err)
	}

	gormListHandler := NewGormListHandler(model, handler.GormDB)

	fmt.Printf("/admin/%s", model.Name())
	handler.App.Get(fmt.Sprintf("/admin/%s", model.Name()), func(c *fiber.Ctx) error {

		// first := handler.TestModels[0].Name()
		// Send the HTML content as the response
		var tmplOutput bytes.Buffer
		err := tmpl.Execute(&tmplOutput, ModelDataDetail{Model: model.Name(), Elements: gormListHandler.ListOjects()})
		if err != nil {
			return err
		}

		// Send the rendered HTML as the response
		c.Set("Content-Type", "text/html")
		return c.SendString(tmplOutput.String())
	})
}

type GormListHandler struct {
	modelType reflect.Type
	db        *gorm.DB
}

func NewGormListHandler(modelTypeMapped reflect.Type, db *gorm.DB) *GormListHandler {
	return &GormListHandler{
		modelType: modelTypeMapped,
		db:        db,
	}
}

func (h *GormListHandler) ListOjects() []interface{} {
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
