package main

import (
	"fmt"
	"time"
)

var chanMap chan map[int]int
var exitChan chan bool

func WriteMap(chanMap chan map[int]int, size int) {
	for i := 1; i <= size; i++ {
		temp := make(map[int]int, 1)
		temp[i] = i
		chanMap <- temp
		fmt.Println("写入数据：", temp)
	}
	close(chanMap)
}
func ReadMap(chanMap chan map[int]int, exitChan chan bool) {
	for {
		val, ok := <-chanMap
		if !ok {
			break
		}
		fmt.Println("读取到：", val)
	}
	exitChan <- true
}

var intChan chan int
var primeChan chan int

func initChan(intChan chan int, size int) {
	for i := 1; i <= size; i++ {
		intChan <- i
	}
	close(intChan)
}
func CheckPrimeChan(intChan, primeChan chan int, exitChan chan bool) {
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		if CheckPrime(num) {
			primeChan <- num
		}
	}
	exitChan <- true

}

func main() {
	t := time.Now()
	size := 100000
	intChan := make(chan int, size)
	primeChan := make(chan int, size)
	exitChan := make(chan bool, 1)
	go initChan(intChan, size)
	checkChannelNum := 8
	for i := 0; i < checkChannelNum; i++ {
		go CheckPrimeChan(intChan, primeChan, exitChan)
	}
	go func() {
		for i := 0; i < checkChannelNum; i++ {
			<-exitChan
		}
		close(primeChan)

	}()

	for {
		value, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(value, "是素数")
	}
	elapsed := time.Since(t)

	fmt.Println("app elapsed:", elapsed)
	/*
		size := 500000
		chanMap := make(chan map[int]int, size)
		exitChan := make(chan int, 1)
		go WriteMap(chanMap, size)
		go ReadMap(chanMap, exitChan)
		for {
			exit := <-exitChan
			if exit != 0 {
				return
			}
		}


	*/
}

func CheckPrime(num int) bool {
	res := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			res = false
		}
	}
	return res
}
