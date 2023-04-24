package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/blockchain", GetBlockchain).Methods(http.MethodGet)
	router.HandleFunc("/api/blockchain", AddBlock).Methods(http.MethodPost)
	router.HandleFunc("/api/blockchain/{id}", UpdateBlock).Methods(http.MethodPut)
	router.HandleFunc("/api/blockchain/{id}", DeleteBlock).Methods(http.MethodDelete)

	router.Use(LoggerMiddleware)

	return router
}
