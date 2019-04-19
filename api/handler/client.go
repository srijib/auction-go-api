package handler

import (
	"encoding/json"
	"net/http"

	"github.com/urmilagera/auction/api/config"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/urmilagera/auction/pkg/client"
	c "github.com/urmilagera/auction/pkg/client"
	e "github.com/urmilagera/auction/pkg/entity_objects"
)

func signup(service client.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error sigining up user"
		var _client *e.Client

		err := json.NewDecoder(r.Body).Decode(&_client)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		data, err := service.FindByUsername(_client.Username)
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
		_client.Password = c.SaltPassowrd(_client.Password)
		_client, err = service.Save(_client)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		w.WriteHeader(http.StatusCreated)

	})
}

func login(service client.UseCase) http.Handler {
	cfg := config.GetAppConfig()
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
		clientDatum := data[0]

		if !c.ComparePasswords(clientDatum.Password, client.Password) {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Password Doesnot Match"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != e.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		jwtmap := clientDatum.GenerateJWT([]byte(cfg.GetAppSecret()))
		if err := json.NewEncoder(w).Encode(jwtmap); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

	})
}

func CreateClientHandlers(r *mux.Router, n negroni.Negroni, clientService client.UseCase) {
	r.Handle("/login", n.With(
		negroni.Wrap(login(clientService)),
	)).Methods("POST", "OPTIONS").Name("login")

	r.Handle("/signup", n.With(
		negroni.Wrap(signup(clientService)),
	)).Methods("POST", "OPTIONS").Name("signup")
}
