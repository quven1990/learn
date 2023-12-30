package algorithm

import (
	"fmt"
	"math/rand"
	"time"
)

//二分查找相关
type BinarySearch struct {
}

//最基本的二分查找法
func (b *BinarySearch) Run() {
	//获取基础数据
	data := GenerateSortedSlice(10, 30)
	// 设置种子，确保随机性
	rand.Seed(time.Now().Unix())

	// 从切片中随机选择一个元素
	target := data[rand.Intn(len(data))]
	//最简单的二分查找法
	targetIndex := b.BinarySearchSimple(data, target)
	fmt.Println("源数据", data, "目标值", target, "目标值索引", targetIndex)
	//找>=N的最左位置
	target = rand.Intn(30)
	res := b.FindLeftMax(data, target)
	fmt.Println("源数据", data, "目标值", target, "结果", res)
	//找<=N的最右位置
	target = rand.Intn(30)
	res = b.FindRightMin(data, target)
	fmt.Println("源数据", data, "目标值", target, "结果", res)
}

//找<=N的最右位置
func (b *BinarySearch) FindRightMin(data []int, target int) (res int) {
	res = -1
	if len(data) == 0 {
		return res
	}
	l := 0
	r := len(data) - 1
	for {
		m := l + (r-l)/2
		if data[m] <= target {
			l = m + 1
			res = m
		} else {
			r = m - 1
		}
		if l > r {
			break
		}
	}
	return
}

//找>=N的最左位置
func (b *BinarySearch) FindLeftMax(data []int, target int) (res int) {
	res = -1
	if len(data) == 0 {
		return res
	}

	l := 0
	r := len(data) - 1
	for {
		m := l + (r-l)/2
		if data[m] >= target {
			r = m - 1
			res = m
		} else {
			l = m + 1
		}
		if l > r {
			break
		}
	}

	return res
}

//找右边的最小值

//最基本的二分查找法
func (b *BinarySearch) BinarySearchSimple(data []int, target int) (targetIndex int) {
	targetIndex = -1
	if len(data) == 0 {
		return targetIndex
	}
	l := 0
	r := len(data) - 1
	for l <= r {
		m := l + (r-l)/2
		if data[m] == target {
			targetIndex = m
			break
		} else if data[m] > target {
			r = m - 1
		} else if data[m] < target {
			l = m + 1
		}
	}
	return targetIndex
}
