// @Description 
// @Author 小游
// @Date 2021/03/12
package q88

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T)  {
	arr1 := []int{1,2,3,0,0,0}
	arr2 := []int{2,5,6}
	merge(arr1,3,arr2,3)
	fmt.Println(arr1)
}
