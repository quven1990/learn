package Stack

//最小栈
//关键思路(多维护一个栈)
//1. 一个栈用来存储数据,一个栈用来存储最小值
//2. 每次push的时候,如果push的值小于等于最小值栈的栈顶元素,则将该值push到最小值栈中
//3. 每次pop的时候,如果pop的值等于最小值栈的栈顶元素,则将该值pop出最小值栈
//4. 每次getMin的时候,直接返回最小值栈的栈顶元素即可
//5. 这样做的好处是,最小值栈的栈顶元素始终是当前栈中的最小值,时间复杂度为O(1)
//6. 缺点是需要多维护一个栈,空间复杂度为O(n)
type MinStack struct {
	data *BasicStack
	min  *BasicStack
}

func MinStackCreate() MinStack {
	stackSize := 10
	return MinStack{
		data: CreateStack(stackSize),
		min:  CreateStack(stackSize),
	}
}

func (this *MinStack) Push(val int) {
	this.data.Push(val)
	//最小栈数据处理
	if this.min.IsEmpty() { //最小栈空仓直接压入
		this.min.Push(val)
	} else {
		if this.min.Peek() <= val { //如果最小栈的栈顶比压入值小，则直接再次压入最小栈的栈顶
			this.min.Push(this.min.Peek())
		} else { //如果最小栈的栈顶比压入值大，则直接压入最新的值
			this.min.Push(val)
		}
	}
}

func (this *MinStack) Pop() {
	this.data.Pop()
	this.min.Pop()
}

func (this *MinStack) Top() int {
	return this.data.Peek()
}

func (this *MinStack) GetMin() int {
	return this.min.Peek()
}
