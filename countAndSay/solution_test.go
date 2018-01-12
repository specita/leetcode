package countAndSay

import "testing"

func TestCountAndSay(t *testing.T) {
	table := []struct {
		n      int
		expect string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
		{6, "312211"},
	}

	for _, i := range table {
		output := CountAndSay(i.n)
		if output != i.expect {
			t.Fatalf("case:%v, output:%v", i, output)
		}
	}
}
