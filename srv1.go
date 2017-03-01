package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func PostUser(w http.ResponseWriter, req *http.Request) {
	sleep := rand.Intn(1000) + 1000
	fmt.Printf("%s %s (%dms) ...", req.Method, req.URL, sleep)
	time.Sleep(time.Duration(sleep) * time.Millisecond)

	http.Post("http://localhost:8002/auth/"+mux.Vars(req)["id"], "", nil)

	fmt.Println(" OK")
}

func main() {
	var bind = ":8001"

	m := mux.NewRouter()
	m.HandleFunc("/users/{id}", PostUser).
		Methods("POST")

	log.Printf("bind = %+v\n", bind)
	log.Fatalln(
		http.ListenAndServe(bind, m),
	)
}
