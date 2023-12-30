package LinkedList

// 链表合并(两个链表都是从小到大排序，并且元素不重复)
func (l *LinkedList) MergeTwoLists(l1 *LinkedListNode, l2 *LinkedListNode) (head *LinkedListNode) {
	//特殊情况兼容
	if l1 == nil {
		head = l2
		return head
	}
	if l2 == nil {
		head = l1
		return head
	}
	var curr1 *LinkedListNode = nil
	var curr2 *LinkedListNode = nil
	var pre *LinkedListNode = nil
	//初始化头节点
	if l1.val <= l2.val {
		head = l1
		curr1 = l1.next
		curr2 = l2
	} else {
		head = l2
		curr1 = l2.next
		curr2 = l1
	}
	pre = head
	for {
		if curr1 == nil || curr2 == nil {
			break
		}
		if curr1.val <= curr2.val {
			pre.next = curr1
			curr1 = curr1.next
		} else {
			pre.next = curr2
			curr2 = curr2.next
		}
		pre = pre.next
	}
	if curr1 == nil {
		pre.next = curr2
	}
	if curr2 == nil {
		pre.next = curr1
	}

	return head
}
