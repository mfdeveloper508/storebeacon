package store

import (
	"github.com/qor/render"
	"github.com/qor/admin"
	"github.com/storebeacon/backend/app/application"
	"github.com/storebeacon/backend/utils/funcmapmaker"
	"github.com/storebeacon/backend/app/api"
	"github.com/storebeacon/backend/models/store"
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
			&render.Config{AssetFileSystem: application.AssetFS.NameSpace("store")},
			"app/store/views",
		),
	}

	funcmapmaker.AddFuncMapMaker(controller.View)


	application.Router.HandleFunc("/", controller.Home)
	application.Router.HandleFunc("/store/", controller.Index)
	application.Router.HandleFunc("/verify/login/", controller.Login)
	application.Router.HandleFunc("/verify/store/", controller.Store)
	application.Router.HandleFunc("/store/register/", controller.Register)

	m.ConfigureAPI(application.Api)
	m.ConfigureAdmin(application.Admin)
}

// ConfigureAPI adds api endpoints
func (m Module) ConfigureAPI(api *api.API) {
	sr := api.Router.PathPrefix("/store").Subrouter()
	sr.HandleFunc("/upload/", UploadPOST).
		Methods("POST")
}

// ConfigureAdmin configured admin
func (m Module) ConfigureAdmin(adm *admin.Admin) {
	adm.AddResource(&store.Store{}, &admin.Config{
		Name: "Stores",
		Menu: []string{"Store Management"},
	})
}


