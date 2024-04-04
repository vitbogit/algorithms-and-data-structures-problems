// Doesn`t work!!! Because even int64 can`t hold big numbers in tests
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var num1, num2, sum int
	for k := 1; l1 != nil || l2 != nil; k *= 10 {
		if l1 != nil {
			num1 += l1.Val * k
			l1 = l1.Next
		}
		if l2 != nil {
			num2 += l2.Val * k
			l2 = l2.Next
		}
	}
	sum = num1 + num2
	l3 := &ListNode{Val: sum % 10}
	head := l3
	sum /= 10
	for sum > 0 {
		l3.Next = &ListNode{Val: sum % 10}
		l3 = l3.Next
		sum /= 10
	}
	return head
}

