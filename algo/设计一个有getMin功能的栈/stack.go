package stask

type Item interface{}

var _ IStack = (*Stack)(nil)

type IStack interface {
	Push(interface{}) bool
	Pop() error
	GetMIn() interface{}
}
type Stack struct {
	Data    []Item
	MinData []Item
}

func NewStack() IStack {
	return &Stack{
		Data:    []Item{},
		MinData: []Item{},
	}
}

func (s *Stack) Push(v interface{}) bool {
	s.Data = append(s.Data, v)
	if len(s.MinData) == 0 {
		s.MinData = append(s.MinData, v)
		return true
	}
	if min, ok := s.MinData[len(s.MinData)-1].(int); ok {
		value := v.(int)
		if value < min {
			s.MinData = append(s.MinData, v)
			return true

		}
	}
	return true
}

func (s *Stack) Pop() error {
	if s.Data[len(s.Data)-1] == s.MinData[len(s.MinData)-1] {
		s.MinData = s.MinData[:len(s.MinData)-1]
	}
	s.Data = s.Data[:len(s.Data)-1]
	return nil
}

func (s *Stack) GetMIn() interface{} {
	if len(s.MinData) == 0 {
		return 0
	}
	return s.MinData[len(s.MinData)-1]
}
