// @Description
// @Author 小游
// @Date 2021/03/26
package q303

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
}
