package transaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/soerjadi/wwalet/internal/model"
)

func (h Handler) topup(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dec := json.NewDecoder(r.Body)

	req := model.TransactionRequest{}

	if err := dec.Decode(&req); err != nil {
		return nil, errors.New("fail to decode request body")
	}

	result, err := h.usecase.Topup(r.Context(), req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result, nil
}

func (h Handler) pay(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dec := json.NewDecoder(r.Body)

	req := model.TransactionRequest{}

	if err := dec.Decode(&req); err != nil {
		return nil, errors.New("fail to decode request body")
	}

	result, err := h.usecase.Payment(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (h Handler) transfer(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dec := json.NewDecoder(r.Body)

	req := model.TransactionRequest{}

	if err := dec.Decode(&req); err != nil {
		return nil, errors.New("fail to decode request body")
	}

	result, err := h.usecase.Transfer(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (h Handler) transactionList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return h.usecase.GetTransactionList(r.Context())
}
