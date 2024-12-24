package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	elements []T
}

func (stk Stack[T]) GetLength() int {
	return len(stk.elements)
}

func (stk *Stack[T]) Push(element T) {
	stk.elements = append(stk.elements, element)
}

func (stk *Stack[T]) Pop() (T, error) {
	if stk.GetLength() == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	lastInd := len(stk.elements) - 1
	poppedElement := stk.elements[lastInd]
	stk.elements = stk.elements[:lastInd]
	return poppedElement, nil

}

func StackOps() {
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	result, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Popped element:", result)
	}
	fmt.Println(stack.GetLength())

	// Strings stack
	stack2 := Stack[string]{}
	stack2.Push("a")
	stack2.Push("b")
	stack2.Push("c")
	stack2.Push("d")

	result2, err := stack2.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Popped element:", result2)
	}
	fmt.Println(stack2.GetLength())
}
