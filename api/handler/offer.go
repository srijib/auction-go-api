package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	e "github.com/urmilagera/auction/pkg/entity_objects"
	"github.com/urmilagera/auction/pkg/offer"
	"gopkg.in/negroni"
)

func createOffer(service offer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var offer *e.Offer
		client := r.Context().Value("me").(*e.Client)
		errorMessage := "Error occured while creating an offer"
		err := json.NewDecoder(r.Body).Decode(&offer)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		// check if offer data is valid else return error
		if !offer.Validate() {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Data error"))
			return
		}
		offer.CreatedBy = client.Username
		offer.Id, err = service.Save(offer)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if err := json.NewEncoder(w).Encode(offer); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

	})
}

func getOffer(service offer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error occured while fetching offers"
		var offers []*e.Offer
		page, err := strconv.Atoi(r.FormValue("page"))
		if err != nil {
			page = 0
		}
		size, err := strconv.Atoi(r.FormValue("size"))
		if err != nil {
			size = 10
		}
		sortKey := r.FormValue("sortKey")

		offers, err = service.Query(page, size, sortKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		w.WriteHeader(http.StatusAccepted)
		if err := json.NewEncoder(w).Encode(offers); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//CreateUserHandlers Maps routes to http handlers
func CreateOfferHandlers(r *mux.Router, n negroni.Negroni, service offer.UseCase) {
	r.Handle("/offer", n.With(
		negroni.Wrap(createOffer(service)),
	)).Methods("POST", "OPTIONS").Name("CreateOffer")

	r.Handle("/offer", n.With(
		negroni.Wrap(getOffer(service)),
	)).Methods("GET", "OPTIONS").Name("GetOffers")
}
