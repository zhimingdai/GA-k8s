package GA__final

import (
	"sort"
)

/*func (chr *Chromosome) execcomputetime() {
	nl := chr.nl
	gl := chr.gl
	var sonsvctime float64
	for i := 0; i < 14; i++ {
		for j := range gl[i] {
			testmicro[i].Selfcomputetime(nl[gl[i][j]])
			if len(testmap[testmicro[i]]) == 0 {
				sonsvctime = 0
			}else{
				for _,sonsvc := range testmap[testmicro[i]]{
					testmicro[sonsvc - 1].Selfcomputetime()
				}
				sonsvctime =
			}

		}

	}
}*/

//这个计算的是所有14个任务的计算时间，也就是一整个服务的计算时间，应该是最大的时间
func (char *Chromosome) Taskcomputetime() float64 {
	var taskcomputetime []float64
	nl := char.Nl
	gl := char.Gl
	for i := 0; i < 14; i++ {
		taskcomputetime = append(taskcomputetime, Computetime(nl, gl, i))
		// fmt.Println("这是第",i,"个任务的计算时间",taskcomputetime)
	}
	sort.Float64s(taskcomputetime)
	return taskcomputetime[len(taskcomputetime)-1]

}

func Selfcomputetime(nl Nodelist, svcnode int, svc int) float64 {
	//fmt.Println("cpu使用了的资源为：",nl[svcnode].cpund)
	//nl[svcnode].cpund代表剩余的cpu
	selfcomputetime := testmicro[svc].cpures / ((testmicro[svc].cpures / (100 - nl[svcnode].Cpund)) * 100)
	return selfcomputetime
}

//i类任务的完成时间等于所有子任务中的最大的完成时间，子任务的完成时间等于所有这个类型的子任务完成时间中最大的一个
//忽略子任务传输给父任务的过程，假设子任务完成后，传输一个信息给父节点后，父节点就可以运行了
//这个计算的是单个任务的计算时间
func Computetime(nl Nodelist, gl Genelist, i int) float64 {
	//fmt.Println("这是第",i,"个任务", "这个是对应表",testmap[testmicro[i]])
	// fmt.Println(testmicro[i])
	//fmt.Println(testmap[testmicro[i]])
	if len(testmap[testmicro[i]]) == 0 {

		var time []float64
		for _, svcnode := range gl[i] { //循环第i个任务的所有个数得到所有的i类型的任务所在的主机号，然后进行相加
			//fmt.Println("主机号：",svcnode)
			time = append(time, Selfcomputetime(nl, svcnode, i))
		}
		sort.Float64s(time) //待确认
		//fmt.Println("计算时间为：",time[len(time)-1])
		return time[len(time)-1]
	} else {
		var selftimes []float64
		for _, svcnode := range gl[i] {
			//fmt.Println("主机号为：",svcnode)
			selftimes = append(selftimes, Selfcomputetime(nl, svcnode, i))
		}
		sort.Float64s(selftimes)
		selftime := selftimes[len(selftimes)-1] //这个是i类任务的最大的完成时间，后面应该循环调用得出子任务的计算时间
		//fmt.Println("自己的计算的shijianwei :",selftime)
		//得到自己这个类型的任务的计算时间，然后准备计算子任务的计算时间，然后进行比较
		var sonsvctime []float64
		for _, j := range testmap[testmicro[i]] { //应该注意range得到的两个值的不同
			//fmt.Println("子任务为：",j)
			sonsvctime = append(sonsvctime, Computetime(nl, gl, j-1)) //这里应该注意，应该是减一，而不是直接是j
		}
		sort.Float64s(sonsvctime)
		sonsvcmaxtime := sonsvctime[len(sonsvctime)-1] //还有待确认是不是最后一个，也有可能是第一个
		//	fmt.Println("子任务的计算时间为：",sonsvcmaxtime)
		allcomputetime := selftime + sonsvcmaxtime
		//fmt.Println("最终的时间为：",allcomputetime)
		return allcomputetime
	}

}

/*func (ms micoservice) Sonsvctime(nl Nodelist,gl Genelist) float64{
    var sonsvctime float64
    if len(testmap[ms]) == 0 {
    	return 0
	}else {
		for i := range testmap[ms]{
			sonsvctime = sonsvctime + testmicro[testmap[ms][i]-1].Selfcomputetime()
		}
		return  sonsvctime
	}
}*/
