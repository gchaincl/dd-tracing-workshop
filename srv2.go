package main

import (
	"errors"
	"fmt"
	stdlog "log"
	"math/rand"
	"net/http"
	"time"

	dd "github.com/gchaincl/dd-go-opentracing"
	"github.com/gorilla/mux"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func init() {
	tracer := dd.NewTracer()
	opentracing.SetGlobalTracer(tracer)
}

func db(ctx opentracing.SpanContext) {
	defer opentracing.StartSpan("SELECT * FROM auth",
		opentracing.ChildOf(ctx),
	).Finish()

	sleep := rand.Intn(1000)
	fmt.Printf("\n\tDB Access (%dms)\n", sleep)
	time.Sleep(time.Duration(sleep) * time.Millisecond)
}

func async(ctx opentracing.SpanContext) {
	span := opentracing.StartSpan("ASYNC JOB",
		opentracing.ChildOf(ctx),
	)
	defer span.Finish()

	sleep := rand.Intn(3000)
	fmt.Printf("\n\tAsync job (%dms) ... ", sleep)
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	span.LogFields(log.Error(errors.New("boom")))
	fmt.Println("OK")
}

func PostAuth(w http.ResponseWriter, req *http.Request) {
	spanCtx, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)
	if err != nil {
		stdlog.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	sleep := rand.Intn(1000)
	fmt.Printf("%s %s (%dms) ...", req.Method, req.URL, sleep)
	time.Sleep(time.Duration(sleep) * time.Millisecond)

	db(spanCtx)
	go async(spanCtx)

	fmt.Println("OK")
}

func main() {
	var bind = ":8002"

	m := mux.NewRouter()
	m.HandleFunc("/auth/{id}", PostAuth).
		Methods("POST")

	stdlog.Printf("bind = %+v\n", bind)
	stdlog.Fatalln(
		http.ListenAndServe(bind, m),
	)
}
