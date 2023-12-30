package algorithm

import (
	"math/rand"
	"sort"
	"time"
)

//生成排序过的切片
func GenerateSortedSlice(n int, maxRange int) []int {
	s := make([]int, n)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		s[i] = rand.Intn(maxRange)
	}

	sort.Ints(s)

	return s
}

//生成随机切片
func GenerateSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())

	//创建一个长度为给定大小的切片
	result := make([]int, size)
	elementMap := make(map[int]bool)

	for i := 0; i < size; i++ {
		//生成随机数
		for {
			element := rand.Intn(size)
			if !elementMap[element] {
				elementMap[element] = true
				result[i] = element
				break
			}
		}
	}

	return result
}
