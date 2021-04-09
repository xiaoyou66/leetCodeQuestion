// @Description
// @Author 小游
// @Date 2021/04/09
package q59

type MaxQueue struct {
	// 这里我们维护了一个正常的队列
	q []int
	// 这里我们则维护了一个递减的队列
	p []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	// 当我们想获取最大值时
	if len(this.q) == 0 {
		return -1
	}
	// 这里我们可以直接返回递减的队列的第一个元素
	return this.p[0]
}

func (this *MaxQueue) Push_back(value int) {
	// 正常队列我们可以直接加入
	this.q = append(this.q, value)
	// 这里我们把队尾小于v的元素全部弹出
	for len(this.p) > 0 && value > this.p[len(this.p)-1] {
		this.p = this.p[:len(this.p)-1]
	}
	// 然后把value插进来
	this.p = append(this.p, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.q) == 0 {
		return -1
	}
	// 当我们出队是需要判断最大值是否为递减队列的最大致，如果是那么递减队列也移出
	if this.q[0] == this.p[0] {
		this.p = this.p[1:]
	}
	value := this.q[0]
	this.q = this.q[1:]
	return value
}
