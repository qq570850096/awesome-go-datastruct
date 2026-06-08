package Sorts

import (
	"math/rand"
	"time"
)

func partition(arr []int, l, r int) int {

	rand.Seed(time.Now().Unix())
	randIndex := rand.Int()%(r-l+1) + l
	arr[l], arr[randIndex] = arr[randIndex], arr[l]

	v := arr[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func partition2(arr []int, l, r int) int {
	rand.Seed(time.Now().Unix())
	randIndex := rand.Int()%(r-l+1) + l
	arr[l], arr[randIndex] = arr[randIndex], arr[l]

	v := arr[l]

	i, j := l+1, r
	for {

		for i <= r && arr[i] < v {
			i++
		}

		for j >= l+1 && arr[j] > v {
			j--
		}
		if i > j {
			break
		}
		arr[j], arr[i] = arr[i], arr[j]
		i++
		j--
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func quickSort3Ways(arr []int, l, r int) {
	if l >= r {
		return
	}

	rand.Seed(time.Now().Unix())
	randIndex := rand.Int()%(r-l+1) + l
	arr[l], arr[randIndex] = arr[randIndex], arr[l]

	v := arr[l]

	lt, gt, i := l, r+1, l+1
	for i < gt {
		if arr[i] < v {
			arr[i], arr[lt+1] = arr[lt+1], arr[i]
			i++
			lt++
		} else if arr[i] > v {
			arr[i], arr[gt-1] = arr[gt-1], arr[i]
			gt--
		} else {
			i++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]
	quickSort3Ways(arr, l, lt-1)
	quickSort3Ways(arr, gt, r)
}

func QuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	QuickSort(arr, l, p-1)
	QuickSort(arr, p+1, r)
}

func Quick3Ways(arr []int, n int) {
	quickSort3Ways(arr, 0, n-1)
}
