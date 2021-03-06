// @Description 
// @Author 小游
// @Date 2021/03/09
package sort

import (
	"fmt"
	"testing"
)

var arr = []int{38,65,97,76,13,27,49}
// [13 27 38 49 65 76 97]


func TestInsert(t *testing.T){
	insertSort(arr)
	fmt.Println(arr)
}

func TestBInsert(t *testing.T){
	bInsertSort(arr)
	fmt.Println(arr)
}
func TestShellSort(t *testing.T){
	shellSort(arr)
	fmt.Println(arr)
}

func TestBubbleSort(t *testing.T){
	bubbleSort(arr)
	fmt.Println(arr)
}

func TestQuickSort(t *testing.T){
	quickSort(arr,0,len(arr)-1)
	fmt.Println(arr)
}

func TestSelect(t *testing.T){
	selectSort(arr)
	fmt.Println(arr)
}

func TestHeap(t *testing.T){
	heapSort(arr)
	fmt.Println(arr)
}

func TestMSort(t *testing.T){
	arr2:=make([]int,len(arr))
	mSort(arr,arr2,0,len(arr)-1)
	fmt.Println(arr)
	fmt.Println(arr2)
}

func TestCountSort(t *testing.T){
	countingSort(arr,13,97)
	fmt.Println(arr)
}

func TestBucketSort(t *testing.T){
	sortBucket(arr,97,5)
	fmt.Println(arr)
}

func TestRadixSort(t *testing.T){
	radixSort(arr,97)
	fmt.Println(arr)
}