// @Description
// @Author 小游
// @Date 2021/04/06
package q35

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 解法一 哈希表
func copyRandomList1(head *Node) *Node {
	// 当我们的节点为空时返回空
	if head == nil {
		return nil
	}
	// 初始化map来存储数据
	link := map[*Node]*Node{}
	// 第一遍遍历我们先存储值
	cur := head
	for cur != nil {
		link[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	// 第二遍我们就可以开始进行复制了
	cur = head
	for cur != nil {
		// 这里我们就可以把我们存储的新节点的next和random都进行赋值了
		link[cur].Next = link[cur.Next]
		link[cur].Random = link[cur.Random]
		cur = cur.Next
	}
	// 返回我们复制好的对象
	return link[head]
}

// 解法二 拆分 原/新 链表
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		tmp := &Node{Val: cur.Val}
		tmp.Next = cur.Next
		cur.Next = tmp
		cur = tmp.Next
	}
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	cur = head.Next
	pre := head
	res := head.Next
	for cur.Next != nil {
		pre.Next = pre.Next.Next
		cur.Next = cur.Next.Next
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = nil
	return res
}
