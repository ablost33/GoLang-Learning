package data_structures

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func PlayWithStack() {
	myStack := &stack.Stack{}
	myStack.Push(1)
	myStack.Push(2)
	fmt.Println(myStack.Peek())
	myStack.Pop()
	fmt.Println(myStack.Peek())
}
