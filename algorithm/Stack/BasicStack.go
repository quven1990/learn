package Stack

type BasicStack struct {
	StackElement []int
	Size         int
	Limit        int
}

//创建栈
func CreateStack(size int) *BasicStack {
	return &BasicStack{
		Limit:        size,
		StackElement: make([]int, size, size),
	}
} //塞入栈数据
func (b *BasicStack) Push(val int) bool {
	if b.Size+1 > b.Limit {
		return false
	}
	b.StackElement[b.Size] = val
	b.Size++
	return true
}

//弹出栈数据
func (b *BasicStack) Pop() int {
	if b.IsEmpty() {
		return -1
	}
	res := b.StackElement[b.Size-1]
	b.Size--
	return res

}

//判断栈是否为空
func (b *BasicStack) IsEmpty() bool {
	if b.Size == 0 {
		return true
	}
	return false
}
