package algorithm

import (
	"fmt"
	"math"
	"math/rand"
)

//基尼系数计算社会财富是否平均分配
func GiniRun() {
	data := makeGiniBaseData()
	var times int
	times = 1000000 //分配次数
	data = jiniExchangeMoney(data, times)

	fmt.Println(data)
	jini := calcGini(data)
	fmt.Println(jini)
}

//社会财富开始分配
func jiniExchangeMoney(data []float64, times int) (res []float64) {
	res = data
	for i := 0; i < times; i++ {
		for k, _ := range data {
			//乱序找一个人
			for { //不能是自己
				randK := rand.Intn(len(data))
				if randK != k {
					if data[k] <= 0 {
						break
					}
					data[randK] = data[randK] + 1
					data[k] = data[k] - 1
					break
				}
			}
		}
	}

	return res
}

//计算基尼系数
func calcGini(data []float64) (jini float64) {
	//total
	var total float64
	var total2 float64
	for _, i := range data {
		total += i
		for _, j := range data {
			total2 += math.Abs(i - j)
		}
	}
	jini = total2 / (2 * total * float64(len(data)))
	return
}

//构造原始数据
func makeGiniBaseData() []float64 {
	data := make([]float64, 0, 100)
	for i := 0; i < 100; i++ {
		data = append(data, 100)
	}
	return data
}
