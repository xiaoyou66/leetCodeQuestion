// @Description
// @Author 小游
// @Date 2021/03/25
package q241

import "strconv"

func diffWaysToCompute(expression string) []int {
	var ways []int
	for i := 0; i < len(expression); i++ {
		c := string(expression[i])
		if c == "+" || c == "-" || c == "*" {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, l := range left {
				for _, r := range right {
					switch c {
					case "+":
						ways = append(ways, l+r)
					case "-":
						ways = append(ways, l-r)
					case "*":
						ways = append(ways, l*r)
					}
				}
			}
		}
	}
	if len(ways) == 0 {
		if a, err := strconv.Atoi(expression); err == nil {
			ways = append(ways, a)
		}
	}
	return ways
}
