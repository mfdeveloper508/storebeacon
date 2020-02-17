package aisle

import (
	"github.com/qor/render"
	"github.com/qor/admin"
	"github.com/storebeacon/backend/app/application"
	"github.com/storebeacon/backend/utils/funcmapmaker"
	"github.com/storebeacon/backend/app/api"
	"github.com/storebeacon/backend/models/aisle"
)

// Module for news controller
type Module struct {
	Config *Config
}

// New instance of the news module
func New(config *Config) *Module {
	return &Module{Config: config}
}

// Config for news module
type Config struct{}

// Configure news module
func (m Module) Configure(application *application.Application) {
	controller := &Controller{
		View: render.New(
			&render.Config{AssetFileSystem: application.AssetFS.NameSpace("aisle")},
			"app/aisle/views",
		),
	}

	funcmapmaker.AddFuncMapMaker(controller.View)

	application.Router.HandleFunc("/aisle/", controller.Index)
	application.Router.HandleFunc("/aisle/edit/{id}", controller.Edit)
	application.Router.HandleFunc("/aisle/remove/{id}", controller.Remove)
	application.Router.HandleFunc("/aisle/add/", controller.Add)
	application.Router.HandleFunc("/aisle/register/", controller.Register)

	m.ConfigureAPI(application.Api)
	m.ConfigureAdmin(application.Admin)
}

// ConfigureAPI adds api endpoints
func (m Module) ConfigureAPI(api *api.API) {

}

// ConfigureAdmin configured admin
func (m Module) ConfigureAdmin(adm *admin.Admin) {
	aisles := adm.AddResource(&aisle.Aisle{}, &admin.Config{
		Name: "Aisles",
		Menu: []string{"Store Management"},
	})
	stores:= adm.GetResource("Stores")
	aisles.Meta(&admin.Meta{Name: "Store", Type: "select_one", Config: &admin.SelectOneConfig{
		SelectMode:         "select_async",
		RemoteDataResource: stores,
	}})
}

