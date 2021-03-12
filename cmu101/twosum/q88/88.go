// @Description 
// @Author 小游
// @Date 2021/03/12
package q88

// 我的解法，这方法不行。。占用内存太大了
func merge1(nums1 []int, m int, nums2 []int, n int)  {
	// 这里我们同样使用双指针
	if n == 0 {
		return
	}
	// 这里我为了方便在创建一个数组
	var re []int
	// 使用两个指针来表示两个数组
	i,j:=0,0
	// 下面我们开始遍历
	for i < m && j < n {
		if  nums1[i] > nums2[j] {
			re = append(re,nums2[j])
			j++
		} else {
			re = append(re,nums1[i])
			i++
		}
	}
	// 合并剩下的部分
	for ;i<m;i++{
		re = append(re,nums1[i])
	}
	for ;j<n;j++{
		re = append(re,nums2[j])
	}
	copy(nums1,re)
}

// 解法2 使用三个指针
func merge(nums1 []int, m int, nums2 []int, n int)  {
	// 这里其实很简单，我们的数组已经排好序了，正的不好排，我们可以从大到小排
	// 合理利用逆向思维,可以让我们的题目变简单
	pos:=m+n-1;m--;n--
	for m>=0 && n>=0{
		// 我们把最大的那个值插入num1的末尾
		if nums2[n] > nums1[m] {
			nums1[pos] = nums2[n]
			n--
		} else {
			nums1[pos] = nums1[m]
			m--
		}
		pos --
	}
	// 遍历剩余部分
	for n >= 0{
		nums1[pos] = nums2[n]
		n--;pos--
	}
	for m >= 0{
		nums1[pos] = nums1[m]
		m--;pos--
	}
}
