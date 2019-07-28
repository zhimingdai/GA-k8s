package GA__final

import (
	"math/rand"
)

var ms1 micoservice = struct {
	msreq  int
	cpures float64
	disk   float64
}{msreq: 3, cpures: 0.1, disk: 1.6}
var ms2 micoservice = struct {
	msreq  int
	cpures float64
	disk   float64
}{msreq: 2, cpures: 1.17, disk: 21}
var ms3 micoservice = struct {
	msreq  int
	cpures float64
	disk   float64
}{msreq: 3, cpures: 2.1, disk: 36}
var ms4 micoservice = struct {
	msreq  int
	cpures float64
	disk   float64
}{msreq: 1, cpures: 0.11, disk: 32}
var ms5 micoservice = struct {
	msreq  int
	cpures float64
	disk   float64
}{msreq: 2, cpures: 1.7, disk: 60}
var ms6 micoservice = struct {
	msreq  int
	cpures float64
	disk   float64
}{msreq: 1, cpures: 1.8, disk: 30}
var ms7 micoservice = struct {
	msreq  int
	cpures float64
	disk   float64
}{msreq: 15, cpures: 0.32, disk: 50}
var ms8 micoservice = struct {
	msreq  int
	cpures float64
	disk   float64
}{msreq: 15, cpures: 0.45, disk: 10}
var ms9 micoservice = struct {
	msreq  int
	cpures float64
	disk   float64
}{msreq: 12, cpures: 0.2, disk: 30}
var ms10 micoservice = struct {
	msreq  int
	cpures float64
	disk   float64
}{msreq: 3, cpures: 1.61, disk: 33}
var ms11 micoservice = struct {
	msreq  int
	cpures float64
	disk   float64
}{msreq: 1, cpures: 0.45, disk: 45.1}
var ms12 micoservice = struct {
	msreq  int
	cpures float64
	disk   float64
}{msreq: 3, cpures: 1.63, disk: 20}
var ms13 micoservice = struct {
	msreq  int
	cpures float64
	disk   float64
}{msreq: 3, cpures: 0.42, disk: 40}
var ms14 micoservice = struct {
	msreq  int
	cpures float64
	disk   float64
}{msreq: 3, cpures: 0.32, disk: 31}

var testmap proconsmap

func Creatinit() (proconsmap, Nodelist) {

	var nl1 Nodelist
	testmap = make(proconsmap)
	testmap[ms1] = proconslist{}
	testmap[ms2] = proconslist{13}
	testmap[ms3] = proconslist{4} //这里与原本的不一样，减除了ms3
	testmap[ms4] = proconslist{}
	testmap[ms5] = proconslist{2, 4, 10, 11, 12}
	testmap[ms6] = proconslist{}
	testmap[ms7] = proconslist{5, 6, 9}
	testmap[ms8] = proconslist{7}
	testmap[ms9] = proconslist{}
	testmap[ms10] = proconslist{12}
	testmap[ms11] = proconslist{12}
	testmap[ms12] = proconslist{}
	testmap[ms13] = proconslist{1, 3}
	testmap[ms14] = proconslist{}
	nl1 = make(Nodelist, 80)
	for i := range nl1 {
		nl1[i].Cpund = 4
		nl1[i].Disknd = 100
	}
	return testmap, nl1
}

var testmicro = allmirco{ms1, ms2, ms3, ms4, ms5, ms6, ms7, ms8, ms9, ms10, ms11, ms12, ms13, ms14}

//var Pop1 []chromosome
func (popson Poplutationgene) Soninit() Popluatin {
	var pop Popluatin
	for i := range popson {
		var chrom Chromosome
		chrom.Gl = popson[i]
		_, chrom.Nl = Creatinit()
		pop = append(pop, chrom)
	}
	return pop

}
func (Pop1 Popluatin) Initchrom() Popluatin {

	for i := 0; i < 200; i++ {
		//	//var chro Ch
		var chro1 Chromosome
		//chro1.genelist
		chro1.Gl = make(Genelist, 15)

		for i := 1; i <= 14; i++ {
			num := testmicro[i-1].msreq
			temp := msnodelist(num)
			chro1.Gl[i] = temp
		}
		chro1.Gl = chro1.Gl[1:]
		_, chro1.Nl = Creatinit()
		//fmt.Println(chro1)  //这是一个染色体，也就是一个个体
		//change(chro1)
		Pop1 = append(Pop1, chro1) //这就是一个种群，有200个个体

	}
	return Pop1
}

func msnodelist(num int) []int {
	//rand.Seed(time.Now().UnixNano())
	//num := rand.Intn(51)
	var gene = make([]int, num)

	for i := 0; i < num; i++ {
		//rand.Seed(time.Now().UnixNano())
		gene[i] = rand.Intn(80)

	}
	return gene
}

func msnodelistsecond(num int) []int {
	//rand.Seed(time.Now().UnixNano())
	//num := rand.Intn(51)
	var gene = make([]int, num)

	for i := 0; i < num; i++ {
		//rand.Seed(time.Now().UnixNano())
		gene[i] = rand.Intn(80)

	}
	return gene
}
