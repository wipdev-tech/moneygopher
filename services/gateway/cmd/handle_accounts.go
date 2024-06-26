package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/wipdev-tech/moneygopher/services/accounts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type account struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Balance     int64  `json:"balance_dollars"`
}

func handleAccountsPost(w http.ResponseWriter, r *http.Request) {
	type In struct {
		PhoneNumber string `json:"phone_number"`
	}

	var in In
	json.NewDecoder(r.Body).Decode(&in)

	if in.PhoneNumber == "" {
		respondError(w, http.StatusBadRequest, "invalid phone number")
		return
	}

	conn, err := grpc.Dial("accounts:"+os.Getenv("ACCOUNTS_PORT"), insecureOpts...)
	if err != nil {
		fmt.Println("failed to dial grpc:", err)
		respondError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	defer conn.Close()

	accClient := accounts.NewAccountsClient(conn)

	ctx := context.Background()
	newID := uuid.NewString()
	resp, err := accClient.CreateAccount(ctx, &accounts.CreateAccountRequest{
		Id:          newID,
		PhoneNumber: in.PhoneNumber,
	})

	if err != nil {
		respondError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	acc := rpcAccountToAccount(resp)
	respondJSON(w, http.StatusCreated, acc)
}

func handleAccountsGet(w http.ResponseWriter, r *http.Request) {
	accountID := r.PathValue("accountID")
	if accountID == "" {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	conn, err := grpc.Dial("accounts:"+os.Getenv("ACCOUNTS_PORT"), insecureOpts...)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	defer conn.Close()

	accClient := accounts.NewAccountsClient(conn)
	resp, err := accClient.GetAccount(
		r.Context(),
		&accounts.GetAccountRequest{Id: accountID},
	)

	if status.Convert(err).Message() == sql.ErrNoRows.Error() {
		respondError(w, http.StatusNotFound, "not found")
		return
	}

	if err != nil {
		respondError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	acc := rpcAccountToAccount(resp)
	respondJSON(w, http.StatusCreated, acc)
}

func rpcAccountToAccount(rpcAcc *accounts.Account) account {
	return account{
		ID:          rpcAcc.Id,
		PhoneNumber: rpcAcc.PhoneNumber,
		Balance:     rpcAcc.Balance.Units,
	}
}
