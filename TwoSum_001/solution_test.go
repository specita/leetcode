package TwoSum

import (
	"testing"
	"reflect"
	"log"
)

func TestSolution(t *testing.T) {
	table := []struct {
		Nums   []int
		Target int
		Expect []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for _, e := range table {
		output := Solution(e.Nums, e.Target)
		if !reflect.DeepEqual(output, e.Expect) {
			log.Fatalf("error case : %v, output: %v", e, output)
		}
	}
}
