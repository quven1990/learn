package LinkedList

//划分链表(小于x的元素放左边，大于x的元素放右边，然后两个链表合并)
func devideLinkList(originLink *LinkedListNode, x int) (res *LinkedListNode) {
	if originLink == nil {
		return res
	}
	//左右分别计算，最终直接合并
	var leftLink *LinkedListNode = nil
	var leftLinkCurr *LinkedListNode = nil
	var rightLink *LinkedListNode = nil
	var rightLinkCurr *LinkedListNode = nil
	for {
		if originLink == nil {
			break
		}
		if originLink.val < x { //左边链表
			if leftLink == nil {
				leftLink = &LinkedListNode{
					val: originLink.val,
				}
				leftLinkCurr = leftLink
			} else {
				leftLinkCurr.next = &LinkedListNode{
					val: originLink.val,
				}
				leftLinkCurr = leftLinkCurr.next
			}
		} else { //右边链表
			if rightLink == nil {
				rightLink = &LinkedListNode{
					val: originLink.val,
				}
				rightLinkCurr = rightLink
			} else {
				rightLinkCurr.next = &LinkedListNode{
					val: originLink.val,
				}
				rightLinkCurr = rightLinkCurr.next
			}
		}
		//fmt.Println(originLink.val)
		originLink = originLink.next
	}
	//fmt.Println("调试信息")
	//fmt.Printf("%#v\n", leftLink)
	//fmt.Printf("%#v\n", rightLink)
	//两个链表合并
	if leftLink == nil {
		res = rightLink
	} else {
		leftLinkCurr.next = rightLink
		res = leftLink
	}

	return res
}
