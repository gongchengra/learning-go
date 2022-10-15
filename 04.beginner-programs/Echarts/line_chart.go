package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"os"
)

func main() {
	sals := []opts.LineData{{Value: 567}, {Value: 612}, {Value: 800},
		{Value: 980}, {Value: 1410}, {Value: 2350}}
	ages := []string{"18", "20", "25", "30", "40", "50"}
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{Title: "Average salary per age",
			Subtitle: "Slovakia"}),
	)
	line.SetXAxis(ages).AddSeries("salaries", sals)
	line.SetSeriesOptions(charts.WithLineChartOpts(
		opts.LineChart{
			Smooth: true,
		}))
	f, _ := os.Create("line_chart.html")
	line.Render(f)
}
