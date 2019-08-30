package main

import (
	_ "dogego/conf"
	"dogego/protocol/http"
	"dogego/routers"
	"log"
	"os"
)

func main() {
	router := routers.NewRouter()
	server := http.NewHttpProtocol(router)

	err := server.Start(os.Getenv("ADDR"))

	if err != nil {
		log.Fatal(err)
	}
}
