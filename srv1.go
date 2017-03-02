//+build srv1

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	dd "github.com/gchaincl/dd-go-opentracing"
	"github.com/gorilla/mux"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func init() {
	tracer := dd.NewTracer()
	opentracing.SetGlobalTracer(tracer)
}

func trace(op string, parent opentracing.Span, req *http.Request) opentracing.Span {
	span := opentracing.StartSpan(op, opentracing.ChildOf(parent.Context()))
	ext.Component.Set(span, "/auth/{id}")
	err := span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	if err != nil {
		log.Fatalln(err)
	}

	return span
}

func PostUser(w http.ResponseWriter, req *http.Request) {
	span := opentracing.StartSpan("Handle POST")
	ext.Component.Set(span, "/users/{id}")
	ext.PeerService.Set(span, "srv1")
	dd.EnvTag.Set(span, "test")
	defer span.Finish()

	sleep := rand.Intn(1000)
	fmt.Printf("%s %s (%dms) ...", req.Method, req.URL, sleep)
	time.Sleep(time.Duration(sleep) * time.Millisecond)

	req, _ = http.NewRequest("POST", "http://localhost:8002/auth/"+mux.Vars(req)["id"], nil)
	child := trace("Call POST", span, req)
	http.DefaultClient.Do(req)
	child.Finish()

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
