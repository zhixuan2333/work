package main

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

var router = mux.NewRouter()
var client *db.Client

func main() {
	client = New()

	// server main method

	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/internal", internalPageHandler)

	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":80", nil)

}

// New init new fire client
func New() *db.Client {
	opt := option.WithCredentialsFile("go-db-48bf0-firebase-adminsdk-qjruj-e562bd373d.json")
	config := &firebase.Config{
		DatabaseURL: "https://go-db-48bf0-default-rtdb.firebaseio.com/",
	}

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal(err)
	}
	client, err = app.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(client)

	return client

}

// // NewUser Create new user func
// func NewUser(newuser string) error {
// 	ref := client.NewRef("users").Child(newuser.ID)
// 	err := ref.Set(context.Background(), newuser)
// 	if err != nil {
// 		log.Printf("Creat new user filed: %e", err)
// 		return err
// 	}
// 	return nil
// }
