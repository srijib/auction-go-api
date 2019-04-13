package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	user "github.com/urmilagera/auction/pkg/client"
	e "github.com/urmilagera/auction/pkg/entity_objects"
)

func signup(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error sigining up user"
		var client *e.Client

		err := json.NewDecoder(r.Body).Decode(&client)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		data, err := service.FindByUsername(client.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if data != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("User Already Exist"))
			return
		}
		client.Password = "12345"
		client.Id, err = service.Save(client)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		w.WriteHeader(http.StatusCreated)

	})
}

func login(service user.UseCase) http.Handler {
	//cfg := config.GetDBConfig()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error occured while finding the Client"
		var client *e.Client
		err := json.NewDecoder(r.Body).Decode(&client)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.FindByUsername(client.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("User Doesn't Exist"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != e.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

	})
}

func CreateUserHandlers(r *mux.Router, n negroni.Negroni, service user.UseCase) {
	r.Handle("/login", n.With(
		negroni.Wrap(login(service)),
	)).Methods("POST", "OPTIONS").Name("login")

	r.Handle("/signup", n.With(
		negroni.Wrap(signup(service)),
	)).Methods("POST", "OPTIONS").Name("signup")
}
