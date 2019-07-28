package main

import (
	"GA--final"
	"fmt"
	"gonum.org/v1/plot/plotter"
	_ "gonum.org/v1/plot/plotter"
	"math/rand"
	"time"
	// "time"
)

func main() {

	var pop1 GA__final.Popluatin
	poptest1 := pop1.Initchrom()
	poptest2 := pop1.Initchrom()
	poptest3 := pop1.Initchrom()

	rand.Seed(time.Now().UnixNano())
	//poptest1  := poptest1
	//poptest2 := poptest2
	var dataslicefirst plotter.XYs
	var dataslicesecond plotter.XYs
	var dataslicethird plotter.XYs
	var dataslice3 plotter.XYs
	var dataslice4 plotter.XYs
	var dataslice5 plotter.XYs
	for i := 0; i <= 4000; i++ {

		fmt.Print("zhe shi di ", i)
		popson, result1, besttimt1 := poptest1.Election()
		//var datafirst plotter.XY
		datafirst := plotter.XY{
			float64((i)),
			result1}
		dataslicefirst = append(dataslicefirst, datafirst)
		data3 := plotter.XY{
			float64((i)),
			besttimt1}
		dataslice3 = append(dataslice3, data3)

		poptest1 = popson.Soninit()
		popson222, result2, besttime2 := poptest2.Electionsecond()
		datasecond := plotter.XY{
			float64((i)),
			result2}
		dataslicesecond = append(dataslicesecond, datasecond)
		data4 := plotter.XY{
			float64(i),
			besttime2}
		dataslice4 = append(dataslice4, data4)
		poptest2 = popson222.Soninit()

		popson3, result3, besttime3 := poptest3.GeneralElection()
		datathird := plotter.XY{
			float64((i)),
			result3}
		dataslicethird = append(dataslicethird, datathird)
		data5 := plotter.XY{
			float64(i),
			besttime3}
		dataslice5 = append(dataslice5, data5)
		poptest3 = popson3.Soninit()
	}
	GA__final.Plot(dataslicefirst, dataslicesecond, dataslicethird, dataslice3, dataslice4, dataslice5)

}
