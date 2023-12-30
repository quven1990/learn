package algorithm

import "fmt"

//最简单的三傻排序

func ThreeFoolRun() {
	//选择排序
	baseData := GenerateSlice(14)
	fmt.Println("原始数据", baseData)
	selectSortRes := selectSort(baseData)
	fmt.Println("选择排序结果", selectSortRes)
	fmt.Println("==========================")

	//冒泡排序
	baseData = GenerateSlice(19)
	fmt.Println("原始数据", baseData)
	bubbleSortRes := bubbleSort(baseData)
	fmt.Println("冒泡排序结果", bubbleSortRes)
	fmt.Println("==========================")
	//插入排序
	baseData = GenerateSlice(21)
	fmt.Println("原始数据", baseData)
	insertSortRes := insertSort(baseData)
	fmt.Println("插入排序结果", insertSortRes)
	fmt.Println("==========================")

	//快速排序
	baseData = GenerateSlice(21)
	fmt.Println("原始数据", baseData)
	quickSortRes := quickSort(baseData)
	fmt.Println("快速排序结果", quickSortRes)
}

//快速排序
func quickSort(data []int) []int {
	return data
}

//插入排序
func insertSort(data []int) []int {
	for i := 1; i < len(data); i++ {
		for j := i - 1; j >= 0; j-- {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

//冒泡排序
func bubbleSort(data []int) []int {

	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}

//选择排序
func selectSort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		minIndex := i
		min := data[i]
		for j := i + 1; j < len(data); j++ {
			if data[j] < min {
				min = data[j]
				minIndex = j
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}
	return data
}
