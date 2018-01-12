package addTwoNumbers

import (
	"testing"
	"reflect"
	"log"
)

func TestSolution(t *testing.T) {
	table := []struct {
		L1     *ListNode
		L2     *ListNode
		Expect *ListNode
	}{
		{listNode([]int{2, 4, 3}), listNode([]int{5, 6, 4}), listNode([]int{7, 0, 8})},
	}

	for i := range table {
		output := Solution(table[i].L1, table[i].L2)
		if !reflect.DeepEqual(table[i].Expect, output) {
			log.Fatalf("case : %v , output: %v", table[i].Expect.ToInts(), output.ToInts())
		}
	}
}
