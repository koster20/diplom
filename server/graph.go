package server

import (
	"diplom/generator"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"math"
	"math/rand"
	"net/http"
)

func generateLineItems() []opts.LineData {

	items := make([]opts.LineData, 0)

	for i := 0; i < 7; i++ {
		generator.Base_function(0.5)
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	number_of_points := 20
	h := 2 * math.Pi / float64(number_of_points)
	// create a new line instance
	X := make([]float64, 0)
	Y1 := make([]opts.LineData, 0)
	Y2 := make([]opts.LineData, 0)
	j := 0
	for i := -math.Pi; i <= math.Pi; i += h {
		X = append(X, i)
		Y1 = append(Y1, opts.LineData{Value: generator.RungeSimp(X[j], number_of_points)})
		Y2 = append(Y2, opts.LineData{Value: generator.RungeTrap(X[j], number_of_points)})
		j++
	}
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		}))

	// Put data into instance
	line.SetXAxis(X).
		AddSeries("Category A", Y1).
		AddSeries("Category B", Y2).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}
