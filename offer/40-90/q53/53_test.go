// @Description
// @Author 小游
// @Date 2021/04/09
package q53

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	//fmt.Println(majorityElement([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(search([]int{0}, 8))
	fmt.Println(search([]int{1, 1, 2}, 1))
}
