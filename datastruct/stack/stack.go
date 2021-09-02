package stack

import "errors"

type Stacker interface {
	Pop() (interface{}, error)
	Push(interface{}) error
	Clear()
	Len() int
	Cap() int
}

var (
	// 栈中元素达到容量上限
	ErrStackIsFull = errors.New("stack is up to it's capcity")
	// 栈中元素为空
	ErrStackIsEmpty = errors.New("stack is empty")
)

// 非线程安全的栈
type Stack struct {
	len   int
	cap   int
	stack []interface{}
}

func NewStack(cap int) *Stack {
	if cap <= 0 {
		cap = 64
	}
	return &Stack{
		len:   0,
		cap:   cap,
		stack: make([]interface{}, 0, cap),
	}
}

func (s *Stack) Pop() (interface{}, error) {
	if s.len == 0 {
		return nil, ErrStackIsEmpty
	}
	s.len--
	v := s.stack[s.len]
	s.stack = s.stack[:s.len]
	return v, nil
}

func (s *Stack) Push(v interface{}) error {
	if s.len == s.cap {
		return ErrStackIsFull
	}
	s.stack = append(s.stack, v)
	s.len++
	return nil
}

func (s *Stack) Clear() {
	if s.len > 0 {
		s.len = 0
		s.stack = make([]interface{}, 0, s.cap)
	}
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) Cap() int {
	return s.cap
}
