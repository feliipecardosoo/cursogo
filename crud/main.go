package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	fmt.Println("Server rodando na porta 5k")
	log.Fatal(http.ListenAndServe(":5000", router))

}
