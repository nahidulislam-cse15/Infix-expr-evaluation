package main

import (
	"fmt"
)

var val1, val2 int
var op byte

type Stacki struct {
	items []int
}

func (stack *Stacki) Push(item int) {
	stack.items = append(stack.items, item)
}
func (stack *Stacki) Pop() int {
	if len(stack.items) == 0 {
		fmt.Println("stack is empty")
	}
	lastIndex := len(stack.items) - 1
	removeItem := stack.items[lastIndex]
	stack.items = stack.items[:lastIndex]

	return removeItem
}


func (stack *Stacki) Top() int {
	if len(stack.items) == 0 {
		fmt.Println("stack is empty")
		return 0
	}
	return stack.items[len(stack.items)-1]

}

func (stack *Stacki) isEmpty() bool {
	return len(stack.items) == 0
}
func (stack *Stacki) size() int {
	return len(stack.items)
}

type Stack struct {
	items []byte
}

func (stack *Stack) Push(item byte) {
	stack.items = append(stack.items, item)
}
func (stack *Stack) Pop() byte {
	if len(stack.items) == 0 {
		fmt.Println("stack is empty")
	}
	lastIndex := len(stack.items) - 1
	removeItem := stack.items[lastIndex]
	stack.items = stack.items[:lastIndex]

	return removeItem
}


func (stack *Stack) Top() byte {
	if len(stack.items) == 0 {
		fmt.Println("stack is empty")
	}
	return stack.items[len(stack.items)-1]
}

func (stack *Stack) isEmpty() bool {
	return len(stack.items) == 0
}
func (stack *Stack) size() int {
	return len(stack.items)
}

// Function to find precedence of operators.
func precedence(op byte) int {
	if op == '+' || op == '-' {
		return 1
	}
	if op == '*' || op == '/' {
		return 2
	}
	return 0
}

// Function to perform arithmetic operations.
func applyOp(a int, b int, op byte) int {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	default:
		return -1
	}
	
}
//calculation and push in stack
func calculate(operand *Stacki,ops *Stack) {
	val2 = operand.Top()
	operand.Pop()

	val1 = operand.Top()
	operand.Pop()

	op = ops.Top()
	ops.Pop()

	operand.Push((applyOp(val1, val2, op)))

}
func main() {

	var ops Stack //char type stack
	var operand Stacki //int type stack 
	expr := "((10*2+2)*(4/2*2))"
	for i := 0; i < len(expr); i++ {
		// whitespace skip 
		if expr[i] == ' ' {
			continue

			//  opening brace, Push it to 'ops'
		} else if expr[i] == '(' {
			ops.Push(expr[i])

			// number Push it to stack for numbers.
			
		} else if expr[i] >= '0' && expr[i] <= '9' {
			val := 0

			//  more  digits 
			for i < len(expr) && expr[i] >= '0' && expr[i] <= '9' {
				val = (val * 10) + int(expr[i]-'0')
				i++
			}
			
			operand.Push(val)
			i--

		} else if expr[i] == ')' {

			for !ops.isEmpty() && ops.Top() != '(' {

				calculate(&operand,&ops)
			}

			// Pop opening brace.
			if !ops.isEmpty() {
				ops.Pop()
			}
		} else {
			// While Top of 'ops' has same or greater// precedence to current op
			//Apply operator on Top of 'ops' to Top two elements in operand stack.
			for !ops.isEmpty() && precedence(ops.Top()) >= precedence(expr[i]) {
				calculate(&operand,&ops)
			}

			// Push current op to 'ops'.
			ops.Push(expr[i])
		}
	}

	for !ops.isEmpty() {
		calculate(&operand,&ops)
	}

	// Top of 'operand' contains result, return it.
	fmt.Println(operand.Top())
}
