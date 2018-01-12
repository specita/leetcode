package addTwoNumbers

//You are given two non-empty linked lists representing two non-negative integers.
//The digits are stored in reverse order and each of their nodes contain a single digit.
//Add the two numbers and return it as a linked list.
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//Example:
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.

type ListNode struct {
	Val  int
	Next *ListNode
}

func listNode(nums []int) *ListNode {
	head := &ListNode{}
	iterator := head
	for _, v := range nums {
		iterator.Next = &ListNode{v, nil}
		iterator = iterator.Next
	}
	return head.Next
}

func (l *ListNode) ToInts() []int {
	if l == nil {
		return []int{}
	}

	iterator := l
	r := []int{iterator.Val}
	for i := iterator.Next; i != nil; i = i.Next {
		r = append(r, i.Val)
	}

	return r
}

func Solution(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(l *ListNode, c chan int) {
		for i := l; i != nil; i = i.Next {
			c <- i.Val
		}
		close(c)
	}(l1, ch1)

	go func(l *ListNode, c chan int) {
		for i := l; i != nil; i = i.Next {
			c <- i.Val
		}
		close(c)
	}(l2, ch2)

	l := new(ListNode)
	h := l
	carry := 0
	for {
		i, ok1 := <-ch1
		j, ok2 := <-ch2

		if !ok1 && !ok2 && carry == 0 {
			break
		}

		l.Next = new(ListNode)
		l.Next.Val = (i + j + carry) % 10
		carry = (i + j + carry) / 10
		l = l.Next
	}

	return h.Next
}
