package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type Bid struct {
	Bidprice float64 `json:"bidprice"`
	Clientid float64 `json:"clientid"`
	Offerid  float64 `json:"offerid"`
}

func CreateBidEndpoint(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var bid Bid
	err := decoder.Decode(&bid)
	if err != nil {
		panic(err)
	}
	fmt.Println(bid.Bidprice)
	// params := mux.Vars(r)
	// var bid Bid
	// _ = json.NewDecoder(r.Body).Decode(&bid)
	// var bidPrice = params["bidprice"]
	// if s, err := strconv.ParseFloat(bidPrice, 64); err == nil {
	// 	bid.bidprice = s
	// }
	// var clientid = params["clientid"]
	// if f, err := strconv.ParseFloat(clientid, 64); err == nil {
	// 	bid.clientid = f
	// }
	// var offerid = params["offerid"]
	// if g, err := strconv.ParseFloat(offerid, 64); err == nil {
	// 	bid.offerid = g
	// }
	// json.NewEncoder(w).Encode(bid)

}

// our main function
func main() {
	router := mux.NewRouter()
	fmt.Println("Hello world")
	router.HandleFunc("/bid", CreateBidEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":9090", router))
}
