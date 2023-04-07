package server

import (
	"diplom/generator"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
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

var saver = opts.ToolBoxFeatureSaveAsImage{
	Show:  true,
	Type:  "png",
	Name:  "baza",
	Title: "huj",
}

var feature = opts.ToolBoxFeature{
	SaveAsImage: &saver,
	DataZoom:    nil,
	DataView:    nil,
	Restore:     nil,
}

var pointer = opts.AxisPointer{
	Type: "cross",
	Snap: true,
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
		charts.WithInitializationOpts(opts.Initialization{Height: "900px", Width: "900px"}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:        true,
			Trigger:     "item",
			TriggerOn:   "mousemove",
			Formatter:   "point",
			AxisPointer: &pointer,
		}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:    false,
			Orient:  "",
			Left:    "",
			Top:     "",
			Right:   "",
			Bottom:  "",
			Feature: &feature,
		}),
	)

	// Put data into instance
	line.SetXAxis(X).
		AddSeries("Category A", Y1, charts.WithLineChartOpts(opts.LineChart{Smooth: false})).
		AddSeries("Category B", Y2, charts.WithLineChartOpts(opts.LineChart{Smooth: false}))
	line.Render(w)
}

func httpserver_delta(w http.ResponseWriter, _ *http.Request) {
	number_of_points := 20
	h := 2 * math.Pi / float64(number_of_points)
	// create a new line instance
	X := make([]float64, 0)
	Y1 := make([]opts.LineData, 0)
	Y2 := make([]opts.LineData, 0)

	j := 0
	for i := -math.Pi; i <= math.Pi; i += h {
		switch j {
		case 1:
			X = append(X, i)
			Y1 = append(Y1, opts.LineData{Value: float32(generator.RungeSimpUp(X[j], number_of_points))})
			Y2 = append(Y2, opts.LineData{Value: float32(generator.RungeTrapUp(X[j], number_of_points))})
		case 2:
			X = append(X, i)
			Y1 = append(Y1, opts.LineData{Value: float32(generator.RungeSimpUp(X[j], number_of_points))})
			Y2 = append(Y2, opts.LineData{Value: float32(generator.RungeTrapUp(X[j], number_of_points))})
		default:
			X = append(X, i)
			Y1 = append(Y1, opts.LineData{Value: float32(generator.RungeSimp(X[j], number_of_points))})
			Y2 = append(Y2, opts.LineData{Value: float32(generator.RungeTrap(X[j], number_of_points))})
		}
		j++
	}
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Height: "900px", Width: "900px"}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:        true,
			Trigger:     "item",
			TriggerOn:   "mousemove",
			Formatter:   "point",
			AxisPointer: &pointer,
		}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:    false,
			Orient:  "",
			Left:    "",
			Top:     "",
			Right:   "",
			Bottom:  "",
			Feature: &feature,
		}),
	)

	// Put data into instance
	line.SetXAxis(X).
		AddSeries("Category A", Y1, charts.WithLineChartOpts(opts.LineChart{Smooth: false})).
		AddSeries("Category B", Y2, charts.WithLineChartOpts(opts.LineChart{Smooth: false}))
	line.Render(w)
}
