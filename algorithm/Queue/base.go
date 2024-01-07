package Queue

import "fmt"

//队列相关学习
type Queue struct {
}

func (q *Queue) Run() (err error) {
	//用队列实现栈
	obj := QueueToStackConstructor(4)
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)

	//fmt.Println(obj.Queue.QueueEle)

	fmt.Println(obj.Top())
	fmt.Println(obj.Top())
	fmt.Println(obj.Top())
	fmt.Println(obj.Top())
	fmt.Println("aaaaaaa")
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	/*
		//循环队列学习
		obj := Constructor(3) //初始化队列长度
		fmt.Println(obj.EnQueue(1))
		fmt.Println(obj.EnQueue(2))
		fmt.Println(obj.EnQueue(3))
		fmt.Println(obj.EnQueue(4))
		fmt.Println(obj.Rear())
		fmt.Println(obj.IsFull())
		fmt.Println(obj.DeQueue())
		fmt.Println(obj.EnQueue(4))
		fmt.Println(obj.Rear())
	*/
	//普通队列
	/*
		qa := new(QueueArray)
		qa.MaxSize = 4
		qa.QueueElement = make([]int, qa.MaxSize, qa.MaxSize)

		_, err = qa.Push(1)
		fmt.Println("队列当前长度为", qa.GetSize())
		qa.Traverse()
		//弹出队列
		ele1, err := qa.Pop()
		if err != nil {
			return err
		}
		fmt.Println("弹出元素为", ele1, "队列当前长度为", qa.GetSize())
		qa.Traverse()
		_, err = qa.Push(2)
		fmt.Println("队列当前长度为", qa.GetSize())
		qa.Traverse()
		//弹出队列
		ele2, err := qa.Pop()
		if err != nil {
			return err
		}
		fmt.Println("弹出元素为", ele2, "队列当前长度为", qa.GetSize())
		qa.Traverse()

		_, err = qa.Push(3)
		fmt.Println("队列当前长度为", qa.GetSize())
		qa.Traverse()
		_, err = qa.Push(4)
		fmt.Println("队列当前长度为", qa.GetSize())
		qa.Traverse()
		if err != nil {
			return err
		}

		ele3, err := qa.Pop()
		if err != nil {
			return err
		}
		fmt.Println("弹出元素为", ele3, "队列当前长度为", qa.GetSize())
		qa.Traverse()
		ele4, err := qa.Pop()
		if err != nil {
			return err
		}
		fmt.Println("弹出元素为", ele4, "队列当前长度为", qa.GetSize())
		qa.Traverse()
	*/
	return
}
