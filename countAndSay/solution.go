package countAndSay

//The count-and-say sequence is the sequence of integers with the first five terms as following:
//
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//1 is read off as "one 1" or 11. （1 被读作 1个1）
//11 is read off as "two 1s" or 21. (2 被读作 2个1)
//21 is read off as "one 2, then one 1" or 1211. （3 被读作 1个2 1个1）
//Given an integer n, generate the nth term of the count-and-say sequence.
//
//Note: Each term of the sequence of integers will be represented as a string.
//
//Example 1:
//
//Input: 1
//Output: "1"
//Example 2:
//
//Input: 4
//Output: "1211"

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	pre := CountAndSay(n - 1)
	var current = pre[0]
	var counter byte = 1
	var result []byte
	for i := 1; i < len(pre); i++ {
		if current == pre[i] {
			counter++
			continue
		}

		result = append(result, counter+byte('0'), current)
		counter = 1
		current = pre[i]
	}

	result = append(result, counter+byte('0'), current)
	return string(result)
}
