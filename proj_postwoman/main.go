package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println(" = = = = = postwoman run server = = = = =")
	router := mux.NewRouter()
	router.HandleFunc("/", GetIndexEndpoint).Methods("GET")
	router.HandleFunc("/api", GetAPIEndpoint).Methods("GET")

	http.ListenAndServe(":8080", router)
}

func GetIndexEndpoint(resp http.ResponseWriter, req *http.Request) {
	message := "CALLED~ GetIndexEndpoint"
	log.Println(message)

	resp.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(resp, "<script>alert(\"move /api\");window.location.pathname=\"/api\"</script>")
}

func GetAPIEndpoint(resp http.ResponseWriter, req *http.Request) {
	message := "CALLED~ GetAPIEndpoint"
	log.Println(message)

	resp.Header().Set("Content-Type", "text/html; charset=utf-8")
	//resp.Write([]byte(`{"message": "` + message + `"}`))

	fmt.Fprintln(resp, "<h1>"+message+"</h1>")
}
