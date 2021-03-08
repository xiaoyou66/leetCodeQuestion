// @Description 
// @Author 小游
// @Date 2021/03/07
package q013

// 解法1
//func romanToInt(s string) int {
//	// 初始化map存储信息
//	info:=map[string]int{"I":1,"V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}
//	sum:=0
//	// 遍历string
//	for i:=0;i<len(s);i++ {
//		num:=0
//		// 首先判断特殊情况
//		if i < len(s)-1{
//			two:=string(s[i])+string(s[i+1])
//			switch two {
//			case "IX":
//				num = 9
//			case "IV":
//				num = 4
//			case "XL":
//				num = 40
//			case "XC":
//				num = 90
//			case "CD":
//				num = 400
//			case "CM":
//				num = 900
//			}
//			if num!=0 {
//				i ++
//				sum += num
//				continue
//			}
//		}
//		sum += info[string(s[i])]
//	}
//	return sum
//}

// 解法二
var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	num, lastint, total := 0, 0, 0
	// 遍历s
	for i := 0; i < len(s); i++ {
		// 获取char（这里获取是从后面往前获取的）
		char := s[len(s)-(i+1) : len(s)-i]
		// 计算num
		num = roman[char]
		// 如果这个值比上一次的要小，那么我们就减一下，否则就加一下
		if num < lastint {
			total = total - num
		} else {
			total = total + num
		}
		lastint = num
	}
	return total
}