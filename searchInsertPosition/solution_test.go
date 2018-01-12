package searchInsertPosition

import "testing"

func TestSearchInsert(t *testing.T) {
	table := []struct {
		nums   []int
		target int
		expect int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 3, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
	}

	for _, e := range table {
		output := SearchInsert(e.nums, e.target)
		if output != e.expect {
			t.Fatalf("case: %v, output: %v", e, output)
		}
	}

}
