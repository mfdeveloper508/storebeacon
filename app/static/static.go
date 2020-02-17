package static

import (
	"net/http"
	"github.com/storebeacon/backend/app/application"
)

// New new home app
func New(config *Config) *App {
	return &App{Config: config}
}

// App home app
type App struct {
	Config *Config
}

// Config home config struct
type Config struct {
	Prefixs []string
	Handler http.Handler
}
var (
	Root = "public"
)

// Configure configure application
func (app App) Configure(application *application.Application) {
	application.Router.PathPrefix("/system/").Handler(http.FileServer(http.Dir(Root)))
	application.Router.PathPrefix("/bower_components/").Handler(http.FileServer(http.Dir(Root)))
	application.Router.PathPrefix("/dist/").Handler(http.FileServer(http.Dir(Root)))
	application.Router.PathPrefix("/plugins/").Handler(http.FileServer(http.Dir(Root)))
}
