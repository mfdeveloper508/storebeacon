package beacon

import (
	"github.com/qor/render"
	"github.com/qor/admin"
	"github.com/storebeacon/backend/app/application"
	"github.com/storebeacon/backend/utils/funcmapmaker"
	"github.com/storebeacon/backend/app/api"
	"github.com/storebeacon/backend/models/beacon"
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
			&render.Config{AssetFileSystem: application.AssetFS.NameSpace("beacon")},
			"app/beacon/views",
		),
	}

	funcmapmaker.AddFuncMapMaker(controller.View)

	application.Router.HandleFunc("/beacons/", controller.Index)
	application.Router.HandleFunc("/beacons/edit/{id}", controller.Edit)
	application.Router.HandleFunc("/beacons/remove/{id}", controller.Remove)
	application.Router.HandleFunc("/beacons/add/", controller.Add)
	application.Router.HandleFunc("/beacons/register/", controller.Register)

	m.ConfigureAPI(application.Api)
	m.ConfigureAdmin(application.Admin)
}

// ConfigureAPI adds api endpoints
func (m Module) ConfigureAPI(api *api.API) {
	api.Router.HandleFunc("", PostProc).
		Methods("POST")
}

// ConfigureAdmin configured admin
func (m Module) ConfigureAdmin(adm *admin.Admin) {
	beacons := adm.AddResource(&beacon.Beacon{}, &admin.Config{
		Name: "Beacons",
		Menu: []string{"Store Management"},
	})

	aisles:= adm.GetResource("Aisles")
	stores:= adm.GetResource("Stores")
	beacons.Meta(&admin.Meta{Name: "Aisle", Type: "select_one", Config: &admin.SelectOneConfig{
		SelectMode:         "select_async",
		RemoteDataResource: aisles,
	}})
	beacons.Meta(&admin.Meta{Name: "Store", Type: "select_one", Config: &admin.SelectOneConfig{
		SelectMode:         "select_async",
		RemoteDataResource: stores,
	}})
}

