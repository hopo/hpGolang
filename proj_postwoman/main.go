package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("= = = postwoman run server = = =")
	router := mux.NewRouter()
	router.HandleFunc("/api", GetAPIEndpoint).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func GetAPIEndpoint(resp http.ResponseWriter, req *http.Request) {
	message := "CALLED~ GetAPIEndpoint"
	log.Println(message)
	resp.Write([]byte(`{"message": "` + message + `"}`))
}
