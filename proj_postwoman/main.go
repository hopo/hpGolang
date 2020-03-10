package main

import (
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

// Todo: Database
// Todo: Sessions
// Todo: Crypt
// Todo: Logger
// Todo: Auth - JWT
// Todo: Validate - go-play
// Todo: Middleware - authrequired

// set templates
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	log.Println(" = = = = = POSTWOMAN Run Server = = = = =")

	router := mux.NewRouter()

	// Todo: how to use allowed
	headers := gohandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := gohandlers.AllowedMethods([]string{"GET", "POST"})
	origins := gohandlers.AllowedOrigins([]string{"127.0.0.1"})

	router.HandleFunc("/", GetIndexEndpoint).Methods("GET")
	router.HandleFunc("/api", GetAPIEndpoint).Methods("GET")

	ch := gohandlers.CORS(headers, methods, origins)

	// set fileserver
	fs := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.ListenAndServe(":9090", ch(router))
}

func GetIndexEndpoint(res http.ResponseWriter, req *http.Request) {
	message := "CALLED~ GetIndexEndpoint"
	log.Println(message)

	err := tpl.ExecuteTemplate(res, "index.html", nil)
	if err != nil {
		log.Panic(err)
	}

	// >case1
	//res.Header().Set("Content-Type", "text/html; charset=utf-8")
	// >case2
	//fmt.Fprintln(res, "<script>alert(\"move /api\");window.location.pathname=\"/api\"</script>")
}

func GetAPIEndpoint(res http.ResponseWriter, req *http.Request) {
	message := "CALLED~ GetAPIEndpoint"
	log.Println(message)

	res.Header().Set("Content-Type", "application/json;")
	res.Write([]byte(`{"message": "` + message + `"}`))
	//res.Header().Set("Content-Type", "text/html; charset=utf-8")
	////fmt.Fprintln(res, "<h1>"+message+"</h1>")
}
