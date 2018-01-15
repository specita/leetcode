package maximumSubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	table := []struct {
		nums   []int
		expect int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	for _, e :=range table {
		output := MaxSubArray(e.nums)
		if output != e.expect {
			t.Fatalf("case:%v, output: %v", e, output)
		}
	}
}
