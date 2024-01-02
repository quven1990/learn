package Stack

//用栈实现队列
type StackToQueue struct {
	In  *BasicStack
	Out *BasicStack
}

//创建对象
func Constructor() StackToQueue {
	res := StackToQueue{
		In:  &BasicStack{},
		Out: &BasicStack{},
	}
	res.In = CreateStack(10)
	res.Out = CreateStack(10)
	return res
}

//important
//转换栈数据
//1 out栈为空了，才能继续往里倒数据
//2 in栈的数据要一次性倒完
func (this *StackToQueue) InToOut() {
	if this.Out.IsEmpty() {
		for {
			if this.In.IsEmpty() {
				break
			}
			this.Out.Push(this.In.Pop())
		}
	}
}

func (this *StackToQueue) Push(x int) {
	this.In.Push(x)
	this.InToOut()
}

func (this *StackToQueue) Pop() int {
	this.InToOut()
	return this.Out.Pop()
}

func (this *StackToQueue) Peek() int {
	this.InToOut()
	return this.Out.Peek()
}

func (this *StackToQueue) Empty() bool {
	return this.In.IsEmpty() && this.Out.IsEmpty()
}
