// @Description 
// @Author 小游
// @Date 2021/03/15
package q347

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T)  {
	fmt.Println(topKFrequent([]int{1,1,1,2,2,3},2))
	fmt.Println(topKFrequent([]int{1},1))
	fmt.Println(topKFrequent([]int{1,2},1))
}