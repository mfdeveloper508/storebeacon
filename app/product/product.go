package product

import (
	"github.com/qor/render"
	"github.com/qor/admin"
	"github.com/storebeacon/backend/app/application"
	"github.com/storebeacon/backend/utils/funcmapmaker"
	"github.com/storebeacon/backend/app/api"
	"github.com/storebeacon/backend/models/product"
)

// Module for product controller
type Module struct {
	Config *Config
}

type AWSConfig struct {
	AwsAccess 	string `json:"aws_access"`
	AwsSecret 		string `json:"aws_secret"`
	AwsRegion 		string `json:"aws_region"`
	AwsBucket 		string `json:"aws_bucket"`
	AwsRule 		string `json:"aws_rule"`
}

// New instance of the news module
func New(config *Config) *Module {
	aws_info = config.AwsConfig
	return &Module{Config: config}
}

// Config for news module
type Config struct{
	AwsConfig AWSConfig
}

var aws_info AWSConfig

// Configure news module
func (m Module) Configure(application *application.Application) {
	controller := &Controller{
		View: render.New(
			&render.Config{AssetFileSystem: application.AssetFS.NameSpace("product")},
			"app/product/views",
		),
	}

	funcmapmaker.AddFuncMapMaker(controller.View)

	application.Router.HandleFunc("/products/", controller.Index)
	application.Router.HandleFunc("/products/edit/{id}", controller.Edit)
	application.Router.HandleFunc("/products/remove/{id}", controller.Remove)
	application.Router.HandleFunc("/products/add/", controller.Add)
	application.Router.HandleFunc("/products/register/", controller.Register)

	m.ConfigureAPI(application.Api)
	m.ConfigureAdmin(application.Admin)
}

// ConfigureAPI adds api endpoints
func (m Module) ConfigureAPI(api *api.API) {
	sr := api.Router.PathPrefix("/products").Subrouter()
	sr.HandleFunc("/upload/", UploadPOST).
		Methods("POST")
	sr.HandleFunc("/cvsupload/", CVSUploadPOST).
		Methods("POST")
	sr.HandleFunc("/bulkupload/", BulkUploadPOST).
		Methods("POST")
}

// ConfigureAdmin configured admin
func (m Module) ConfigureAdmin(adm *admin.Admin) {
	Products := adm.AddResource(&product.Product{}, &admin.Config{
		Name: "Products",
		Menu: []string{"Product Management"},
	})
	aisles:= adm.GetResource("Aisles")
	Products.Meta(&admin.Meta{Name: "Aisle", Type: "select_one", Config: &admin.SelectOneConfig{
		SelectMode:         "select_async",
		RemoteDataResource: aisles,
	}})
}
