package main

import (
	"math/rand"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/opts"
)

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func main() {
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)
}
