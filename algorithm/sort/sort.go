package sort

import (
	"math"
)

// TODO::排序

// 快排
func Qsort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return []int{arr[1], arr[0]}
		}
		return arr
	}
	qsort(arr)
	return arr
}

func qSortPartition(arr []int) int {
	// 小于2个元素
	// if len(arr) <= 1 {
	// 	return 0
	// }
	// 2个元素
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
		return 0
	}
	// 3个元素以上
	arr[0], arr[len(arr)/2] = arr[len(arr)/2], arr[0]
	i, j := 1, len(arr)-1
	for i <= j {
		if arr[i] <= arr[0] {
			i++
			continue
		}
		// 避免无效数据交换
		for ; arr[j] >= arr[0] && j >= i; j-- {
		}
		// 有效数据交换
		if j > i {
			arr[i], arr[j] = arr[j], arr[i]
			j--
			i++
		}
	}
	arr[0], arr[j] = arr[j], arr[0]
	return j
}

func qsort(arr []int) {
	mid := qSortPartition(arr)
	// 至少2个才有意义
	if mid > 1 {
		qsort(arr[:mid])
	}
	// 至少2个才有意义
	if mid < len(arr)-2 {
		qsort(arr[mid+1:])
	}
}

// 归并排序
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return []int{arr[1], arr[0]}
		}
		return arr
	}
	return mergeSort(arr)
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			arr[1], arr[0] = arr[0], arr[1]
		}
		return arr
	}
	return merge(mergeSort(arr[:len(arr)/2]), mergeSort(arr[len(arr)/2:]))
}

func merge(a, b []int) []int {
	if len(b) == 0 {
		return a
	}
	if len(a) == 0 {
		return b
	}
	c := make([]int, len(a)+len(b))
	for ai, bi, ci := 0, 0, 0; ci < len(c); ci++ {
		// 只剩下b
		if ai >= len(a) {
			copy(c[ci:], b[bi:])
			return c
		}
		// 只剩下a
		if bi >= len(b) {
			copy(c[ci:], a[ai:])
			return c
		}
		if a[ai] <= b[bi] {
			c[ci] = a[ai]
			ai++
		} else {
			c[ci] = b[bi]
			bi++
		}
	}
	return c
}

// 插入排序
func InsertionSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	if len(arr) == 2 {
		if arr[0] <= arr[1] {
			return arr
		}
		return []int{arr[1], arr[0]}
	}
	for i := 1; i < len(arr); i++ {
		j := i - 1
		for ; j >= 0 && arr[i] < arr[j]; j-- {
		}
		// 有效数据交换
		if j != i-1 {
			// 只需要交换相邻元素
			if j == i-2 {
				arr[j+1], arr[i] = arr[i], arr[j+1]
				continue
			}
			tmp := arr[i]
			copy(arr[j+2:], arr[j+1:i])
			arr[j+1] = tmp
		}
	}
	return arr
}

func CounterSort(arr []int) []int {
	// 空数组或者一个元素
	if len(arr) < 2 {
		return arr
	}
	// 两个元素
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
		return arr
	}
	positiveCount, negativeCount := 0, 0
	positiveMin, positiveMax := math.MaxInt64, -1
	negativeMin, negativeMax := -1, math.MinInt64
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 0 {
			positiveCount++
			if arr[i] < positiveMin {
				positiveMin = arr[i]
			}
			if arr[i] > positiveMax {
				positiveMax = arr[i]
			}
		} else {
			negativeCount++
			if arr[i] > negativeMax {
				negativeMax = arr[i]
			}
			if arr[i] < negativeMin {
				negativeMin = arr[i]
			}
		}
	}
	// 原地放回原来的数组
	aIndex := 0
	// 先处理负数
	var negativeArr []int
	if positiveCount == 0 {
		negativeArr = arr
	} else {
		// 避免overlay
		negativeArr = make([]int, negativeCount)
	}
	if negativeCount > 0 {
		counter := make([]int, negativeMax-negativeMin+1)
		for i := 0; i < len(arr); i++ {
			if arr[i] < 0 {
				counter[-(arr[i]-negativeMax)]++
			}
		}
		for i := len(counter) - 1; i >= 0; i-- {
			val := negativeMax - i
			for j := 0; j < counter[i]; j++ {
				negativeArr[aIndex] = val
				aIndex++
			}
		}

	}
	// 再处理正数
	if positiveCount > 0 {
		counter := make([]int, positiveMax-positiveMin+1)
		for i := 0; i < len(arr); i++ {
			if arr[i] >= 0 {
				counter[arr[i]-positiveMin]++
			}
		}
		for i := 0; i < len(counter); i++ {
			if counter[i] > 0 {
				val := i + positiveMin
				for j := 0; j < counter[i]; j++ {
					arr[aIndex] = val
					aIndex++
				}
			}
		}
	}
	if negativeCount > 0 {
		copy(arr[:negativeCount], negativeArr)
	}
	return arr
}
