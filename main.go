package main

import (
	"net/http"
	"flag"
	"web-go/message"
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
		Router{"POST", "/message", message.MessageHandler},
	}

	mux := NewMux(router)
	http.ListenAndServe(":" + *port, mux)
}
