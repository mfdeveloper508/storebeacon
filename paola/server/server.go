package server

import (
	"fmt"
	"log"
	"net/http"
)

// Server config
type Server struct {
	Hostname string `json:"hostname"`
	HTTPPort int    `json:"http_port"`
}

// Starts HTTP listener
func Run(handlers http.Handler, s Server) {
	log.Println("Running HTTP on " + httpAddress(s))

	// Start HTTP listener
	log.Fatal(http.ListenAndServe(httpAddress(s), handlers))
}

func httpAddress(s Server) string {
	return s.Hostname + ":" + fmt.Sprintf("%d", s.HTTPPort)
}
