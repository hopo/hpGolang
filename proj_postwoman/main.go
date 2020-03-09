package main

import (
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	log.Println(" = = = = = POSTWOMAN Run Server = = = = =")

	router := mux.NewRouter()

	// todo: how to use allowed
	headers := gohandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := gohandlers.AllowedMethods([]string{"GET", "POST"})
	origins := gohandlers.AllowedOrigins([]string{"127.0.0.1"})

	router.HandleFunc("/", GetIndexEndpoint).Methods("GET")
	router.HandleFunc("/api", GetAPIEndpoint).Methods("GET")

	ch := gohandlers.CORS(headers, methods, origins)

	http.ListenAndServe(":8080", ch(router))
}

func GetIndexEndpoint(resp http.ResponseWriter, req *http.Request) {
	message := "CALLED~ GetIndexEndpoint"
	log.Println(message)

	err := tpl.ExecuteTemplate(resp, "index.gohtml", nil)
	if err != nil {
		log.Panic(err)
	}

	// >case1
	//resp.Header().Set("Content-Type", "text/html; charset=utf-8")
	// >case2
	//fmt.Fprintln(resp, "<script>alert(\"move /api\");window.location.pathname=\"/api\"</script>")
}

func GetAPIEndpoint(resp http.ResponseWriter, req *http.Request) {
	message := "CALLED~ GetAPIEndpoint"
	log.Println(message)

	resp.Header().Set("Content-Type", "application/json;")
	resp.Write([]byte(`{"message": "` + message + `"}`))
	//resp.Header().Set("Content-Type", "text/html; charset=utf-8")
	////fmt.Fprintln(resp, "<h1>"+message+"</h1>")
}
