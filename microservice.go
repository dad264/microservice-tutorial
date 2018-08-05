package main

import (
	"fmt"
	"log"
	"microservice/api"
	"net/http"
	"os"

	"go.uber.org/zap"
	//"github.com/PacktPublishing/Cloud-Native-Go/api"
)

var build string

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	// compile passing -ldflags "-X main.Build <build sha1>"

	// fmt.Printf("Using build: %s\n", Build)
	sugar.Infof("Using build: %s", build)

	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", api.EchoHandleFunc)
	http.HandleFunc("/api/hello", api.HelloHandleFunc)

	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)

	// var url = "http://website.com"

	// sugar.Infow("failed to fetch URL",
	// 	// Structured context as loosely typed key-value pairs.
	// 	"url", url,
	// 	"attempt", 3,
	// 	"backoff", time.Second,
	// )

	sugar.Infof("starting Listener")
	http.ListenAndServe(port(), nil)
	sugar.Infof("stopping Listener")
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to Cloud Native Go (Update).")
	sugar.Infow("calling index")
}
