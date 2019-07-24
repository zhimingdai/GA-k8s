package GA__final

import (
	"fmt"
	"math/rand"
	"sort"
)


func  (poptest Popluatin)Election( ) (Poplutationgene,float64,float64) {
fmt.Println("poptest的长度为：",len(poptest))
	var validitychrom []int
	var invaliditychrom []int
	//chrotest := GA__final.Initchrom()
	//fmt.Println("diyigefangan wei :",chrotest[1])
	//GA__final.Change(chrotest[1])
	//var pop1 Popluatin
	//poptest := pop1.Initchrom()  这个要在函数外面引用执行，然后调用poptest的election方法
	//poptest[1].Remainresour(nodelist)
	//fmt.Println(nodelist)

	//type pop2 map[float64]Nodelist
	//var popval = make(pop2)
	Chrombalancevali := make(Chrombalance)
	var Chrosominvali Poplutationgene
//这个函数循环所有的200个个体，判断出哪几个是可以用的，因为有些个体一开始就是错误的，那么后面也就没有必要浪费时间去进行计算的功能，直接交叉或者突变
	for i:=0;i< 200 ;i++  {
		//_,nodelist :=  Creatinit()  //这个也要在函数外面引用，然后当作参数应用
		//fmt.Println("测试测试：",i)
		poptest[i].Remainresour()
		if  poptest[i].Judegevalidityfirst(){
			//能够使用
			validitychrom = append(validitychrom,i)
			//singlebalan := poptest[i].Singlebalance(nodelist)
			//useragebalan := poptest[i].Useageblance(nodelist)
			clusterbalan :=poptest[i].Clusterbalance()
			//fmt.Println("这个是基因的表现型：",poptest[i])
			taskcomputetime := poptest[i].Taskcomputetime()
			fmt.Println("计算时间为1：",taskcomputetime)
			//taskcomputetime = taskcomputetime - rand.Float64()
			allelement := (clusterbalan + (10-taskcomputetime)) /2
			//allelement = allelement - rand.Float64()
			Chrombalancevali[allelement] =  poptest[i]
			//crossover
			//mutation
			//fmt.Println("第",i,"个个体的资源种类均衡度为：",singlebalan)
			//fmt.Println("第",i,"个个体的个体的使用均衡度为：",useragebalan)
			//fmt.Println("第",i,"个个体的个体均衡度为：",clusterbalan)
			//popval[clusterbalan] = nodelist

		}else{
			 //不能使用
			invaliditychrom = append(invaliditychrom,i)
			Chrosominvali = append(Chrosominvali,poptest[i].Gl)

		}
	}
 //Chrombalancevali代表能够使用的个体，然后调用下面的函数对其进行处理得到相应的子代
 fmt.Println("这个是第一种遗传算法")
	Newpopgene,result,besttime := Chrombalancevali.Chronsonvali(len(validitychrom))
	Newpopgene = Chronsoninvali(Newpopgene,len(invaliditychrom))
	//for i := range Newpopgene {
		//fmt.Println("xinde erzi ", Newpopgene[i])
	//}



	//for i := range Chrosominvali{
		//fmt.Println("zhe xie dou shi bu neng yong  de :",Chrosominvali[i])
	//}
    //fmt.Println(testslice)
	//var num []int
	//num = make([]int,10)
	//rand.Seed(time.Now().UnixNano())
    /*for i := 0;i <=50;i++{

    a,_ := Rws(testslice)
		switch a {
		case 0: num[0]++
		case 1: num[1]++
		case 2: num[2]++
		case 3: num[3]++
		case 4: num[4]++
		case 5: num[5]++
		case 6: num[6]++
		case 7: num[7]++
		case 8: num[8]++
		case 9: num[9]++




		}
    //fmt.Println(a,b)
    }*/
	//fmt.Println(num)



	//for i := range sortclubalan {
		//fmt.Println(sortclubalan[i])
		//fmt.Println(sortclubalan[i],Chbalansort[sortclubalan[i]])
	//}
	//fmt.Println(popval)
	fmt.Println(len(validitychrom ))
	return  Newpopgene,result,besttime
}
func  (poptest Popluatin)Electionsecond( ) (Poplutationgene,float64,float64) {
	var  validitychrom []int
	var invaliditychrom []int
	//Chrombalancevali := make(Chrombalance)
	Chrombalancevali222 := make(Chrombalance)
	var Chrosominvali Poplutationgene
	//这个函数循环所有的200个个体，判断出哪几个是可以用的，因为有些个体一开始就是错误的，那么后面也就没有必要浪费时间去进行计算的功能，直接交叉或者突变
	for i:=0;i< 200 ;i++  {

		poptest[i].Remainresour()
		noderemain, nodeuseout := poptest[i].Judegevaliditysecond()
		//fmt.Println("xixixiix",len(noderemain))
		//fmt.Println("xixixiix222",len(nodeuseout))
		if  len(nodeuseout) == 0 {   //nodeuseout为0代表所有的主机都是有空余资源的，因此这个就是可以使用的
			//能够使用
			validitychrom = append(validitychrom,i)

			clusterbalan :=poptest[i].Clusterbalance()

			//fmt.Println("均衡度为：",clusterbalan)
			//增加计算时间这个选择函数
			taskcomputetime := poptest[i].Taskcomputetime()
			 fmt.Println("计算时间为2：",taskcomputetime)
			allelement :=  (clusterbalan + (10-taskcomputetime) )/ 2
			//Chrombalancevali[clusterbalan] =  poptest[i]
			Chrombalancevali222[allelement] =  poptest[i]


		}else{
//就是不符合规范的子代，这个nodeuseout不为0，也就是代表有的主机是没有空余资源的，有的是有的，
//
			invaliditychrom = append(invaliditychrom,i)
			//主要的修改的地方，也就是把那些使用过度的节点上的任务分配到那些使用为0的节点上面去
			poptest[i].Gl =Chronsoninvalisecond(poptest[i].Gl,noderemain , nodeuseout)
			Chrosominvali = append(Chrosominvali,poptest[i].Gl)




		}
	}
	fmt.Println("second",len(validitychrom ))
	
	//Newpopgene := Chrombalancevali.Chronsonvali(len(validitychrom))
	fmt.Println("这是第二种遗传算法")
	Newpopgene,result,besttime := Chrombalancevali222.Chronsonvali(len(validitychrom))
	for i :=0;i<=len(Chrosominvali)/2 ;i++{
		Newpopgene = append(Newpopgene,Chrosominvali[i])
	}
	Newpopgene = Chronsoninvali(Newpopgene, 200-len(Newpopgene))
	fmt.Println("xinde zi qun de geshu ",len(Newpopgene))


	return  Newpopgene,result,besttime
}

func  Chronsoninvalisecond (gl Genelist, noderemain []int, nodeuseout[]int) Genelist{
          for i := 0;i<14;i++{


         	for j := range gl[i]{
         		for m := range nodeuseout{
         			if gl[i][j] == nodeuseout[m]{
						gl[i][j] = noderemain[rand.Intn(len(noderemain))]
					}
					break
				}


			}
		 }
		 return  gl
}

//这个是否可以重新进行修改
func  Chronsoninvali (Newpopgene Poplutationgene, b int) Poplutationgene{
	for i:=0;i< b ;i++{
		//	//var chro Ch
		var chro1 Genelist
		//chro1.genelist
		chro1  = make(Genelist,15)

		for i := 0;i< 14;i++{
			num := testmicro[i].msreq
			temp := msnodelist(num)
			chro1[i] = temp
		}
		Newpopgene = append(Newpopgene,chro1)



	}
	return Newpopgene
}

//根据原先的可以得到有效解的个数，得到相同数目的子代，子代中包括最好的，也包括后面交叉得到的子代
func(Chrombalancevali Chrombalance)  Chronsonvali(numvaliditychrom int ) (Poplutationgene,float64,float64){
	//fmt.Println("进来的时候的长度为：",numvaliditychrom)
	//for i := range Chrombalancevali{
	//	fmt.Println("每个个体对应的均衡度为：",i,Chrombalancevali[i])
	//}
	//var Chrombalanceson Chrombalance
	sortclubalan := Chrombalancevali.Sortbybalance()
	//fmt.Println("zheshi waimainde sort",sortclubalan)
	//for i := range sortclubalan{
		//fmt.Println(sortclubalan[i])
		//fmt.Println(Chrombalancevali[sortclubalan[i]])
	//}
	//fmt.Println(sortclubalan)

	//fmt.Println("paixu hou de ",sortclubalan)
	//fmt.Println("neng  shi yong de  changdu wei :",len(validitychrom))
	//这个就是新的子代的基因型
	//将最好的直接放入到下一代的种群当中,
	//应该记录下当前最好的基因的表现状态，然后化成图形
	fmt.Println("zhe chi zui hao de wei :",sortclubalan[len(sortclubalan)-1])
	//result得到最优的解并且返回，然后画图
	result := sortclubalan[len(sortclubalan)-1]
	Newpopgene,besttime  := Bestgene(sortclubalan,Chrombalancevali)
	sortclubalan = sortclubalan[:len(sortclubalan)-1]
	//进来的时候的个数为1
	if numvaliditychrom == 1{
		return Newpopgene,result,besttime
	}else if numvaliditychrom == 2{
		Newpopgene = append(Newpopgene,Chrombalancevali[sortclubalan[len(sortclubalan)-1]].Gl)
		//fmt.Println("当长度为2是：",len(Newpopgene))
		return Newpopgene,result,besttime
	}else{
		finalselecpop,selectionProbability := Firstfitness(sortclubalan,Chrombalancevali)
		for  len(Newpopgene) <= numvaliditychrom {
			chromosomeBa := finalselecpop[Rws(selectionProbability)].Gl
			//母亲染色体
			//fmt.Println("baba",chromosomeBa)
			chromosomeMa := finalselecpop[ Rws(selectionProbability)].Gl
			//fmt.Println("mama",chromosomeMa)
			//var popchromosomeSon  Poplutationgene

			son1, son2 := Crossover( chromosomeMa ,chromosomeBa )
			Newpopgene = append(Newpopgene, son1)
			Newpopgene = append(Newpopgene, son2)

		}
		Newpopgene1  := Newpopgene[:numvaliditychrom]


		return Newpopgene1,result,besttime
	}

	//防止只有一个个体能得到解的情况
	//if len(sortclubalan) == 1{
		//var temp []Genelist
		//temp = append(Newpopgene,Chrombalancevali[sortclubalan[len(sortclubalan)-1]].Gl)
		//return temp,result
	/*}else if len(sortclubalan) != 0 {
		//fmt.Println(sortclubalan)
	//	fmt.Println("zheg shi zui hao de ",Newpopgene)

		//删除了最好的元素的继续使用轮盘算子进行更新子代的工作
		finalselecpop,selectionProbability := Firstfitness(sortclubalan,Chrombalancevali)
		//for i := range finalselecpop{

		//	fmt.Println("zuihoude key wei :",i)
			//fmt.Println("zui hou de value wei:",finalselecpop[i])
		//}

		//循环让子代的个数和原来的一样
		for  len(Newpopgene) <= numvaliditychrom {
			chromosomeBa := finalselecpop[Rws(selectionProbability)].Gl
			//母亲染色体
			//fmt.Println("baba",chromosomeBa)
			chromosomeMa := finalselecpop[ Rws(selectionProbability)].Gl
			//fmt.Println("mama",chromosomeMa)
			//var popchromosomeSon  Poplutationgene

			son1, son2 := Crossover( chromosomeMa ,chromosomeBa )
			Newpopgene = append(Newpopgene, son1)
			Newpopgene = append(Newpopgene, son2)

		}
		Newpopgene1  := Newpopgene[:numvaliditychrom]


		return Newpopgene1,result
	}else {
		//fmt.Println("新的子代的长度为：",len(Newpopgene))
		Newpopgene1 := Newpopgene[:numvaliditychrom]

		return Newpopgene1, result
	}*/


}
//从有效的种群中选出一个最优的个体
func Bestgene(sortclubalan []float64,Chrombalancevali Chrombalance) ([]Genelist,float64){

	var Newpopgene []Genelist
	Newpopgene = append(Newpopgene,Chrombalancevali[sortclubalan[len(sortclubalan)-1]].Gl)
	temp := Chrombalancevali[sortclubalan[len(sortclubalan)-1]]
	result2 :=  temp.Taskcomputetime()
	fmt.Println("最好的计算时间为：",result2)
	bestnl := Chrombalancevali[sortclubalan[len(sortclubalan)-1]].Nl
	var usednode int
	for j := range bestnl{
		if bestnl[j].Cpund != 4 || bestnl[j].Disknd != 100{
			usednode ++
		}
	}
	fmt.Println("使用的节点的个数为：",usednode)
	delete(Chrombalancevali,sortclubalan[len(sortclubalan)-1])
	//sortclubalan = sortclubalan[:len(sortclubalan)-1]
	//fmt.Println(sortclubalan)
	return Newpopgene,result2
}





func Crossover(chromosomeMa,chromosomeBa Genelist) (Genelist,Genelist){
	var chromosomeSongene1 Genelist
	var chromosomeSongene2 Genelist
	for i := 0;i <14;i++ {
		num := func() int {
			var temp int
			if len(chromosomeMa[i]) >= len(chromosomeBa[i]) {
				temp = len(chromosomeBa[i])
			} else {
				temp = len(chromosomeMa[i])
			}
			return temp

		}()
		index := rand.Intn(num)
		//fmt.Println(num)
		//fmt.Println(index)

       var gene Gene
		gene  = append(gene, chromosomeBa[i][:index]...)
		gene = append(gene, chromosomeMa[i][index:]...)
		chromosomeSongene1 = append(chromosomeSongene1,gene)
		var gene2 Gene
		gene2 = append(gene2, chromosomeMa[i][:index]...)
		gene2 = append(gene2, chromosomeBa[i][index:]...)
		chromosomeSongene2 = append(chromosomeSongene2,gene2)
	}
		//chromosomeSon = append(chromosomeSon,chromosomeSongene1)
		//chromosomeSon = append(chromosomeSon,chromosomeSongene2)
//fmt.Println(chromosomeSongene2)
//	fmt.Println(chromosomeSongene1)
//popchromosomeSon = append(popchromosomeSon,chromosomeSongene1)
//popchromosomeSon = append(popchromosomeSon,chromosomeSongene2)
	//for  i := range popchromosomeSon  {
	//	fmt.Println(popchromosomeSon[i])
	//}


return  chromosomeSongene1,chromosomeSongene2

}



//计算第一个个选择算子后的适应度大小
func Firstfitness(sortclubalan []float64,Chbalansort Chrombalance) (Chrombalance,[]float64){
   fmt.Println("sortclubalan :",sortclubalan)
	diff := (sortclubalan[len(sortclubalan)-1] - sortclubalan[0]) / 3
	fmt.Println("diff :",diff)
	//fmt.Println(diff)
	var firsrpop []float64
	var secondpop []float64
	var thirdpop []float64
	for i := range Chbalansort {
		if i >= sortclubalan[0] && i < sortclubalan[0]+diff {
			thirdpop = append(thirdpop, i)
		} else if i >= sortclubalan[0]+diff && i < sortclubalan[0]+2*diff {
			secondpop = append(secondpop, i)
		} else {
			firsrpop = append(firsrpop, i)
		}
	}
	avefitnessfirst := func(pop []float64) float64 {
		if len(pop) == 0{
			return 0
		}else {
			var sum float64
			for i := range pop {
				sum = sum + pop[i]
			}
			sum = sum / float64(len(pop))
			return sum
		}
	}(firsrpop)
	avefitnesssecond := func(pop []float64) float64 {
		if len(pop) == 0{
			return 0
		}else {
			var sum float64
			for i := range pop {
				sum = sum + pop[i]
			}
			sum = sum / float64(len(pop))
			return sum
		}
	}(secondpop)
	avefitnessthird := func(pop []float64) float64 {
		if len(pop) == 0{
			return 0
		}else{
		var sum float64
		for i := range pop {
			sum = sum + pop[i]
		}
		sum = sum / float64(len(pop))
		return sum
		}
	}(thirdpop)
	//fmt.Println(avefitnessfirst)
	//fmt.Println(avefitnesssecond)
	//fmt.Println(avefitnessthird)
	//fmt.Println(firsrpop)
	//fmt.Println(secondpop)
	//fmt.Println(thirdpop)
	avefitnessall := avefitnessfirst + avefitnesssecond + avefitnessthird
	//fmt.Println("11111aaaaaaaaaaa:= ",avefitnessall)
	//selectionprobality := []float64{avefitnessfirst / avefitnessall, avefitnesssecond / avefitnessall, avefitnessthird / avefitnessall}
	var finalchbalance Chrombalance
	finalchbalance = make(Chrombalance)
	var finalselecproblity []float64
	for i := range firsrpop{

		//Chbalansort[firsrpop[i]]
		finalselecproblityfirst  :=  (firsrpop[i]  /avefitnessall ) / float64(len(firsrpop))
		finalchbalance[finalselecproblityfirst] = Chbalansort[firsrpop[i]]
		finalselecproblity = append(finalselecproblity,finalselecproblityfirst)
	}
	//fmt.Println("11111aaaaaaaaaaa:= ",finalselecproblity)
	for i := range secondpop{

		//Chbalansort[firsrpop[i]]
		finalselecproblitysecond  :=  (secondpop[i]  /avefitnessall ) / float64(len(secondpop))
		//fmt.Println("333222222aaaaaaaaaaa:= ",finalselecproblitysecond )
		finalchbalance[finalselecproblitysecond] = Chbalansort[secondpop[i]]
		finalselecproblity = append(finalselecproblity,finalselecproblitysecond)
	}
	//fmt.Println("222222aaaaaaaaaaa:= ",finalselecproblity)
	for i := range thirdpop{

		//Chbalansort[firsrpop[i]]
		finalselecproblitythird  :=  (thirdpop[i]  /avefitnessall ) / float64(len(thirdpop))
		finalchbalance[finalselecproblitythird] = Chbalansort[thirdpop[i]]
		finalselecproblity = append(finalselecproblity,finalselecproblitythird)
	}
	//fmt.Println("33333aaaaaaaaaaa:= ",finalselecproblity)
//fmt.Println("aaaaaaaaaaa:= ",finalselecproblity)
    return   finalchbalance,finalselecproblity
}


func Rws(selectionProbability []float64) float64{
	sum := 0.0
	//rand.Seed(time.Now().UnixNano())
   //fmt.Println("xuanzhe de = :",selectionProbability)
	r := rand.Float64()
	//fmt.Println("r = :",r)
	var index  int
	for index < len(selectionProbability) {
		sum += selectionProbability[index]
		if sum >= r {
			break
		}
		index++
	}
	return selectionProbability[index]
}


//选择函数，按照适应度函数的大小进行排序，然后返回相应的map和对应的slice
func (Chblan Chrombalance)Sortbybalance () []float64{
	var sortclubalan []float64
	//var sumclubalance float64
	//var Chbalansort Chrombalance
	//Chbalansort = make(Chrombalance)
	//var num int
	//var num1 int
	//for i := range Chblan{
		//fmt.Println("每个个体对应的均衡度为：",i,Chrombalancetest[i])
		//num ++
		//sumclubalance = sumclubalance + i

	//}

	for i  := range Chblan{
		//i = i / sumclubalance
		//num1 ++
		//fmt.Println("i wei :",i)
		//fmt.Println("一共有：",num)
		sortclubalan = append(sortclubalan,i)

		//Chbalansort[i] = j
	}

	//fmt.Println("排序前的选择函数大小为：",Chblan)
	 sort.Float64s(sortclubalan)
	//fmt.Println("排序hou的悬着函数大小为：",Chbalansort)
	//sort.Float64s(sortclubalan)

	return sortclubalan   //这个就是排好序的各个的

}



	

