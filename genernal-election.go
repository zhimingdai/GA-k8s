package GA__final

import "fmt"

func (poptest Popluatin) GeneralElection() (Poplutationgene, float64, float64) {
	fmt.Println("poptest的长度为：", len(poptest))
	var validitychrom []int
	var invaliditychrom []int
	Chrombalancevali := make(Chrombalance)
	var Chrosominvali Poplutationgene
	//这个函数循环所有的200个个体，判断出哪几个是可以用的，因为有些个体一开始就是错误的，那么后面也就没有必要浪费时间去进行计算的功能，直接交叉或者突变
	for i := 0; i < 200; i++ {
		poptest[i].Remainresour()
		if poptest[i].Judegevalidityfirst() {
			validitychrom = append(validitychrom, i)
			clusterbalan := poptest[i].Clusterbalance()
			taskcomputetime := poptest[i].Taskcomputetime()
			fmt.Println("计算时间为1：", taskcomputetime)
			allelement := (clusterbalan + (10 - taskcomputetime)) / 2
			Chrombalancevali[allelement] = poptest[i]
		} else {
			invaliditychrom = append(invaliditychrom, i)
			Chrosominvali = append(Chrosominvali, poptest[i].Gl)

		}
	}
	fmt.Println("这个是最开始的遗传算法")
	//应该对这里进行修改
	Newpopgene, result, besttime := Chrombalancevali.Chronsonvaligeneral(len(validitychrom))
	Newpopgene = Chronsoninvali(Newpopgene, len(invaliditychrom))
	fmt.Println(len(validitychrom))
	return Newpopgene, result, besttime
}
func (Chrombalancevali Chrombalance) Chronsonvaligeneral(numvaliditychrom int) (Poplutationgene, float64, float64) {

	sortclubalan := Chrombalancevali.Sortbybalance()

	fmt.Println("zhe chi zui hao de wei :", sortclubalan[len(sortclubalan)-1])

	result := sortclubalan[len(sortclubalan)-1]
	Newpopgene, besttime := Bestgene(sortclubalan, Chrombalancevali)
	sortclubalan = sortclubalan[:len(sortclubalan)-1]

	if numvaliditychrom == 1 {
		return Newpopgene, result, besttime
	} else if numvaliditychrom == 2 {
		Newpopgene = append(Newpopgene, Chrombalancevali[sortclubalan[len(sortclubalan)-1]].Gl)

		return Newpopgene, result, besttime
	} else {
		finalselecpop, selectionProbability := Firstfitnessgenernal(sortclubalan, Chrombalancevali)
		for len(Newpopgene) <= numvaliditychrom {
			chromosomeBa := finalselecpop[Rws(selectionProbability)].Gl

			chromosomeMa := finalselecpop[Rws(selectionProbability)].Gl

			son1, son2 := Crossover(chromosomeMa, chromosomeBa)
			Newpopgene = append(Newpopgene, son1)
			Newpopgene = append(Newpopgene, son2)

		}
		Newpopgene1 := Newpopgene[:numvaliditychrom]

		return Newpopgene1, result, besttime
	}

}
func Firstfitnessgenernal(sortclubalan []float64, Chbalansort Chrombalance) (Chrombalance, []float64) {
	var sum float64
	for i := range sortclubalan {
		sum = sum + sortclubalan[i]
	}
	var finalselecproblity []float64
	var finalchbalance Chrombalance
	finalchbalance = make(Chrombalance)
	for i := range sortclubalan {
		finalselecproblity = append(finalselecproblity, sortclubalan[i]/sum)
		finalchbalance[sortclubalan[i]/sum] = Chbalansort[sortclubalan[i]]
	}

	return finalchbalance, finalselecproblity
}
