// @Description 10大经典排序算法go语言实现
// @Author 小游
// @Date 2021/03/09
package sort

import (
	"math"
)

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
	// 这里我们使用筛选法来建初堆，我们依次把、[n/2],[n/2]-1,....,1的节点都调整为堆即可
	n := len(arr) - 1
	// 根据堆的性质，在完全二叉树中所有序号大于 [n/2]的节点都是叶子，
	// 所以我们只需要调整 1 - n/2 这个区间把这几个点调整为大根堆就行
	// 这里调整堆传入了一个开始点和结束点
	for i :=n/2;i>=0;i--{
		adjustHeap(arr,i,n)
	}
}

// 调整堆
func adjustHeap(arr []int,s int,m int)  {
	// 保存起始点的值作为临变量
	rc:=arr[s]
	// 我们使用筛选法来调整堆
	// 这里使用了堆的性质
	// 当 ki>=k2i 且 ki >= k2i+1 或 ki <=k2i 且 ki <= k2i+1
	for j:=2*s;j<=m;j*=2{
		if j<m && arr[j]<arr[j+1]{ j++ }
		if rc >= arr[j] { break}
		arr[s] = arr[j]
		s=j
	}
	arr[s] = rc
}

// 堆排序
// 时间复杂度 nlogn
// 这里耗时操作在于筛选和调整 筛选需要n-1次 然后n个节点的二叉树深度为logn +1
// 空间复杂度 1 因为只需要一个临时空间
func heapSort(arr []int)  {
	// 先建堆，保证初始的时候是一个堆
	createHeap(arr)
	// 开始进行for循环，进行堆排序
	for i:=len(arr)-1;i>0;i--{
		// 这里我们把最大的那个最后一个（这个指的是当前堆的最大一个）互换
		// 因为我们堆排序后0号是最大的
		arr[0],arr[i] = arr[i],arr[0]
		// 我们i号是最大的了，所以我们忽略i，从i-1开始进行调整
		adjustHeap(arr,0,i-1)
	}
}

// 合并表
// 注意这里我们arr2是一个空数组，然后arr1的low和mid是一个排好序的数组，mid+1和high也是一个排好序的数组
// 这个函数的作用就是把两个数组合并为一个数组，放入arr2中
func Merge(arr1 []int,arr2 []int,low int,mid int,high int)  {
	// 这里我们使用i和j来表示arr1数组的两部分，然后k就表示arr2数组的位置
	i,j,k:=low,mid+1,low
	// 只要有一个数组到头了，我们就介绍排序
	for i <= mid && j<=high {
		// 因为我们要合并两个数组，所以我们需要判断一下数组的两部分到底谁大
		if arr1[i] > arr1[j] {
			// 注意，k表示arr2的位置，然后j表示的是 mid+1 到high这个数组
			arr2[k] = arr1[j];k++;j++
		} else {
			// 注意，k表示arr2的位置，然后j表示的是 low 到mid 这个数组
			arr2[k] = arr1[i];k++;i++
		}
	}
	// 剩余部分因为已经是一个有序的数组了，所以我们直接按顺序加到arr2就行
	for i<=mid {
		arr2[k] = arr1[i];i++;k++
	}
	for j<=high {
		arr2[k] = arr1[j];j++;k++
	}
}

// 2-路归并排序
// 时间复杂度nlogn
// 空间复杂度 n 因为我们排序的时候需要申请n个空间来存储
func mSort(arr []int,arr2 []int,low int,high int)  {
	// 当low和high相等的时候，这个时候大小为1，我们直接放到arr2就行
	if low == high{
		arr2[low] = arr[low]
	} else {
		// 我们准备开始进行拆分,我们需要把我们的数组拆成两部分
		arr3:=make([]int,high+1)
		// 直接取mid作为数组的分界点
		mid:= (low + high) / 2
		// 首先我们把low和mid 进行排序并放入arr3中
		mSort(arr,arr3,low,mid)
		// 然后我们把mid+1和high 进行排序并放入arr3中
		mSort(arr,arr3,mid+1,high)
		// 最后我们得到的数组在arr3中已经排好序并分成了两部分
		// 因为arr2是一个空数组，所以我们把arr3再进行排序，最后就是合并为一个有序的表并放入arr2中
		Merge(arr3,arr2,low,mid,high)
	}
}

// 计数排序,这里我们传入一个数组，一个最小值，一个最大值
// 时间复杂度 n+k
// 空间复杂度 n+k
func countingSort(arr []int, minvalue, maxValue int){
	// 首先我们建立一个统计数组,为了节省空间我们直接取最大值和最小值的差+1的空间，这个时候最大值在最后一个，最小值在第一个
	count:=make([]int,maxValue-minvalue+1)
	// 遍历数组，统计出现的次数，注意我们计算的时候需要减minValue
	for _,v:=range arr{
		count[v-minvalue] ++
	}
	// 最后我们输出数组，这里我们使用i来标记arr数组
	i:=0
	// 遍历count
	for k,v:=range count{
		// 如果count数组为0那么就不管，如果不为0那么就为几，我们就输出几次
		if v!=0 {
			// 输出v次
			for j:=0;j<v;j++{
				arr[i] = minvalue + k
				i ++
			}
		}
	}
}

// 桶排序，这里我们传入一个最大值，还有一为个桶的数量
// 时间复杂度 为n+k
// 空间复杂度为 n+k
func sortBucket(arr []int,max int,num int) {
	// 创建桶，大小为我们传入桶的数量
	buckets := make([][]int,num)
	var index int
	// 遍历数组
	for _,v := range arr{
		// 分配桶 index = value * (n-1)/k 这里我们使用公式来确放入的位置
		index = v * (num-1) / max
		// 把我们的数字放入桶内
		buckets[index] = append(buckets[index],v)
	}
	// 桶内排序
	tmpPos := 0
	// 这里我们遍历每一个桶
	for k:=0; k < num; k++ {
		// 获取桶里面数据的长度
		bucketLen := len(buckets[k])
		// 如果我们这个桶大小为0，那么我们就对桶里面进行排序
		if bucketLen>0{
			// 使用插入排序对数组进行排序
			insertSort(buckets[k])
			// 我们把排序好的部分复制到数组中
			copy(arr[tmpPos:],buckets[k])
			// 数组位置加上桶的大小，开始保存下一个桶
			tmpPos +=bucketLen
		}
	}
}

// 基数排序，这里我们传入数组和这个数组的最大值
func radixSort(theArray []int,max int){
	// 因为前面我们获取了最大值，这里我们获取一下最大值的位数
	var count = 0
	for max % 10>0{
		max /= 10
		count++
	}
	// 给桶中对应的位置放数据
	for i:=0; i<count; i++ {
		// 根据不同的位数，我们来获取10的n次方，用于后面获取每位的值
		theData := int(math.Pow10(i))
		// 建立空桶
		bucket:=make([][]int,10)
		// 遍历数组
		for k:=0; k<len(theArray); k++{
			// 这里我们进行取余操作，目的是为了获取这个数组在每位上的值
			theResidue := (theArray[k]/theData) %10
			// 获取到了位数后我们直接放入对应的桶中
			bucket[theResidue] = append(bucket[theResidue],theArray[k])
		}

		// 一遍循环完之后需要把数组二维数据进行重新排序，比如数组开始是10 1 18 30 23 12 7 5 18 233 144 ，循环个位数
		// 循环之后的结果为10 30 1 12 23 233 144 5 7 18 18 ，然后循环十位数，结果为1 5 7 10 12 18 18 23 30 233 144
		// 最后循环百位数，结果为1 5 7 10 12 18 18 23 30 144 233
		var x = 0
		// 遍历我们的桶
		for p:=0; p<len(bucket); p++ {
			for q:=0; q<len(bucket[p]); q++ {
				// 这里按照桶的结构，按桶的顺序把桶里面的数据放入数组中
				if bucket[p][q]!=0 {
					theArray[x] = bucket[p][q]
					x++
				}else {
					break
				}
			}
		}
	}
}
