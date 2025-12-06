package Sorts

import (
	"math/rand"
	"time"
)

// 对arr[l...r]部分进行partition操作
// 返回p, 使得arr[l...p-1] < arr[p] ; arr[p+1...r] > arr[p]
func partition (arr []int,l,r int) int {
	// 随机在arr[l...r]的范围中, 选择一个数值作为标定点pivot
	rand.Seed(time.Now().Unix())
	randIndex := rand.Int()%(r-l+1)+l
	arr[l],arr[randIndex] = arr[randIndex],arr[l]

	v := arr[l]
	j := l
	for i := l+1; i<=r; i++ {
		if arr[i] < v {
			j++
			arr[j],arr[i] = arr[i],arr[j]
		}
	}
	arr[l],arr[j] = arr[j],arr[l]
	return j
}

// 双路快速排序的partition
// 返回p, 使得arr[l...p-1] < arr[p] ; arr[p+1...r] > arr[p]
func partition2(arr []int,l,r int) int {
	rand.Seed(time.Now().Unix())
	randIndex := rand.Int()%(r-l+1)+l
	arr[l],arr[randIndex] = arr[randIndex],arr[l]

	v := arr[l]
	// 不变式：arr[l+1...i) <= v; arr(j...r] >= v
	i,j := l+1,r
	for {
		// 注意这里的边界, arr[i] < v, 不能是arr[i] <= v
		// 思考一下为什么?
		for i <= r && arr[i] < v {
			i++
		}
		// 注意这里的边界, arr[j] > v, 不能是arr[j] >= v
		// 思考一下为什么?
		for j >= l + 1 && arr[j] > v {
			j--
		}
		if i >  j {
			break
		}
		arr[j],arr[i] = arr[i],arr[j]
		i++
		j--
	}
	arr[l],arr[j] = arr[j],arr[l]
	return j
}
// 递归的三路快速排序算法
func quickSort3Ways(arr []int,l,r int)  {
	if l >= r {
		return
	}
	// 对于小规模数组, 我们其实可以使用插入排序进行优化
	// 随机在arr[l...r]的范围中, 选择一个数值作为标定点pivot
	rand.Seed(time.Now().Unix())
	randIndex := rand.Int()%(r-l+1)+l
	arr[l],arr[randIndex] = arr[randIndex],arr[l]

	v := arr[l]
	// 小于区间：arr[l+1...lt] < v
	// 大于区间：arr[gt...r] > v
	// 等于区间：arr[lt+1...i) == v
	lt,gt,i := l,r+1,l+1
	for i < gt {
		if arr[i] < v {
			arr[i],arr[lt + 1] = arr[lt + 1],arr[i]
			i++
			lt++
		} else if arr[i]>v {
			arr[i],arr[gt-1] = arr[gt-1],arr[i]
			gt--
		} else {
			i++
		}
	}
	arr[l],arr[lt] = arr[lt],arr[l]
	quickSort3Ways(arr,l,lt-1)
	quickSort3Ways(arr,gt,r)
}

func QuickSort(arr []int,l,r int)  {
	if l >= r {
		return
	}
	p := partition(arr,l,r)
	QuickSort(arr,l,p-1)
	QuickSort(arr,p+1,r)
}

func Quick3Ways(arr []int,n int)  {
	quickSort3Ways(arr,0,n-1)
}
