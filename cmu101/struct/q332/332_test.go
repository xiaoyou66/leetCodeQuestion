// @Description
// @Author 小游
// @Date 2021/03/26
package q332

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	defer fmt.Println("12")
	{
		defer fmt.Println("defer runs")
		fmt.Println("block ends")
	}
	fmt.Println("main ends")

}
