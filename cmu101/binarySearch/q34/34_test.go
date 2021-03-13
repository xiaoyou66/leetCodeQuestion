// @Description 
// @Author 小游
// @Date 2021/03/13
package q34

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T)  {
	fmt.Println(searchRange([]int{5,7,7,8,8,10},8))
	fmt.Println(searchRange([]int{},0))
	fmt.Println(searchRange([]int{1},0))
}
