// @Description
// @Author 小游
// @Date 2021/04/06
package q31

import "container/list"

func validateStackSequences(pushed []int, popped []int) bool {
	// 初始化一个栈，模拟
	stack := list.New()
	// 声明节点表示位置
	i := 0
	// 遍历整个pushed节点
	for _, v := range pushed {
		// 先把当前节点放入栈中
		stack.PushBack(v)
		// 判断栈顶元素是否和popped的元素相同，相同我们就出栈
		for stack.Len() != 0 && stack.Back().Value.(int) == popped[i] {
			stack.Remove(stack.Back())
			i++
		}
	}
	// 最后如果栈为空，就说明我们遍历完毕了
	return stack.Len() == 0
}
