// @Description
// @Author 小游
// @Date 2021/03/06
package main

import (
	"fmt"
)

func main() {
	a := 1
	defer fmt.Println("the value of a1:",a)
	defer func() {
		fmt.Println("the value of a2:",a)
	}()
	a++
}