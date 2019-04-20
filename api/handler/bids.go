package handler

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/urmilagera/auction/pkg/client"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/urmilagera/auction/pkg/bid"
	"github.com/urmilagera/auction/pkg/entity_objects"
	e "github.com/urmilagera/auction/pkg/entity_objects"
	"github.com/urmilagera/auction/pkg/offer"

	"net/http"
)

func placeBid(bidService bid.UseCase, offerService offer.UseCase, clientService client.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var _bid *entity_objects.Bid
		errorMessage := "Error occured while Placing a Bid"
		client := r.Context().Value("me").(*e.Client)
		err := json.NewDecoder(r.Body).Decode(&_bid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error occured while Placing a Bid"))
			return
		}
		offer, err := offerService.Find(_bid.OfferId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error occured while Placing a Bid"))
			return
		}

		if offer.BidPrice >= _bid.BidPrice {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error occured while Placing a Bid. BidPrice < Previous Bid Price"))
			return
		}

		_bid.Timestamp = time.Now()
		_bid.ClientId = client.Id
		_bid, err = bidService.Save(_bid)

		user, err := clientService.Find(client.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Authentication Failed"))
			return
		}

		_bid.Client.Id = user.Id
		_bid.Client.Email = user.Email
		_bid.Client.Username = user.Username

		offer, err = offerService.Update(_bid.OfferId, "bid_id", _bid.Id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error occured in Placing a Bid"))
			return
		}

		if err := json.NewEncoder(w).Encode(_bid); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

	})

}

func acceptBid(bidService bid.UseCase, offerService offer.UseCase, clientService client.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		client := r.Context().Value("me").(*e.Client)
		bidID, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		errorMessage := "Error Accepting Bid"

		_bid, err := bidService.Update(bidID, "accepted", true)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error Placing Bid"))
			return
		}

		_, err = offerService.Update(_bid.OfferId, "sold", true)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error Placing Bid"))
			return
		}
		user, err := clientService.Find(client.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Authentication Failed"))
			return
		}
		_bid.Timestamp = time.Now()
		_bid.Client.Id = user.Id
		_bid.Client.Email = user.Email
		_bid.Client.Username = user.Username
		_bid.ClientId = user.Id

		if err := json.NewEncoder(w).Encode(_bid); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

	})
}

func CreateBidHandlers(r *mux.Router, n negroni.Negroni, bidService bid.UseCase, offerService offer.UseCase, clientService client.UseCase) {
	r.Handle("/bids", n.With(
		negroni.Wrap(placeBid(bidService, offerService, clientService)),
	)).Methods("POST", "OPTIONS").Name("placeBid")

	r.Handle("/bids/{id}", n.With(
		negroni.Wrap(acceptBid(bidService, offerService, clientService)),
	)).Methods("PUT", "OPTIONS").Name("acceptBid")

}
