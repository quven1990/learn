package Stack

import "fmt"

type Stack struct {
}

func (s *Stack) Run() (err error) {
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

	return err
}
