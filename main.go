package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/urmilagera/auction-go-api/api/config"
	"github.com/urmilagera/auction-go-api/api/handler"
	"github.com/urmilagera/auction-go-api/api/infrastructure/mongo_db"
	"github.com/urmilagera/auction-go-api/api/middleware"
	"github.com/urmilagera/auction-go-api/pkg/bid"
	"github.com/urmilagera/auction-go-api/pkg/offer"

	"github.com/urfave/negroni"
)

func main() {
	config := config.GetDBConfig()

	mPool, session := mongo_db.GetMongoPool(
		config.GetDatabaseHostname(),
		config.GetDatabasePort(),
		config.GetConnectionPool(),
	)

	defer session.Close()
	defer mPool.Close()

	fmt.Println("Hello")
	r := mux.NewRouter()

	offerRepo := offer.CreateMongoRepository(mPool, config.GetDatabaseName())
	bidRepo := bid.CreateMongoRepository(mPool, config.GetDatabaseName())

	offerService := offer.CreateService(offerRepo)
	bidService := bid.CreateService(bidRepo)

	apiMiddleware := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)

	handler.CreateOfferHandlers(r, *apiMiddleware, offerService)
	handler.CreateBidHandlers(r, *apiMiddleware, bidService, offerService)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := &http.Server{
		Addr:    ":" + config.GetAppServerPort(),
		Handler: context.ClearHandler(http.DefaultServeMux),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Server is UP!!!!")
}
