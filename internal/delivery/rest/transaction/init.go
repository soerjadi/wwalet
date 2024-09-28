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
	r.Use(mux.CORSMethodMiddleware(r))

	topup := r.PathPrefix("/topup").Subrouter()
	topup.Use(middleware.OnlyLoggedInUser(h.userUsecase, h.cfg))
	topup.HandleFunc("/", rest.HandlerFunc(h.topup).Serve).Methods(http.MethodPost)

	transfer := r.PathPrefix("/transfers").Subrouter()
	transfer.Use(middleware.OnlyLoggedInUser(h.userUsecase, h.cfg))
	transfer.HandleFunc("", rest.HandlerFunc(h.transfer).Serve).Methods(http.MethodPost)

	pay := r.PathPrefix("/pay").Subrouter()
	pay.Use(middleware.OnlyLoggedInUser(h.userUsecase, h.cfg))
	pay.HandleFunc("", rest.HandlerFunc(h.pay).Serve).Methods(http.MethodPost)

	trx := r.PathPrefix("/transactions").Subrouter()
	trx.Use(middleware.OnlyLoggedInUser(h.userUsecase, h.cfg))
	trx.HandleFunc("", rest.HandlerFunc(h.transactionList).Serve).Methods(http.MethodPost)
}
