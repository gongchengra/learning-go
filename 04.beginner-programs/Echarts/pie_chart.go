package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"os"
)

func main() {
	destinations := []opts.PieData{{Name: "Croatia", Value: 22},
		{Name: "Bohemia", Value: 34}, {Name: "Bulgaria", Value: 18},
		{Name: "Spain", Value: 5}, {Name: "Others", Value: 21}}
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeChalk}),
		charts.WithTitleOpts(opts.Title{Title: "Popular destinations"}),
	)
	pie.AddSeries("destinations", destinations)
	f, _ := os.Create("pie_chart.html")
	pie.Render(f)
}
