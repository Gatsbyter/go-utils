package stack

type IStack interface {
	Push(data interface{})
	Pop() interface{}
	Peak() interface{}
	Size() int
	ToSlice() []interface{}
}

type Stack struct {
	list []interface{}
}

func NewStack() IStack {
	return &Stack{list: []interface{}{}}
}

func (s *Stack) Push(data interface{}) {
	s.list = append(s.list, data)
}

func (s *Stack) Pop() interface{} {
	length := s.Size()
	if length == 0 {
		return nil
	}

	res := s.list[length-1]
	s.list = s.list[:length-1]
	return res
}

func (s *Stack) Peak() interface{} {
	length := s.Size()
	if length == 0 {
		return nil
	}

	return s.list[length-1]
}

func (s *Stack) Size() int {
	return len(s.list)
}

func (s *Stack) ToSlice() []interface{} {
	return s.list
}
