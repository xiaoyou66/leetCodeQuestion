// @Description
// @Author 小游
// @Date 2021/03/26
package q332

type Stack struct {
	data []string
}

func (s *Stack) Push(k string) {
	s.data = append(s.data, k)
}
func (s *Stack) Pop() (x string) {
	x = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return
}
func (s *Stack) Empty() bool {
	return len(s.data) == 0
}
func (s *Stack) Top() string {
	return s.data[len(s.data)-1]
}

func findItinerary(tickets [][]string) []string {
	var ans []string
	if len(tickets) == 0 {
		return ans
	}
	hash := map[string][]string{}
	// 遍历tickets
	for _, v := range tickets {
		hash[v[0]] = append(hash[v[0]], v[1])
	}
	s := Stack{}
	s.Push("JFK")
	for !s.Empty() {
		next := s.Top()
		if _, ok := hash[next]; ok {
			ans = append(ans, next)
			s.Pop()
		} else {
			s.Push(hash[next][0])
		}
	}
	return ans
}
