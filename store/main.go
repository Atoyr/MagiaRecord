package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/atoyr/MagiaRecord/store/models"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("firebase key file not set")
		return
	}
	args := flag.Args()
	file := args[0]
	_, err := os.Stat(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile(file)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("Failed to 1: %v", err)
	}

	mergeMagicalGirls(ctx, app)
	mergeFirestore(ctx, app, "effectTypes")
	mergeFirestore(ctx, app, "effectCategories")
	mergeFirestore(ctx, app, "effectActions")
	mergeFirestore(ctx, app, "effectActivations")
	mergeEffects(ctx, app)
}

func mergeMagicalGirls(ctx context.Context, app *firebase.App) error {
	// load json
	bytes, err := ioutil.ReadFile("../data/magicalGirls.json")
	if err != nil {
		log.Fatalf("Fail to load magicalGirls.json : %v", err)
	}
	// JSONデコード
	var magicalGirls []models.MagicalGirl
	if err := json.Unmarshal(bytes, &magicalGirls); err != nil {
		log.Fatalf("Fail to Convert magicalGirls.json : %v", err)
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
	return nil
}

func mergeFirestore(ctx context.Context, app *firebase.App, name string) error {
	// load json
	bytes, err := ioutil.ReadFile("../data/" + name + ".json")
	if err != nil {
		log.Fatalf("Fail to load %s.json : %v", name, err)
	}
	// JSONデコード
	var sli []string
	if err := json.Unmarshal(bytes, &sli); err != nil {
		log.Fatalf("Fail to Convert %s.json : %v", name, err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to 2: %v", err)
	}
	defer client.Close()

	type dummy struct {
		Name string `json:"name"`
	}

	for i := range sli {
		_, err = client.Collection("private/v1/"+name).Doc(sli[i]).Set(ctx, dummy{Name: sli[i]})
		if err != nil {
			log.Fatalf("Failed adding alovelace: %v", err)
		}
	}
	return nil
}

func mergeEffects(ctx context.Context, app *firebase.App) error {
	// load json
	bytes, err := ioutil.ReadFile("../data/effects.json")
	if err != nil {
		log.Fatalf("Fail to load effects.json : %v", err)
	}
	// JSONデコード
	var effects []models.Effect
	if err := json.Unmarshal(bytes, &effects); err != nil {
		log.Fatalf("Fail to Convert effects.json : %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to 2: %v", err)
	}
	defer client.Close()

	for i := range effects {
		_, err = client.Collection("private/v1/effects").Doc(effects[i].Key).Set(ctx, effects[i])
		if err != nil {
			log.Fatalf("Failed adding alovelace: %v", err)
		}
	}

	iter := client.Collection("private/v1/effects").Documents(ctx)
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
	return nil
}
