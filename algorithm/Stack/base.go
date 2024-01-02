package Stack

import "fmt"

type Stack struct {
}

func (s *Stack) Run() (err error) {
	/*
		//基本的栈操作
		obj := CreateStack(3)
		fmt.Println(obj.Push(1))
		fmt.Println(obj.Push(2))
		fmt.Println(obj.Pop())
		fmt.Println(obj.Push(3))
		fmt.Println(obj.Pop())
		fmt.Println(obj.Push(4))
		fmt.Println(obj.Pop())

		fmt.Println(obj.Pop())
		fmt.Println(obj.Pop())
		fmt.Println(obj.Pop())
	*/
	//用栈实现队列
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	fmt.Println(obj.Peek())
	fmt.Println(obj.Peek())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	return err
}
