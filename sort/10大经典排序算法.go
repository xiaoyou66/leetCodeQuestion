// @Description 10大经典排序算法go语言实现
// @Author 小游
// @Date 2021/03/09
package sort

// 插入排序
// 时间复杂度 n^2
// 因为我们这里是两层for循环，运气好的话，如果是有序的那么只需要执行一层，时间复杂度就是n
// 运气不好的话时间复杂度就是 n^2 平均的时间复杂度为 n^2
// 空间复杂度 1 ，因为我们只使用了一个临时变量，然后就没申请其他的了，所以空间复杂度为 1
func insertSort(arr []int){
	// 使用j来表示当前排序的位置
	var j = 0
	for i:=1;i<len(arr);i++{
		// 使用临时变量存储
		tmp:=arr[i]
		// 这个就是核心部分，首先我们让j处于i上
		// 当j-1大于tmp也就是arr[i]的时候,我们就需要把数字向后移动
		for j = i ; j>0 && arr[j-1] > tmp ; j--{
			arr[j] = arr[j-1]
		}
		// for循环的顺序 就是先j=1，然后执行判断语句，成功的话我们就执行函数体，最后才执行j--
		// 所以我们这里移动后，j实际上减了1，这样我们就可以直接替换了
		arr[j] = tmp
	}
}

// 折半插入排序（这个不算10大排序算法）
// 时间复杂度为 n^2，为什么是n^2呢，外面这一层不必多说就是n
// 里面这层，虽然有两个for循环，但是实际情况下，这个执行的时间是随n线性变化的，所以也可以看成n
// 空间复杂度为 1 这里为什么不是3呢，因为我们这个空间复杂度反映的是一个趋势
// 无论我们的n是多大，我们始终只用了三个变量，所以空间复杂度为1
func bInsertSort(arr []int){
	var j = 0
	for i:=1;i<len(arr);i++{
		// 存储临时变量
		tmp := arr[i]
		// 赋值low和high，这里high等于i-1
		// 这里为什么i-1，因为我们最后计算的时候要确保low指针+1不会越界
		low,high,mid:=0,i-1,0
		// 我们使用指针来查找需要插入的位置
		for low <= high {
			mid = (low+high) /2
			if arr[mid] > tmp {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		// 这里说一下为什么最后这个low就是我们应该放入的位置
		// low,high,mid的值是两个一组的，上一组是a[i]和a[mid]未经比较以前的，
		// 下一组是比较后，经过调整的。调整结果要么是low=mid+1，要么是high=mid-1。
		// 不知大家从上面的数据看出了什么。对！我们现在可以清楚的肯定：新元素的插入位置就是low。
		// 并且还发现high比low要小一，即high+1才与low相等。这是显然的，否则while循环如何结束。
		// 我想此刻大家的疑惑肯定是解开了。那么插入位置的写法就很随意了:low和high+1都行！
		// fmt.Println(low,mid,high)
		// 这里为什么要取low因为我们结束后j一定为是low，所以我们可以直接替换
		// 我也尝试过high不减一，low<high 但是就是调不出来，我也不知道为什么
		for j=i;j>low;j-- {
			arr[j] = arr[j-1]
		}
		arr[low] = tmp
	}
}

// 希尔排序
// 时间复杂度 n^1.3 这个是别人试验验证的结果
// 空间复杂度 不随n变化就是1
func shellSort(arr []int) {
	// 首先我们就是要确定增量大小，默认情况下，我们取n=len(arr)/2，后面我们的增量就会不断的缩小
	for gap:=len(arr)/2;gap>0;gap/=2{
		// 确定好增量后我们就从gap开始,这里为什么是i等于gap呢，因为我们后面是j-gap，所以我们需要从这个增量开始计算
		for i:=gap;i<len(arr);i++{
			// 用一个变量来暂存i的位置
			j:=i
			// 这里就开始比较了，我们是从gap开始的，j-gap 是我们比的第二个数字，然后我们进行交换
			for j-gap>=0 && arr[j] < arr[j-gap] {
				// 如果发现右边大于左边的，我们就交换一下
				arr[j],arr[j-gap] = arr[j-gap],arr[j]
				// 我们这里就是确保后面排序后再给前面排序
				j-=gap
			}
		}
	}
}

// 冒泡排序
// 时间复杂度 n^2 因为我们用了两层循环，这种两层循环的一般都是n^2
// 空间复杂度 1 这个和上面一样，不解释
func bubbleSort(arr []int)  {
	flag:=true
	// 首先我们直接设置i，i表示当前未排序的区间最大值,同时为了避免重复比较，设置一个flag来进行标记
	for i:=len(arr)-1;i>0 && flag;i-- {
		// 把flag置为false，如果没有排序那么我们就退出循环
		flag = false
		// 然后我们只需要按顺序排好i前面的就行了
		for j:=0;j<i;j++ {
			// 冒泡排序是比较相邻两个位置的元素
			if arr[j] > arr[j+1] {
				// flag置为true表示已经排过序了
				flag = true
				// 交换
				arr[j+1],arr[j] = arr[j],arr[j+1]
			}
		}
	}
}

// 快速排序-返回枢纽位置并进行排序，传入一个low指针，一个high指针，然后给数组进行简单排序
// 这里我们的数组排序好后返回一个中间位置的指针
func partition(arr []int,low int,high int) int {
	// 一般情况下我们直接取low点作为枢纽，然后我们就要进行排序，把比low小的排左边，比low大的排右边
	tmp:=arr[low]
	// 当low大于high时我们的排序就已经好了,这个时候其实low就是high同时这个值为我们的枢纽
	for low < high {
		// 首先我们从high哪里往左遍历，找出第一个比tmp小的点
		for low < high && arr[high] >= tmp {
			high --
		}
		// 找到这个比tmp小的点后我们就替换过去，这个时候我们的high其实留了一个空
		arr[low] = arr[high]
		// 因为上面high有一个空，所以这时候我们就可以从low指针开始出发，找到一个比tmp大的
		for low < high && arr[low] <= tmp {
			low ++
		}
		// 找到后我们就可以替换了，这个时候我们就把high给替换了，此时low就没用了
		// 然后我们就可以进行下一轮比较了，然后我们可以继续把比tmp小的点放到low这里
		arr[high] = arr[low]
	}
	// fmt.Println(low,high)
	// 最后替换完后，low和high其实就指向同一个位置了，这时我们就可以把low换成tmp了
	arr[low] = tmp
	// 返回low指针
	return low
}

// 快速排序
// 时间复杂度 nlogn 这东西推导比较复杂
// 具体参考 https://www.zhihu.com/question/22393997
// 空间复杂度 最好情况nlogn（注意这个log以2为底），最坏情况就是n
func quickSort(arr []int,low int,high int)  {
	// 这里也要判断，要不然会堆栈溢出
	if low < high {
		// 给数组进行简单排序，同时留下一个mid位置的指针
		mid:=partition(arr,low,high)
		// 这里我们使用递归来分别递归左右两部分
		quickSort(arr,low,mid-1)
		quickSort(arr,mid+1,high)
	}
}

// 选择排序
// 时间复杂度 n^2 这个很简单，因为是for循环嵌套
// 空间复杂度为1
func selectSort(arr []int){
	j:=0
	for i := 0; i < len(arr); i++ {
		// 记录最小值的位置
		min:=i
		// 找出最小的那个值，这里我们从i开始往后找，找到一个最小的
		for j=i+1;j<len(arr);j++{
			if arr[j] < arr[min]{
				min = j
			}
		}
		// 最小值进行替换一下
		arr[i],arr[min] = arr[min],arr[i]
	}
}

// 建初堆
func createHeap(arr []int)  {

}

// 调整堆
func adjustHeap(arr []int,s int,m int)  {

}

// 堆排序
func heapSort(arr []int)  {

}