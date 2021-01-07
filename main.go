package main

import (
	"log"
	"net/http"

	"methods"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {

	svc := methods.StringService{}

	// start counting the time between requests,
	// if the time interval was longer than the
	// defined, it persists the items contained
	// in queue
	go methods.Timer()

	PersistHandler := httptransport.NewServer(
		methods.MakePersistEndpoint(svc),
		methods.DecodePersistRequest,
		methods.EncodeResponse,
	)

	http.Handle("/persist", PersistHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
