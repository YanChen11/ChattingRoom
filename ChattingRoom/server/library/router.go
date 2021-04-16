package library

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/ws", Chat)

	c := cors.New(cors.Options{
		AllowedOrigins:         []string{"*"},
	})

	handler := c.Handler(myRouter)

	log.Fatal(http.ListenAndServe(":8081", handler))
}