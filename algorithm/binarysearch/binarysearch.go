package binarysearch

// 二分法检索

// 二分法检索本质：在一个有序集合里找到目标（存在性问题），这个目标表达为一个等式或者不等式，
// 而找到第一个XXX或者最后一个XXX，是对”存在性问题“的向左或者向右迭代，所以它们本质是一样的。
// 根据集合的搜索模式，结合具体实现，可以发现以下结论：
// 1. BinarySearchXXX (XXX: Equal | Less | Greater | LessOrEqual) 共用同一套模板,解决（存在性）问题
// 2. BinarySearchXXXFirst (XXX: Equal | Less | Greater | LessOrEqual) 共用同一套模板，解决第一个（存在性）问题
// 3. BinarySearchXXXFirst (XXX: Equal | Less | Greater | LessOrEqual) 共用同一套模板，解决最后一个（存在性）问题
func BinarySearchEqual(arr []int, num int) int {
	// 二分查找 指定值 ，返回数组下标，无则返回-1
	// 空数组
	if len(arr) == 0 {
		return -1
	}
	// 1个元素
	if len(arr) == 1 {
		if arr[0] == num {
			return 0
		}
		return -1
	}
	// 不可能存在 边界快速判定
	if arr[0] > num || arr[len(arr)-1] < num {
		return -1
	}
	// 可能存在
	left, right := 0, len(arr)-1 // 闭区间可以避免边界溢出
	for left <= right {
		mid := left + (right-left)/2 // 还是为了防止溢出
		// 找到
		if arr[mid] == num {
			return mid
		}
		// 可能在左边
		if arr[mid] > num {
			right = mid - 1
			continue
		}
		// 可能在右边
		if arr[mid] < num {
			left = mid + 1
		}

	}
	return -1
}

func BinarySearchEqualFirst(arr []int, num int) (index int) {
	// 二分查找数组中第一个等于目标值的元素，并返回其数组下标
	// 一直往左边找
	index = -1
	end := len(arr)
	for end > 0 {
		if end = BinarySearchEqual(arr[:end], num); end == -1 {
			break
		}
		index = end
	}
	return index
}

func BinarySearchEqualLast(arr []int, num int) int {
	// 二分查找数组中第一个等于目标值的元素，并返回其数组下标
	// 一直往右边找
	index := -1
	start := 0
	for start < len(arr) {
		if start = BinarySearchEqual(arr[start:], num); start == -1 {
			break
		}
		index += start + 1
		start = index + 1
	}
	return index
}

func BinarySearchLessLast(arr []int, num int) int {
	// 二分查找数组中最后一个小于目标值的元素，并返回其数组下标
	// 如果存在，一定在目标值的左边
	// 空数组
	if len(arr) == 0 {
		return -1
	}
	// 1个元素
	if len(arr) == 1 {
		if arr[0] < num {
			return 0
		}
		return -1
	}
	// 不可能存在 边界简易判断
	if arr[0] >= num {
		return -1
	}
	// 一定存在
	left, right := 0, len(arr)-1
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if arr[mid] < num {
			break
		}
		right = mid - 1
	}
	// 可能右边有更大的比目标值小的元素 , 如果右边不存在，右边区间返回-1, 则mid就是所找下标 mid = -1 + mid + 1
	return BinarySearchLessLast(arr[mid+1:right+1], num) + mid + 1
}

func BinarySearchGreaterFirst(arr []int, num int) int {
	// 二分查找数组中第一个大于目标值的元素，并返回其数组下标
	// 如果存在 一定在目标值的右边
	// 空数组
	if len(arr) == 0 {
		return -1
	}
	// 1个元素
	if len(arr) == 1 {
		if arr[0] > num {
			return 0
		}
		return -1
	}
	// 不存在
	if arr[len(arr)-1] <= num {
		return -1
	}
	// 一定存在
	var mid int
	left, right := 0, len(arr)-1
	for left <= right {
		mid = left + (right-left)/2
		// 一定在右边
		if arr[mid] > num {
			break
		}
		left = mid + 1

	}
	// 左边可能还有比更小的且比目标大的元素
	if index := BinarySearchGreaterFirst(arr[left:mid], num); index != -1 {
		return left + index
	}
	return mid
}

func BinarySearchLess(arr []int, num int) int {
	// 二分查找 小于指定值的元素 ，返回数组下标，无则返回-1
	// 空数组
	if len(arr) == 0 {
		return -1
	}
	// 1个元素
	if len(arr) == 1 {
		if arr[0] < num {
			return 0
		}
		return -1
	}
	// 不存在
	if arr[0] >= num {
		return -1
	}
	// 一定存在
	var mid int
	left, right := 0, len(arr)-1
	for left <= right {
		mid = left + (right-left)/2
		if arr[mid] < num {
			break
		}
		right = mid - 1

	}
	return mid
}

func BinarySearchLessLast_0(arr []int, num int) int {
	// 二分查找数组中最后一个小于目标值的元素，并返回其数组下标
	// 如果存在，一定在右边
	index := -1
	start := 0
	for start < len(arr) {
		if start = BinarySearchLess(arr[start:], num); start == -1 {
			break
		}
		index += start + 1
		start = index + 1
	}
	return index
}

func BinarySearchGreater(arr []int, num int) int {
	// 二分查找大于指定值的元素 ，返回数组下标，无则返回-1
	// 空数组
	if len(arr) == 0 {
		return -1
	}
	// 1个元素
	if len(arr) == 1 {
		if arr[0] > num {
			return 0
		}
		return -1
	}
	// 不存在
	if arr[len(arr)-1] <= num {
		return -1
	}
	// 一定存在
	left, right := 0, len(arr)-1
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if arr[mid] > num {
			break
		}
		left = mid + 1
	}
	return mid
}

func BinarySearchGreaterFirst_0(arr []int, num int) int {
	// 二分查找数组中第一个大于目标值的元素，并返回其数组下标
	// 如果存在 一定在左边
	index := -1
	end := len(arr)
	for end > 0 {
		if end = BinarySearchGreater(arr[:end], num); end == -1 {
			break
		}
		index = end
	}
	return index
}

func BinarySearchLessOrEqual(arr []int, num int) int {
	// 二分查找 小于等于指定值的元素 ，返回数组下标，无则返回-1
	// 空数组
	if len(arr) == 0 {
		return -1
	}
	// 1个元素
	if len(arr) == 1 {
		if arr[0] <= num {
			return 0
		}
		return -1
	}
	// 不存在
	if arr[0] > num {
		return -1
	}
	// 一定存在
	var mid int
	left, right := 0, len(arr)-1
	for left <= right {
		mid = left + (right-left)/2
		if arr[mid] <= num {
			break
		}
		right = mid - 1
	}
	return mid
}

func BinarySearchLessOrEqualLast(arr []int, num int) int {
	// 二分查找数组中最后一个小于等于目标值的元素，并返回其数组下标
	// 如果存在，一定在右边
	index := -1
	start := 0
	for start < len(arr) {
		if start = BinarySearchLessOrEqual(arr[start:], num); start == -1 {
			break
		}
		index += start + 1
		start = index + 1
	}
	return index
}

func BinarySearchGreaterOrEqual(arr []int, num int) int {
	// 二分查找大于等于指定值的元素 ，返回数组下标，无则返回-1
	// 空数组
	if len(arr) == 0 {
		return -1
	}
	// 1个元素
	if len(arr) == 1 {
		if arr[0] >= num {
			return 0
		}
		return -1
	}
	// 不存在
	if arr[len(arr)-1] < num {
		return -1
	}
	// 一定存在
	left, right := 0, len(arr)-1
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if arr[mid] >= num {
			break
		}
		left = mid + 1
	}
	return mid
}

func BinarySearchGreaterOrEqualFirst(arr []int, num int) int {
	// 二分查找数组中第一个大于目标值的元素，并返回其数组下标
	// 如果存在 一定在左边
	index := -1
	end := len(arr)
	for end > 0 {
		if end = BinarySearchGreaterOrEqual(arr[:end], num); end == -1 {
			break
		}
		index = end
	}
	return index
}
