package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	// "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type requestBody struct {
	Shorten string
}


func Redirect(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	sa := option.WithCredentialsFile("C:/Users/TheWisePigeon/Documents/gogoFront/api/creds.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	responseBody, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalf("Something went wrong %v", err)
	}

	var shorten requestBody
	err = json.Unmarshal(responseBody, &shorten)

	if err != nil {
		log.Fatalf("Something went wrong %v", err)
	}
	foundValues := make([]map[string]interface{}, 0)
	iter := client.Collection("links").Where("short", "==", "6053817c").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 
		}
		fmt.Println(doc.Data())
		foundValues = append(foundValues, doc.Data())
	}
	if len(foundValues)==0{
		fmt.Fprintf(w, "<h1>Link deleted or disabled</h1>")
		return
	}
	// urlToBeRedirectedTo  := fmt.Sprint(foundValues[0]["url"])
	http.Redirect(w, r, "https://google.com", http.StatusFound)
}
