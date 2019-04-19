package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/urfave/negroni"
	"github.com/urmilagera/auction/api/config"
	"github.com/urmilagera/auction/api/handler"
	"github.com/urmilagera/auction/api/middleware"
	"github.com/urmilagera/auction/pkg/bid"
	"github.com/urmilagera/auction/pkg/client"
	e "github.com/urmilagera/auction/pkg/entity_objects"
	"github.com/urmilagera/auction/pkg/offer"
)

func main() {
	config := config.GetAppConfig()
	var DB *gorm.DB

	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.GetDatabaseUsername(),
		config.GetDatabasePassword(),
		config.GetDatabaseName(),
		config.GetDatabaseCharset())

	db, err := gorm.Open(config.GetDatabaseDialect(), dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}
	DB = e.DBMigrate(db)

	r := mux.NewRouter()

	clientRepo := client.CreateMysqlRepository(DB)
	offerRepo := offer.CreateMysqlRepository(DB)
	bidRepo := bid.CreateMysqlRepository(DB)

	clientService := client.CreateService(clientRepo)
	offerService := offer.CreateService(offerRepo)
	bidService := bid.CreateService(bidRepo)

	authMiddleware := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)

	//Middleware for all other routes that require authentication
	apiMiddleware := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.JwtMiddleware(config.GetAppSecret())),
		negroni.HandlerFunc(middleware.LoginMiddleware(clientService)),
		negroni.NewLogger(),
	)
	handler.CreateClientHandlers(r, *authMiddleware, clientService)
	handler.CreateOfferHandlers(r, *apiMiddleware, offerService)
	handler.CreateBidHandlers(r, *apiMiddleware, bidService, offerService)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	error := http.ListenAndServe(":8040", context.ClearHandler(http.DefaultServeMux))
	if error != nil {
		fmt.Println(error.Error())
	}
}
