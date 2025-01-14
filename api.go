package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddr string
	storage Storage
}

func NewApiServer(listenAddr string, storage Storage) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
		storage: storage,
	}
}

func (server *ApiServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHandlerFunc(server.handleAccount))

	//setup server
	slog.Info("Starting server on ", slog.String("port", server.listenAddr))

	srvr := http.Server{
		Addr:    server.listenAddr,
		Handler: router,
	}

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srvr.ListenAndServe(); err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	//handling gracefully server shutdown

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	err := srvr.Shutdown(ctx)

	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")
}

func (server *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return server.handleGetAccount(w, r)
	}

	if r.Method == http.MethodPost {
		return server.handleCreateAccount(w, r)
	}

	if r.Method == http.MethodDelete {
		return server.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (server *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	account := NewAccount("Anthony", "GG")
	return WriteJson(w, http.StatusOK, account)
}

func (server *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (server *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (server *ApiServer) handleTransferMoney(w http.ResponseWriter, r *http.Request) error {
	return nil
}



func WriteJson(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
}

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type apiProcessFunc func(http.ResponseWriter, *http.Request) error

func makeHandlerFunc(f apiProcessFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			WriteJson(w, http.StatusBadRequest, ErrorResponse{
				Error:   true,
				Message: err.Error(),
			})
		}
	}
}
