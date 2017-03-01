package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func db() {
	sleep := rand.Intn(1000)
	fmt.Printf("\n\tDB Access (%dms)\n", sleep)
	time.Sleep(time.Duration(sleep) * time.Millisecond)
}

func async() {
	sleep := rand.Intn(1000)
	fmt.Printf("\n\tAsync job (%dms)\n", sleep)
	time.Sleep(time.Duration(sleep) * time.Millisecond)
}

func PostAuth(w http.ResponseWriter, req *http.Request) {
	sleep := rand.Intn(1000) + 1000
	fmt.Printf("%s %s (%dms) ...", req.Method, req.URL, sleep)
	time.Sleep(time.Duration(sleep) * time.Millisecond)

	db()
	go async()

	fmt.Println("OK")
}

func main() {
	var bind = ":8002"

	m := mux.NewRouter()
	m.HandleFunc("/auth/{id}", PostAuth).
		Methods("POST")

	log.Printf("bind = %+v\n", bind)
	log.Fatalln(
		http.ListenAndServe(bind, m),
	)
}
