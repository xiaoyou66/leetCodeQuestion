// @Description
// @Author 小游
// @Date 2021/03/25
package q384

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	s := Constructor([]int{4, 9, 5})
	fmt.Println(s.Shuffle())
	fmt.Println(s.Reset())
}
