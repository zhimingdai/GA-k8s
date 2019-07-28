package GA__final

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func Plot(data1 plotter.XYs, data2 plotter.XYs, datagernal plotter.XYs, data3 plotter.XYs, data4 plotter.XYs, data5 plotter.XYs) {
	rand.Seed(int64(0))

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "迭代的次数"
	p.Y.Label.Text = "得到的结果"

	err = plotutil.AddLinePoints(p,
		"First", data1,
		"Second", data2,
		"third", datagernal,
		"3", data3,
		"4", data4,
		"5", data5)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

// randomPoints returns some random x, y points.
