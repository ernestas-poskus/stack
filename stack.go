package stack

import (
	"sync"
)

// Stack data structure implementation
type Stack struct {
	head *node
	size int
	lock *sync.Mutex
}

type node struct {
	value interface{}
	next  *node
}

// New create strack object
func New() *Stack {
	s := &Stack{}
	s.lock = &sync.Mutex{}
	return s
}

// Len returns stack size
func (s *Stack) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.size
}

// Push adds new node to stack
func (s *Stack) Push(data interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	n := &node{value: data}

	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n
	}
	s.size++
}

// Pop node value from stack
func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	var n *node
	if s.head != nil {
		n = s.head
		s.head = n.next
		s.size--
	}

	if n == nil {
		return nil
	}

	return n.value
}
