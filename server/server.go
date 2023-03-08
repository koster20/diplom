package server

import (
	"net/http"
)

func ShowGraph() {
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)
}
