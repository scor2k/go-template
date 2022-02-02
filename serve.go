package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func enableJsonCors(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	enableJsonCors(&w)

	msg, _ := json.Marshal(map[string]interface{}{"result": "ok", "msg": fmt.Sprintf("APP_NAME v%v", appVersion)})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(msg)
}

// Serve - start API server and handle requests
func Serve() {
	router := mux.NewRouter()

	router.HandleFunc("/health/check", HealthCheck).Methods("GET")

	router.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE")
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Headers", "*")
			w.WriteHeader(http.StatusOK)
		})

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Info("Start listening on :8080")
	log.Fatal(srv.ListenAndServe())
}
