package main

import (
	"net/http"
	"flag"
	"web-go/message"
	"web-go/gateway"
)

var (
	port = flag.String("port", "8080", "accept port")
)

func main() {
	flag.Parse()
	if *port == "" {
		flag.PrintDefaults()
		return
	}
	router := []Router{
		Router{"GET", "/", IndexHandler},
		Router{"POST", "/api/v1", gateway.Handler},
		Router{"POST", "/message", message.Handler},
	}

	mux := NewMux(router)
	http.ListenAndServe(":" + *port, mux)
}
