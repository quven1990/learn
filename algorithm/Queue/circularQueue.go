package Queue

//循环队列
type MyCircularQueue struct {
	Size     int //用来计算是否溢出
	Left     int
	Right    int
	QueueEle []int
	Limit    int //总长度
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Limit:    k,
		QueueEle: make([]int, k, k),
	}
}

//写入队列元素
func (this *MyCircularQueue) EnQueue(value int) bool {
	//判断是否溢出
	if (this.Size + 1) > this.Limit { //溢出了直接返回false
		return false
	}
	this.Size++
	this.QueueEle[this.Right] = value //给队列赋值
	if this.Right+1 > this.Limit-1 {  //超出界限从头计算
		this.Right = 0
	} else {
		this.Right++
	}
	return true
}

//弹出队列元素
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() { //空队列 无法弹出
		return false
	}
	if (this.Left + 1) == this.Limit {
		this.Left = 0
	} else {
		this.Left++
	}
	this.Size-- //长度-1
	return true
}

//弹出队列元素
func (this *MyCircularQueue) DeQueueWithVal() int {
	if this.IsEmpty() { //空队列 无法弹出
		return -1
	}
	if (this.Left + 1) == this.Limit {
		this.Left = 0
	} else {
		this.Left++
	}
	this.Size-- //长度-1
	var index int
	if this.Left-1 < 0 {
		index = this.Limit - 1
	} else {
		index = this.Left - 1
	}
	return this.QueueEle[index]
}

//返回队列第一个元素
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	//根据left计算下标
	//var index int
	//
	//if this.Left==0
	//index = this.Left - 1
	return this.QueueEle[this.Left]
}

//弹出尾部元素
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	var index int
	if this.Right == 0 {
		index = this.Limit - 1
	} else {
		index = this.Right - 1
	}

	return this.QueueEle[index]
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.Size == 0 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	if this.Size == this.Limit {
		return true
	}
	return false
}
