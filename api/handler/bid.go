package handler

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/urmilagera/auction/pkg/bid"
	e "github.com/urmilagera/auction/pkg/entity_objects"
	"github.com/urmilagera/auction/pkg/offer"

	"net/http"
)

func placeBid(bidService bid.UseCase, offerService offer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var _bid *e.Bid
		errorMessage := "Error occured while Placing a Bid"
		client := r.Context().Value("me").(*e.Client)
		err := json.NewDecoder(r.Body).Decode(&_bid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error occured while Placing a Bid"))
			return
		}

		offer, err := offerService.Find(_bid.OfferID)
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

		offer, err = offerService.Update(_bid.OfferID, "bidprice", _bid.BidPrice)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error occured in Placing a Bid"))
			return
		}

		_bid.Username = client.Username
		_bid.Id, err = bidService.Save(_bid)
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

func acceptBid(bidService bid.UseCase, offerService offer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		bidID := e.StringToID(id)
		errorMessage := "Error Accepting Bid"

		_bid, err := bidService.Update(bidID, "accepted", true)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error Placing Bid"))
			return
		}

		_, err = offerService.Update(_bid.OfferID, "sold", true)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error Placing Bid"))
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

func CreateBidHandlers(r *mux.Router, n negroni.Negroni, bidService bid.UseCase, offerService offer.UseCase) {
	r.Handle("/bids", n.With(
		negroni.Wrap(placeBid(bidService, offerService)),
	)).Methods("POST", "OPTIONS").Name("placeBid")

	r.Handle("/bids/{id}", n.With(
		negroni.Wrap(acceptBid(bidService, offerService)),
	)).Methods("PUT", "OPTIONS").Name("acceptBid")

}
