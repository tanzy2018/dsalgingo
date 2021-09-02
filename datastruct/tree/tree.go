package tree

import "fmt"

type CompareResult int

const (
	CompareResultLess CompareResult = iota
	CompareResultEqual
	CompareResultGreater
	CompareResultUndefined
)

var (
	// 空树
	ErrEmptyTree = NewError(0, "the nodes binary tree is empty")

	// 找不到节点
	ErrTargetNotFound = NewError(1, "the target is not found in binary tree")

	// 数据类型冲突
	ErrTypeConflict = NewError(2, "the type of the given target is conflict with the tree's type")

	// 未定义比较结果
	ErrUndefinedCompareResult = NewError(3, "the compare result is undefined")

	// 非法参数
	ErrIllegalParams = NewError(4, "the input param is illegal")
)

type Treer interface {
	Search(Comparabler) Noder
	Add(Comparabler)
	Del(Comparabler)
	Update(Comparabler)
	PostOrderTraverse() []Noder
	PreOrderTraverse() []Noder
	InOrderTraverse() []Noder
	LazyPostOrderTraverse() Noder
	HasNextPostOrderTraverse() bool
	ResetPostOrderTraverse()
	LazyPreOrderTraverse() Noder
	HasNextPreOrderTraverse() bool
	ResetPreOrderTraverse()
	LazyInOrderTraverse() Noder
	HasNextInOrderTraverse() bool
	ResetInOrderTraverse()
	Count() int
	Err() error
	Type() string
}

type Comparabler interface {
	Compare(target Comparabler) CompareResult
}

type Noder interface {
	Value() Comparabler
	Values() []Comparabler
}

type Error struct {
	Code int
	Msg  string
}

func (b Error) Error() string {
	return fmt.Sprintf("code=%d\t\tmsg=%s", b.Code, b.Msg)
}

func NewError(code int, msg string) Error {
	return Error{code, msg}
}
