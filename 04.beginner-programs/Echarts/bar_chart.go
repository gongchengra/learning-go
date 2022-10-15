package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"os"
)

func main() {
	medals := []opts.BarData{{Value: 46},
		{Value: 38}, {Value: 29},
		{Value: 22}, {Value: 13}, {Value: 11}}
	countries := []string{"USA", "China", "UK", "Russia",
		"South Korea", "Germany"}
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{Title: "Olympic Gold medals in London"}))
	bar.SetXAxis(countries)
	bar.AddSeries("medals", medals)
	f, _ := os.Create("bar_chart.html")
	bar.Render(f)
}
