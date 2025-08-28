package c

type Stack[T any] struct {
	Items []T
}

func (s *Stack[T]) Add(f T) {
	s.Items = append(s.Items, f)
}

func (s *Stack[T]) Pop() (item T, ok bool) {
	if len(s.Items) > 0 {
		item = s.Items[len(s.Items)-1]
		s.Items = s.Items[:len(s.Items)-1]
	} else {
		ok = false
	}
	return
}
