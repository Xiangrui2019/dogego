package main

import (
	_ "dogego/conf"
	"dogego/routers"
	"log"
	"net/http"
	"os"
)

func main() {
	router := routers.NewRouter()

	httpServer := http.Server{
		Addr:    os.Getenv("ADDR"),
		Handler: router,
	}

	err := httpServer.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
