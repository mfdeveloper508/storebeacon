package application

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/qor/admin"
	"github.com/qor/auth"
	"github.com/qor/assetfs"

	"github.com/storebeacon/backend/app/api"
	"github.com/storebeacon/backend/app/route"
	"github.com/storebeacon/backend/app/route/middleware/logrequest"
)

type AppModuleInterface interface {
	Configure(*Application)
}

// Application main application
type Application struct {
	*Config
}

type Config struct {
	httpMux *http.ServeMux
	Router  *mux.Router
	Admin   *admin.Admin
	Auth    *auth.Auth
	Api     *api.API
	AssetFS assetfs.Interface
}

func New(cfg *Config) *Application {
	if cfg == nil {
		cfg = &Config{}
	}

	if cfg.Router == nil {
		cfg.httpMux = http.NewServeMux()
	}

	if cfg.Admin != nil {
		cfg.Admin.MountTo("/admin", cfg.httpMux)
	}

	// qor auth routes
	if cfg.Auth != nil {
		cfg.httpMux.Handle("/auth/", cfg.Auth.NewServeMux())
	}

	// export downloads folder
	s := http.StripPrefix("/downloads/", http.FileServer(http.Dir("./downloads/")))
	cfg.httpMux.Handle("/downloads/", s)

	cfg.Router = route.Routes()

	Api := api.New(&api.Config{
		Router: cfg.Router.PathPrefix("/api").Subrouter(),
	})

	if cfg.AssetFS == nil {
		cfg.AssetFS = assetfs.AssetFS()
	}

	cfg.Api = Api

	return &Application{
		Config: cfg,
	}
}

// GetMux provides the *http.ServeMux for this application
func (application *Application) GetMux() *http.ServeMux {
	h := logrequest.Handler(application.Router)

	h = handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT"}),
		handlers.AllowedHeaders([]string{
			"Content-Type",
			"X-Requested-With",
			"Authorization",
		}),
	)(h)
	application.httpMux.Handle("/", h)

	return application.httpMux
}

func (application *Application) GetRawMux() *http.ServeMux {
	return application.httpMux
}

// Use attaches a module to the application
func (application *Application) Use(app AppModuleInterface) {
	app.Configure(application)
}
