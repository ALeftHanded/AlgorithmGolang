package SinglyLinkedList

// 160. 相交链表
// https://leetcode-cn.com/problems/intersection-of-two-linked-lists

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	h1, h2 := headA, headB
	for h1 != h2 {
		h1 = h1.Next
		h2 = h2.Next
		if h1 == nil{
			h1 = headB
		}
		if h2 == nil{
			h2 = headA
		}
	}
	return h1
}