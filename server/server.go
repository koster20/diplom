package server

import (
	"net/http"
)

func ShowGraph() {
	http.HandleFunc("/", httpserver)
	http.HandleFunc("/delta", httpserver_delta)
	http.ListenAndServe(":8081", nil)
}
