package user

import (
	"net/http"

	"github.com/soerjadi/wwalet/internal/config"
	"github.com/soerjadi/wwalet/internal/delivery/middleware"
	"github.com/soerjadi/wwalet/internal/delivery/rest"
	"github.com/soerjadi/wwalet/internal/usecase/user"

	"github.com/gorilla/mux"
)

func NewHandler(usecase user.Usecase, cfg *config.Config) rest.API {
	return &Handler{
		usecase: usecase,
		cfg:     cfg,
	}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.Use(middleware.OnlyLoggedInUser(h.usecase, h.cfg))
	r.HandleFunc("/register", rest.HandlerFunc(h.register).Serve).Methods(http.MethodPost)
	r.HandleFunc("/login", rest.HandlerFunc(h.login).Serve).Methods(http.MethodPost)
	r.HandleFunc("/update", rest.HandlerFunc(h.update).Serve).Methods(http.MethodPost)
}
