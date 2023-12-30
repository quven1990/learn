package LinkedList

import "fmt"

// 链表相加计算   2->4->3 + 5->6->4 = 7->0->8
func (l *LinkedList) AddTwoNumbers(l1 *LinkedListNode, l2 *LinkedListNode) (head *LinkedListNode) {
	if l1 == nil {
		head = l2
		return head
	}
	if l2 == nil {
		head = l1
		return head
	}
	head = nil
	temp := new(LinkedListNode)
	temp = nil
	//sum 数字相加结果
	//curr 当前位是几
	//carry 进位
	var sum, curr, carry int
	for {
		if (l1 == nil) && (l2 == nil) {
			break
		}
		if l1 == nil {
			sum = l2.val + carry
		} else if l2 == nil {
			sum = l1.val + carry
		} else {
			sum = l1.val + l2.val + carry
		}
		curr = int(sum % 10)
		fmt.Println("sum:", sum, "curr:", curr, "carry:", carry)
		if head == nil {
			//head = l.InitHeadNode(head, curr)
			head = &LinkedListNode{
				val: curr,
			}
			temp = head
		} else {
			temp.next = &LinkedListNode{
				val: curr,
			}
			temp = temp.next

			//head = head.next
			//l.AddNodeTail(head, curr)
		}
		carry = int(sum / 10) //进位
		if l1 != nil {
			l1 = l1.next
		}
		if l2 != nil {
			l2 = l2.next
		}
	}
	if carry > 0 { //补充最后一位
		temp.next = &LinkedListNode{
			val: carry,
		}
	}

	return head
}
