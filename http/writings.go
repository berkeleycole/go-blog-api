package http

import (
	"blog-api/writings"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router, writings *writings.Service) {
	r.HandleFunc("/writings", index(writings)).Methods("GET")
	r.HandleFunc("/writings/{id}", show(writings)).Methods("GET")
	r.HandleFunc("/writings", create(writings)).Methods("POST")

	// TODO - CRUD OPERATIONS
}

func index(writingsSrv *writings.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		results, err := writingsSrv.Index(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = json.NewEncoder(w).Encode(&results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func create(writingsSrv *writings.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		fmt.Fprintf(w, "creating the record\n")

		var e writings.NewWriting

		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		writing, err := writingsSrv.Create(ctx, e)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = json.NewEncoder(w).Encode(&writing); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func show(writings *writings.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		fmt.Fprintf(w, "getting the record\n")
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		writing, err := writings.Show(ctx, vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(&writing)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}
}
