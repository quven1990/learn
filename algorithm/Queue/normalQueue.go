package Queue

import (
	"errors"
	"fmt"
)

type QueueArray struct { //数组实现队列
	Left         int   //左指针
	Right        int   //右指针
	QueueElement []int //队列元素
	MaxSize      int   //最大容量
}

//计算队列长度
func (qa *QueueArray) GetSize() (n int) {
	return qa.Right - qa.Left
}

//写入数据
func (qa *QueueArray) Push(ele int) (val int, err error) {
	if qa.Right+1 > qa.MaxSize {
		err = errors.New("超出队列最大长度" + fmt.Sprintf("%s", qa.MaxSize))
		return 0, err
	}
	qa.QueueElement[qa.Right] = ele
	qa.Right += 1 //右指针右移
	return val, nil
}

//弹出数据
func (qa *QueueArray) Pop() (val int, err error) {
	//判断是否还有数据
	if qa.Right-qa.Left <= 0 {
		err = errors.New("该队列无元素")
		return 0, err
	}
	val = qa.QueueElement[qa.Left]
	qa.QueueElement[qa.Left] = 0
	qa.Left += 1
	return val, nil
}

//遍历队列中所有数据
func (qa *QueueArray) Traverse() {
	fmt.Println("开始遍历队列")

	fmt.Println(qa.QueueElement)
}
