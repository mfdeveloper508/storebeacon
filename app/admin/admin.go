package admin

import (
	"github.com/storebeacon/backend/app/application"
)

// App home app
type Module struct {
	Config *Config
}

// New new home app
func New(config *Config) *Module {
	return &Module{Config: config}
}

// Config home config struct
type Config struct{}

// ConfigureApplication configure application
func (app Module) Configure(application *application.Application) {
	Admin := application.Admin

	SetupAdminUsers(Admin)
	SetupDashboard(Admin)
}
