package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	// Route: /books/{title} (GET) ReadBook in handlers.go
    route.HandleFunc("/books/{title}", ReadBook).Methods("GET")

    route.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })

	// Route: / - list all routes
	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	    var routes []string
	    route.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	        path, _ := route.GetPathTemplate()
	        methods, _ := route.GetMethods()
	        routeStr := fmt.Sprintf("%s %s", methods, path)
	        routes = append(routes, routeStr)
	        return nil
	    })
	
	    w.Header().Set("Content-Type", "text/html; charset=utf-8")
	    fmt.Fprintf(w, "<h1>Registrerade routes</h1><ul>")
	    for _, route := range routes {
	        fmt.Fprintf(w, `<li><a href="%s">%s</a></li>`, strings.Split(route, " ")[1], route)
	    }
		fmt.Fprintf(w, "<li><a href=\"/static/\">/static</a></li>")
	    fmt.Fprintf(w, "</ul>")
	})   

	// Route: /static - file server
	fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	port := "8888"
    fmt.Printf("Servern körs på http://localhost:%s\n", port)
    http.ListenAndServe(":" + port, route)
}