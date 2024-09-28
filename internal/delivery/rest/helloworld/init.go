package helloworld

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soerjadi/wwalet/internal/delivery/rest"
)

func NewHandler() rest.API {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/hello", rest.HandlerFunc(h.hello).Serve).Methods(http.MethodGet)
}
