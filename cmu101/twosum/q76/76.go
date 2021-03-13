// @Description 
// @Author 小游
// @Date 2021/03/13
package q76

func minWindow(s string, t string) string {
	// 当s和t都为空时我们返回空字符串
	if s == "" || t == "" {
		return ""
	}
	// 使用两个数组来存储t和s字符串出现的次数
	var tCount,sCount [256]int
	// result表示最后的结果
	// left表示滑动窗口的左指针，right表示滑动窗口的右指针
	// finalLeft表示最后的左指针，finalRight表示最后的右指针，我们通过这两个指针得到最后的结果
	// minW表示滑动窗口最小的宽度
	// count表示
	result,left,right,finalLeft,finalRight,minW,count:="",0,-1,-1,-1,len(s)+1,0
	// 先统计t字符串里面每个字符出现的次数
	for _,v:=range t{
		tCount[v] ++
	}
	// 遍历整个s
	for left < len(s) {
		// 首先我们需要让右指针进行遍历，直到找出完全包含T的字符串时停止
		if right+1 < len(s) && count < len(t) {
			right++
			// 统计一下字符
			sCount[s[right]]++
			// 这里当sCount比tCount小时，我们就让count++
			// 只要当前我们遍历所在位置的字符比t小时，我们就让count+1，否则不加，避免重复计数
			if sCount[s[right]] <= tCount[s[right]] {
				count ++
			}
		} else {
		// 经过上面的操作，这我们的右指针就已经包含所有的t了
		// 这个时候我们就需要更新左指针，来获取最小字符串

			// 当我们的左右指针小于最小窗口值，同时又满足包含t字符时
			if right-left+1<minW && count == len(t) {
				// 这里我们就可以更新一下最小窗口了
				minW = right - left +1
				// 同时我们更新一下最后的左指针和右指针
				finalLeft = left
				finalRight = right
			}
			// 当left指针所指向的字符串包含t中的字符串时，我们就必须要把count-1操作
			// 因为我们的窗口要包含t所有的子串
			if sCount[s[left]] == tCount[s[left]]{
				count--
			}
			// sCount-1操作，说明我们的左指针移动了一位
			sCount[s[left]]--
			left ++
		}
	}
	// 当我们最后的左指针不为-1时，就说明我们找到了相等的子串
	if finalLeft!=-1 {
		// 这里我们把我们的结果贴上去
		result = s[finalLeft:finalRight+1]
	}
	return result
}