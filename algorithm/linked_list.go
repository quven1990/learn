package algorithm

import (
	"fmt"
)

type LinkedList struct {
}

type LinkedListNode struct {
	val  int
	next *LinkedListNode
}

var headNode *LinkedListNode = nil

func (l *LinkedList) Run() error {
	//头插法
	/*
		headNode = l.InitHeadNode(headNode, 0)
		l.AddNodeHead(headNode, 1)
		l.AddNodeHead(headNode, 2)
		l.AddNodeHead(headNode, 3)
		l.AddNodeHead(headNode, 4)
		l.AddNodeHead(headNode, 5)
		l.AddNodeHead(headNode, 6)
		l.AddNodeHead(headNode, 7)
		l.Traverse(headNode)*/
	//尾插法
	/*
		headNode = l.InitHeadNode(headNode, 0)
		l.AddNodeTail(headNode, 1)
		l.AddNodeTail(headNode, 2)
		l.AddNodeTail(headNode, 3)
		l.AddNodeTail(headNode, 4)
		l.AddNodeTail(headNode, 5)
		l.AddNodeTail(headNode, 6)
		l.AddNodeTail(headNode, 7)
		l.Traverse(headNode)*/

	//删除链表节点
	/*
		l.DeleteNode(headNode, 1)
		l.DeleteNode(headNode, 7)
		l.Traverse(headNode)*/

	//链表反转
	/*
		headNode = l.Reverse(headNode)
		l.Traverse(headNode)
	*/
	/*
		//链表合并
		//构建两个链表
		var headNode1 *LinkedListNode = nil
		headNode1 = l.InitHeadNode(headNode1, 2)
		l.AddNodeTail(headNode1, 4)
		l.AddNodeTail(headNode1, 8)
		l.AddNodeTail(headNode1, 9)
		l.Traverse(headNode1)
		var headNode2 *LinkedListNode = nil
		headNode2 = l.InitHeadNode(headNode2, 1)
		l.AddNodeTail(headNode2, 5)
		l.AddNodeTail(headNode2, 6)
		l.AddNodeTail(headNode2, 10)
		l.AddNodeTail(headNode2, 11)
		l.Traverse(headNode2)
		//合并两个链表
		mergeHead := l.MergeTwoLists(headNode1, headNode2)
		l.Traverse(mergeHead)
	*/
	/*
		//两个链表相加
		var headNode1 *LinkedListNode = nil
		headNode1 = l.InitHeadNode(headNode1, 4)
		l.AddNodeTail(headNode1, 8)
		l.AddNodeTail(headNode1, 9)
		l.Traverse(headNode1)
		var headNode2 *LinkedListNode = nil
		headNode2 = l.InitHeadNode(headNode2, 2)
		l.AddNodeTail(headNode2, 6)
		l.AddNodeTail(headNode2, 7)
		l.AddNodeTail(headNode2, 9)
		l.Traverse(headNode2)
		//两个链表相加
		addRes := l.AddTwoNumbers(headNode1, headNode2)
		l.Traverse(addRes)
	*/
	/*
		//划分链表
		var OriginLink *LinkedListNode = nil
		OriginLink = l.InitHeadNode(OriginLink, 10)
		l.AddNodeTail(OriginLink, 4)
		l.AddNodeTail(OriginLink, 8)
		l.AddNodeTail(OriginLink, 6)
		l.AddNodeTail(OriginLink, 2)
		l.AddNodeTail(OriginLink, 0)
		l.AddNodeTail(OriginLink, 9)
		l.AddNodeTail(OriginLink, 3)
		l.Traverse(OriginLink)
		var x int
		rand.Seed(time.Now().UnixNano())
		x = rand.Intn(11) // 生成0到10的随机数
		//x = -1
		res := devideLinkList(OriginLink, x)
		fmt.Println("随机数为:", x)
		l.Traverse(res)*/
	return nil
}

//遍历链表
func (l *LinkedList) Traverse(node *LinkedListNode) error {
	if node == nil {
		return nil
	}
	fmt.Println("开始遍历链表")
	for {
		fmt.Println(node.val)
		node = node.next
		if node == nil {
			break
		}
	}
	return nil
}

//初始化链表头部
func (l *LinkedList) InitHeadNode(node *LinkedListNode, val int) (head *LinkedListNode) {
	if node == nil {
		head = &LinkedListNode{
			val:  val,
			next: nil,
		}
		return head
	}
	return nil
}

//链表增加节点(尾插法)
func (l *LinkedList) AddNodeTail(node *LinkedListNode, val int) error {

	if node == nil {
		return nil
	}
	//找到尾节点
	for {
		tail := node.next
		if tail == nil { //尾节点
			temp := new(LinkedListNode)
			temp.val = val
			temp.next = nil
			node.next = temp
			break
		}
		node = node.next
	}
	return nil
}

//链表增加节点(头插法)
func (l *LinkedList) AddNodeHead(node *LinkedListNode, val int) error {
	if node == nil { //空链表直接赋值
		node = &LinkedListNode{
			val:  val,
			next: nil,
		}
		return nil
	}

	temp := new(LinkedListNode)
	temp.val = val
	temp.next = node.next
	node.next = temp
	return nil
}
