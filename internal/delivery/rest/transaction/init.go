package transaction

import (
	"net/http"

	"github.com/soerjadi/wwalet/internal/config"
	"github.com/soerjadi/wwalet/internal/delivery/middleware"
	"github.com/soerjadi/wwalet/internal/delivery/rest"
	"github.com/soerjadi/wwalet/internal/usecase/transaction"
	"github.com/soerjadi/wwalet/internal/usecase/user"

	"github.com/gorilla/mux"
)

func NewHandler(usecase transaction.Usecase, userUsecase user.Usecase, cfg *config.Config) rest.API {
	return &Handler{
		usecase:     usecase,
		userUsecase: userUsecase,
		cfg:         cfg,
	}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.Use(middleware.OnlyLoggedInUser(h.userUsecase, h.cfg))
	r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/topup", rest.HandlerFunc(h.topup).Serve).Methods(http.MethodGet)
	r.HandleFunc("/transfers", rest.HandlerFunc(h.transfer).Serve).Methods(http.MethodGet)
	r.HandleFunc("/pay", rest.HandlerFunc(h.pay).Serve).Methods(http.MethodGet)
	r.HandleFunc("/transactions", rest.HandlerFunc(h.transactionList).Serve).Methods(http.MethodGet)
}
