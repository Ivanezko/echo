package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var router *mux.Router

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Print("init...")
	rand.Seed(time.Now().UnixNano())
	log.Print("init...done")
}

func main() {
	log.Print("main...")

	router = mux.NewRouter()

	router.PathPrefix("/").HandlerFunc(echo)
	router.HandleFunc("/sys-live", live).Methods("GET")
	router.HandleFunc("/sys-ready", live).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{},
		Debug:            false,
	})

	srv := "0.0.0.0:3000"
	log.Println("Server listen on: " + srv)
	handler := c.Handler(router)
	err := http.ListenAndServe(srv, handler)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("main...done")
}

func echo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	log.Print("Requested: " + r.RequestURI)

	fmt.Fprintf(w, "your request PROD12:\n <pre>%+v</pre>\n", spew.Sdump(r))

}

func live(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", "OK")
}
