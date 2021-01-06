package main

import (
	"container/list"
	"log"
	"net/http"

	"methods"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {

	//TODO
	//it is possible run RPC concurrently to this server
	//using goroutines and channels
	//I will define how to implement it later

	//init queue
	queue := list.New()

	svc := methods.StringService{}

	PersistHandler := httptransport.NewServer(
		methods.MakePersistEndpoint(svc, queue),
		methods.DecodePersistRequest,
		methods.EncodeResponse,
	)

	http.Handle("/persist", PersistHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
