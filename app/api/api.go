package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// API api module
type API struct {
	*Config
	MemberMiddleware alice.Chain
}

// Config holds path prefix and router
type Config struct {
	Router *mux.Router
}

// New creates new instance of API
func New(cfg *Config) *API {
	return &API{
		Config:           cfg,
		MemberMiddleware: alice.New(ForceLoginMiddleware),
	}
}

// BadRequest writes to response writer
func BadRequest(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

// InternalServerError writes to response writer
func InternalServerError(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
