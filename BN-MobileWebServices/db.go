package main

import (
	"log"

	"github.com/gorilla/mux"

	"github.com/go-bongo/bongo"
)

var connection = conDb()
var routers = router()

func conDb() *bongo.Connection {
	config := &bongo.Config{
		ConnectionString: "localhost",
		Database:         "BN",
	}
	connection, err := bongo.Connect(config)
	if err != nil {
		log.Fatal(err)
	}
	return connection
}
func router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return router
}
