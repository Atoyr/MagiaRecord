package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	firebase "firebase.google.com/go"
	"github.com/atoyr/MagiaRecord/store/models"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	// load json
	bytes, err := ioutil.ReadFile("../data/magicalGirls.json")
	if err != nil {
		log.Fatal(err)
	}
	// JSONデコード
	var magicalGirls []models.MagicalGirl
	if err := json.Unmarshal(bytes, &magicalGirls); err != nil {
		log.Fatal(err)
	}

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("./_magia-record-database.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("Failed to 1: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to 2: %v", err)
	}
	defer client.Close()

	for i := range magicalGirls {
		_, err = client.Collection("private/v1/magicalGirls").Doc(magicalGirls[i].Key).Set(ctx, magicalGirls[i])
		if err != nil {
			log.Fatalf("Failed adding alovelace: %v", err)
		}
	}

	iter := client.Collection("private/v1/magicalGirls").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
}
