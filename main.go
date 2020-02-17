// @APIVersion 1.0.0
// @APITitle STORE API
// @APIDescription API for STORE app.
// @BasePath http://host:port/api
package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/qor/admin"
	"github.com/qor/location"
	"github.com/qor/media"
	"github.com/storebeacon/backend/paola/auth"
	"github.com/storebeacon/backend/paola/database"
	"github.com/storebeacon/backend/paola/jsonconfig"
	"github.com/storebeacon/backend/paola/server"

	adminapp "github.com/storebeacon/backend/app/admin"
	"github.com/storebeacon/backend/app/application"
	"github.com/storebeacon/backend/app/static"
	"github.com/storebeacon/backend/models/migrations"
	"github.com/storebeacon/backend/models/admins"
	"github.com/storebeacon/backend/app/store"
	"github.com/storebeacon/backend/app/beacon"
	"github.com/storebeacon/backend/app/aisle"
	"github.com/storebeacon/backend/app/product"
)

var (
	config = &Configuration{}
)

type Configuration struct {
	Server   server.Server           `json:"server"`
	Database database.DatabaseConfig `json:"database"`
	AWS 	  product.AWSConfig 	 `json:"aws"`
}

// ParseJSON unmarshals bytes to structs
func (c *Configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

// LoadComponents parses the JSON file and sets up components
func LoadComponents() *Configuration {
	configPath := "config.json"
	jsonconfig.Load(configPath, config)

	database.Configure(config.Database)
	// Configure database rds or local test
	host := os.Getenv("RDS_HOSTNAME")
	port := os.Getenv("RDS_PORT")
	username := os.Getenv("RDS_USERNAME")
	password := os.Getenv("RDS_PASSWORD")
	db := os.Getenv("RDS_DB_NAME")

	if host != "" {
		configDatabase := database.DatabaseConfig{}
		portInt, _ := strconv.Atoi(port)

		configDatabase.Host = host
		configDatabase.Port = portInt
		configDatabase.User = username
		configDatabase.Password = password
		configDatabase.Database = db

		database.Configure(configDatabase)
	}
	database.Connect()

	media.RegisterCallbacks(database.Conn)

	// Update the database
	migrations.Migrate(database.Conn)

	location.GoogleAPIKey = os.Getenv("GOOGLE_API_KEY")

	return config
}

func main() {
	LoadComponents()

	Admin := admin.New(&admin.AdminConfig{
		SiteName: "STORE API",
		DB:       database.Conn,
		Auth:     auth.AdminAuth{},
	})

	app := application.New(&application.Config{
		Admin: Admin,
		Auth:  auth.NewAuth(),
	})

	// check if there are admins, if not create one
	adminUsers, _ := admins.GetAdminUsers()
	if len(adminUsers) == 0 {
		adm := admins.Admin{
			Name:     "Admin",
			Email:    "admin@admin.com",
			Password: "admin123!",
			Role:     "Admin",
		}
		database.Conn.Create(&adm)
	}

	app.Use(adminapp.New(&adminapp.Config{}))
	app.Use(static.New(&static.Config{}))
	app.Use(store.New(&store.Config{}))
	app.Use(aisle.New(&aisle.Config{}))
	app.Use(beacon.New(&beacon.Config{}))
	app.Use(product.New(&product.Config{AwsConfig:config.AWS}))

	server.Run(app.GetMux(), config.Server)
}
