package data_structures

type Stack struct {
	Store []string
}

func (s *Stack) Push(element string) {
	s.Store = append(s.Store, element)
}

func (s *Stack) Pop() string {
	if len(s.Store) == 0 {
		return ""
	}
	lastElementIndex := len(s.Store) - 1
	popped := s.Store[lastElementIndex]
	s.Store = s.Store[:lastElementIndex]
	return popped
}

type Queue struct {
	Store []string
}

func (q *Queue) Enqueue(element string) {
	q.Store = append(q.Store, element)
}

func (q *Queue) Dequeue() string {
	if len(q.Store) == 0 {
		return ""
	}
	element := q.Store[0]
	q.Store = q.Store[1:]
	return element
}
