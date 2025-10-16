package server

import (
	"authentication-service/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct {
	Router *mux.Router
	cors   *cors.Cors
}

func New() *Server {
	return &Server{
		Router: mux.NewRouter(),
		cors: cors.New(cors.Options{
			AllowedHeaders:   []string{"*"},
			AllowedMethods:   []string{"GET", "PUT", "DELETE", "POST", "OPTIONS"},
			AllowedOrigins:   []string{utils.Env.URLSite},
			AllowCredentials: true,
		}),
	}
}

func (c *Server) routeDefault(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"Message": "OK"}`))
}

func (c *Server) Listen() {
	c.Router.HandleFunc("/", c.routeDefault).Methods("GET")

	handler := c.cors.Handler(c.Router)

	fmt.Println("Run on port http://localhost:3033/")

	log.Fatal(http.ListenAndServe(":3033", handler))

}
