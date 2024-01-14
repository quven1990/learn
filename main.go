package main

import (
	"fmt"
	"test/algorithm/Stack"
	"time"
)

type ServerTemp struct {
	Addr     string
	Port     string
	Timeout  time.Duration
	MaxConns int
}

/*
func NewServerWithTimeout(addr, port string, timeout time.Duration) *Server {
	return &Server{addr: addr, port: port, timeout: timeout}
}
func NewServerWithMaxConns(addr, port string, maxConns int) *Server {
	return &Server{addr: addr, port: port, maxConns: maxConns}
}
func NewServerWithTimeoutAndMaxConns(addr, port string, timeout time.Duration, maxConns int) *Server {
	return &Server{addr: addr, port: port, timeout: timeout, maxConns: maxConns}
}

func NewDefaultServer(addr, port string, timeout time.Duration, maxConns int) *Server {
	return &Server{addr: addr, port: port, timeout: timeout, maxConns: maxConns}
}

*/

type ServerBuilder struct {
	Server
}

func (sb *ServerBuilder) Create(addr, port string) *ServerBuilder {
	sb.Server.addr = addr
	sb.Server.port = port
	return sb
}
func (sb *ServerBuilder) WithTimeout(timeout time.Duration) *ServerBuilder {
	sb.Server.timeout = timeout
	return sb
}
func (sb *ServerBuilder) WithMaxConns(maxConns int) *ServerBuilder {
	sb.Server.maxConns = maxConns
	return sb
}
func (sb *ServerBuilder) Build() Server {
	if sb.timeout == 0 {
		sb.timeout = 5 * time.Microsecond
	}
	if sb.maxConns == 0 {
		sb.maxConns = 20
	}
	return sb.Server
}

type Server struct {
	addr     string
	port     string
	timeout  time.Duration
	maxConns int
}
type Option func(*Server)

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}
func MaxConns(maxConns int) Option {
	return func(s *Server) {
		s.maxConns = maxConns
	}
}
func Create(addr, port string, options ...Option) (*Server, error) {
	s := Server{
		addr:     addr,
		port:     port,
		timeout:  500 * time.Microsecond,
		maxConns: 20,
	}
	for _, option := range options {
		option(&s)
	}
	return &s, nil
}

// chapter3/sources/slice_unbind_orig_array.go

func SeliceTest(a [3]int) {
	a[0] = 123
	/*
		u := []int{11, 12, 13, 14, 15}
		fmt.Println("array:", u) // [11, 12, 13, 14, 15]
		s := u[1:3]
		fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // [12, 13]
		s = append(s, 24)
		fmt.Println("after append 24, array:", u)
		fmt.Printf("after append 24, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
		s = append(s, 25)
		fmt.Println("after append 25, array:", u)
		fmt.Printf("after append 25, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
		s = append(s, 26)
		fmt.Println("after append 26, array:", u)
		fmt.Printf("after append 26, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)

		s[0] = 22
		fmt.Println("after reassign 1st elem of slice, array:", u)
		fmt.Printf("after reassign 1st elem of slice, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
	*/

}

func main() {
	//根据基尼系数计算财富分配是否合理
	//algorithm.GiniRun()
	//三傻算法
	//algorithm.ThreeFoolRun()
	//二分查找法相关
	//res := algorithm.BinarySearch{}
	//res.Run()
	//链表相关
	//obj := LinkedList.LinkedList{}
	//obj.Run()
	//队列相关(循环队列和普通队列)
	//obj := Queue.Queue{}
	//err := obj.Run()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//栈相关
	obj := Stack.Stack{}
	err := obj.Run()
	if err != nil {
		fmt.Println(err)
	}
	//time.Sleep(1*time.Second)

	/*
		var a [3]int = [3]int{1, 2, 3}

		SeliceTest(a)
		fmt.Println(a)
	*/

	/*
		server, err := Create("test_addr", "test_port", Timeout(100*time.Second))
		fmt.Println(server)
		fmt.Println(err)
	*/

	//server := new(ServerBuilder).Create("test_port", "test_port").WithMaxConns(10).Build()
	//fmt.Println("aaaaa")
	//roun.Test()
	//var arr = []int{7, 3, 10, 4, 5, 9, 123, 6, 8, 54, 47, 789}
	//选择排序
	//res := selectSort(arr)
	//冒泡排序
	//res := bubbleSort(arr)
	//插入排序
	//res := insertSort(arr)
	//快速排序
	//res := quickSort(arr)
	//a := 100
	//b := 2
	//SwapBit(a, b)
	//arr := []int{1, 2, 2, 3, 3, 1, 5, 5, 5, 6, 6, 7, 7, 8}
	//left, right := findTwoOddTimesNumber(arr)
	//fmt.Println(left, right)

	//二分法
	//var arr = []int{1, 4, 7, 11, 23, 56, 198}
	//num := 57
	//index := BindrySearch(arr, num)
	//fmt.Println(index, arr[index])

	//递归找最大数
	//var arr = []int{1, 6, 8, 4, 2, 6, 9}
	//fmt.Println(findMaxByRecursion(arr, 0, len(arr)-1))

}

func findMaxByRecursion(arr []int, l, r int) int {
	if l == r {
		return arr[l]
	}
	fmt.Println("aaaaaaa")

	middle := l + (r-l)/2
	fmt.Println("l:middle", l, middle)
	leftMax := findMaxByRecursion(arr, l, middle)
	fmt.Println("bbbbbbbbb")
	fmt.Println("l:middle", middle+1, r)
	rightMax := findMaxByRecursion(arr, middle+1, r)
	fmt.Println("ccccccccc")

	return Max(leftMax, rightMax)
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//二分法
//在一个有序数组中，找某个数看是否存在
//在一个有序数组中，找出》=某个数的最左位置
//局部最小值问题
func BindrySearch(arr []int, num int) (index int) {
	if len(arr) < 2 {
		return
	}
	max_index := len(arr) - 1
	min_index := 0
	for {
		if min_index > max_index {
			index = -1
			break
		}
		middle_index := min_index + (max_index-min_index)/2
		if arr[middle_index] > num {
			max_index = middle_index - 1
		} else if arr[middle_index] < num {
			min_index = middle_index + 1
		} else {
			index = middle_index
			break
		}
	}

	return
}

//一个数组中有两种数出现了奇数次，其他数都出现了偶数次，怎么找到这一个数
func findTwoOddTimesNumber(arr []int) (left, right int) {
	for i := 0; i < len(arr); i++ {
		left ^= arr[i]
	}
	temp := left & (^left + 1)

	for i := 0; i < len(arr); i++ {
		if arr[i]&temp == 0 {
			right ^= arr[i]
		}
	}
	left = right ^ left
	return left, right
}

//一个数组中有一种数出现了奇数次，其他数都出现了偶数次，怎么找到这一个数
func findOddTimesNumber(arr []int) (eor int) {
	for i := 0; i < len(arr); i++ {
		eor ^= arr[i]
	}
	return eor
}
func SwapBit(a, b int) {
	fmt.Println(a, b)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}

//快速排序
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	middle := arr[0]
	var left []int
	var right []int
	for i := 1; i < len(arr); i++ {
		if arr[i] > middle {
			right = append(right, arr[i])
		} else {
			left = append(left, arr[i])
		}
	}

	middle_s := []int{middle}
	left = quickSort(left)
	right = quickSort(right)
	arr = append(append(left, middle_s...), right...)
	return arr
}

//插入排序
func insertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				Swap(arr, j, j+1)
			}
		}
	}
	return arr
}

//冒泡排序
func bubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for e := len(arr) - 1; e > 0; e-- {
		for i := 0; i < e; i++ {
			if arr[i] > arr[i+1] {
				Swap(arr, i, i+1)
			}
		}
	}
	return arr
}

//选择排序
func selectSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		var minIndex int = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr = Swap(arr, i, minIndex)
	}
	return arr

}
func Swap(arr []int, i, j int) []int {
	temp := arr[j]
	arr[j] = arr[i]
	arr[i] = temp
	return arr
}
