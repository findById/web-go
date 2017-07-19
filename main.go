package main

import (
	"flag"
	"net/http"

	"github.com/findById/web-go/api"
	"github.com/findById/web-go/gateway"
)

var (
	port = flag.String("port", "8080", "accept port")
)

func init() {
	parseConfig("config.json", &Config)
}

func main() {
	flag.Parse()
	if *port == "" {
		flag.PrintDefaults()
		return
	}
	router := []Router{
		Router{"GET", "/", IndexHandler},
		Router{"POST", "/api/v1", gateway.Handler},
		Router{"POST", "/api/message", api.MessageHandler},
	}

	mux := NewMux(router)
	http.ListenAndServe(Config.Host+":"+Config.Port, mux)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, 世界!"))
}
