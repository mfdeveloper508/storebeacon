package route

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/storebeacon/backend/app/route/middleware/logrequest"
)

// Load providers a handler with all routes attached
func Load() http.Handler {
	return middleware(Routes())
}

func Routes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	return r
}

func middleware(h http.Handler) http.Handler {
	h = logrequest.Handler(h)

	return h
}
