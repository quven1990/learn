package Queue

//用队列实现栈
type QueueToStack struct {
	Queue *MyCircularQueue
}

func QueueToStackConstructor(k int) QueueToStack {
	res := QueueToStack{
		Queue: &MyCircularQueue{},
	}
	res.Queue.Limit = k
	res.Queue.QueueEle = make([]int, k, k)
	return res
}

//important  关键点
//插入元素的时候，需要把前面已有的元素全部重新插入到尾部 即可实现栈
//example
// 1<-2<-3 想继续插入4 则需要
//2<-3<-4<-1 (把1插入尾部)
//3<-4<-1<-2 (把2插入尾部)
//4<-1<-2<-3 (把3插入尾部)
func (this *QueueToStack) Push(x int) {
	size := this.Queue.Size
	this.Queue.EnQueue(x)
	//fmt.Println("aaaaaaa")
	//fmt.Println(this.Queue.QueueEle)
	for i := 0; i < size; i++ {
		//fmt.Println("bbbbb")
		temp := this.Queue.DeQueueWithVal()
		//fmt.Println(temp)
		this.Queue.EnQueue(temp)
		//fmt.Println(this.Queue.QueueEle)
	}
}

func (this *QueueToStack) Pop() int {
	return this.Queue.DeQueueWithVal()
}

func (this *QueueToStack) Top() int {
	return this.Queue.Front()
}

func (this *QueueToStack) Empty() bool {
	return this.Queue.IsEmpty()
}
