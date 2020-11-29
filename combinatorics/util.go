package combinatorics

// IntList ...
type IntList struct {
	first *intListNode
	last  *intListNode
	len   int
}

type intListNode struct {
	child *intListNode
	value int
}

// NewIntList ...
func NewIntList() *IntList {
	return &IntList{nil, nil, 0}
}

// Len ...
func (list *IntList) Len() int {
	return list.len
}

// Add ...
func (list *IntList) Add(elem int) {
	node := intListNode{nil, elem}
	if list.first == nil {
		list.first = &node
		list.last = &node
	} else {
		list.last.child = &node
		list.last = &node
	}
	list.len++
}

// Concat ...
func (list *IntList) Concat(other *IntList) {
	if list.first == nil {
		*list = *other
	} else if other.first == nil {
		// Do nothing
	} else {
		list.last.child = other.first
		list.last = other.last
		list.len += other.len
	}
}

// Each ...
func (list *IntList) Each(f func(elem int)) {
	cur := list.first
	for cur != nil {
		f(cur.value)
		cur = cur.child
	}
}

// ToA ...
func (list *IntList) ToA() []int {
	a := make([]int, list.Len())
	{
		index := 0
		list.Each(func(elem int) {
			a[index] = elem
			index++
		})
	}
	return a
}

// IntStack ...
type IntStack struct {
	last *intStackNode
}

type intStackNode struct {
	parent *intStackNode
	value  int
}

// NewIntStack ...
func NewIntStack() *IntStack {
	return &IntStack{nil}
}

// Push ...
func (s *IntStack) Push(elem int) {
	node := intStackNode{s.last, elem}
	s.last = &node
}

// Pop ...
func (s *IntStack) Pop() int {
	value := s.last.value
	s.last = s.last.parent
	return value
}

// Peek ...
func (s *IntStack) Peek() int {
	return s.last.value
}

// Empty ...
func (s *IntStack) Empty() bool {
	return s.last == nil
}

// FuncStack ...
type FuncStack struct {
	last *funcStackNode
}

type funcStackNode struct {
	parent *funcStackNode
	value  func()
}

// NewFuncStack ...
func NewFuncStack() *FuncStack {
	return &FuncStack{nil}
}

// Push ...
func (s *FuncStack) Push(elem func()) {
	node := funcStackNode{s.last, elem}
	s.last = &node
}

// Pop ...
func (s *FuncStack) Pop() func() {
	value := s.last.value
	s.last = s.last.parent
	return value
}

// Peek ...
func (s *FuncStack) Peek() func() {
	return s.last.value
}

// Empty ...
func (s *FuncStack) Empty() bool {
	return s.last == nil
}
