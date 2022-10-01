package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"encoding/json"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"io"

	"github.com/google/uuid"
)

type longUrl struct {
	Url string
	Short string
}

func Handler(w http.ResponseWriter, r *http.Request) {
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
	requestBody, err := (io.ReadAll(r.Body))
	if err != nil {
		log.Fatalf("Something went wrong %v", err)
	}
	var url longUrl
	var generated = uuid.New().String()
	generated = strings.Split(generated, "-")[0]
	fmt.Println(generated)
	url.Short = generated
	err = json.Unmarshal(requestBody, &url)
	if err != nil {
		log.Fatalf(`Something went wrong %v`, err)
	}

	_, _, err = client.Collection("links").Add(ctx, map[string]interface{}{
		"url": url.Url,
		"short":  generated,
	})

	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)
}
