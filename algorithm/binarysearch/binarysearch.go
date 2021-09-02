package binarysearch

// 二分法检索

func BinarySearchEqual(arr []int, num int) int {
	// 二分查找 指定值 ，返回数组下标，无则返回-1
	// 输入数组为空
	if len(arr) == 0 {
		return -1
	}
	// 数组中只有一个元素的简化判断
	if len(arr) == 1 {
		if arr[0] == num {
			return 0
		}
		return -1
	}
	// 第一次简化判断，首尾大小
	if arr[0] > num || arr[len(arr)-1] < num {
		return -1
	}

	// num可能在数组中
	left, right := 0, len(arr)-1

	var mid int
	for left <= right {
		mid = left + (right-left)/2
		// 找到目标值
		if arr[mid] == num {
			return mid
		}
		// 目标值可能在右边
		if arr[mid] < num {
			left = mid + 1
			continue
		}
		// 目标值可能在左边
		if arr[mid] > num {
			right = mid - 1
		}
	}
	return -1
}

func BinarySearchEqualFirst(arr []int, num int) (index int) {
	// 二分查找数组中第一个等于目标值的元素，并返回其数组下标
	index = -1
	end := len(arr)
	for end > 0 {
		end = BinarySearchEqual(arr[:end], num)
		if end == -1 {
			break
		}
		index = end
	}
	return index
}

func BinarySearchEqualLast(arr []int, num int) (index int) {
	// 二分查找数组中第一个等于目标值的元素，并返回其数组下标
	index = -1
	start := 0
	for start < len(arr) {
		start = BinarySearchEqual(arr[start:], num)
		if start == -1 {
			break
		}
		index += start + 1
		start = index + 1

	}
	return index
}

func BinarySearchLessLast(arr []int, num int) int {
	// 二分查找数组中第一个小于目标值的元素，并返回其数组下标
	// 空数组
	if len(arr) == 0 {
		return -1
	}
	// 一个元素的数组
	if len(arr) == 1 && arr[0] < num {
		return 0
	}
	// 简单判断
	if arr[0] >= num {
		return -1
	}

	// 找出第一个小于目标值的元素
	left, right := 0, len(arr)-1
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if arr[mid] >= num {
			right = mid - 1
			continue
		}
		if arr[mid] < num {
			return mid + BinarySearchLessLast(arr[mid+1:right+1], num) + 1
		}
	}
	return -1

}

func BinarySearchGreaterFirst(arr []int, num int) (index int) {
	// 二分查找数组中第一个大于目标值的元素，并返回其数组下标
	// 空数组
	if len(arr) == 0 {
		return -1
	}
	// 一个元素
	if len(arr) == 1 && arr[0] > num {
		return 0
	}
	// 简易判断
	if arr[len(arr)-1] <= num {
		return -1
	}
	// 找到第一个大于目标值的元素
	var mid int
	left, right := 0, len(arr)-1
	for left <= right {
		mid = left + (right-left)/2
		if arr[mid] <= num {
			left = mid + 1
			continue
		}
		if arr[mid] > num {
			index := BinarySearchGreaterFirst(arr[left:mid], num)
			if index == -1 {
				return mid
			}
			return mid - (mid - left) + index
		}
	}
	return -1
}
