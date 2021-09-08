package arithmetic

// 四则运算
// 二进制数字
type bInt = int

const (
	bIntZero bInt = 0
	bIntOne  bInt = 1
)

// 二进制半加器
func binaryhalfAdder(a, b bInt) (res, overflow bInt) {
	if a&b == bIntOne {
		return bIntZero, bIntOne
	}
	return a ^ b, bIntZero
}

// 二进制全加器 8位
func BinaryFullAdder8(a, b [8]bInt) (res [8]bInt, overflow bInt) {
	overflow = 0
	for i := 7; i >= 0; i-- {
		res, overflow_0 := binaryhalfAdder(a[i], b[i])
		res, overflow_1 := binaryhalfAdder(res, overflow)
		overflow, _ = binaryhalfAdder(overflow_0, overflow_1)
		a[i] = res
	}
	return a, overflow
}
