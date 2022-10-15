package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"os"
)

func main() {
	temps := []opts.ScatterData{{Value: -7.3}, {Value: -3.4}, {Value: -5.0},
		{Value: -0.9}, {Value: -2.2}, {Value: 4.8}, {Value: 5.1}, {Value: -1.9},
		{Value: 0}, {Value: 2.6}}
	dates := []string{"Jan 1", "Jan 10", "Jan 12", "Jan 20", "Jan 30", "Feb 1",
		"Feb 2", "Feb 5", "Feb 8", "Feb 12"}
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{Title: "Temperatures"}),
	)
	scatter.SetXAxis(dates).AddSeries("temps", temps)
	f, _ := os.Create("scatter_chart.html")
	scatter.Render(f)
}
