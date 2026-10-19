package main

import "fmt"

type Node struct {
	chislo string
	next   *Node
}

type Stack struct {
	head *Node
}

func InitStack(stack *Stack) {
	stack.head = nil
}
func isEmptyStack(stack *Stack) bool {
	return stack.head == nil
}
func push(stack *Stack, chislo string) {
	new_node := &Node{chislo, stack.head}
	stack.head = new_node
}
func pop(stack *Stack) {
	if isEmptyStack(stack) {
		fmt.Print("no")
	}
	stack.head = stack.head.next
}
func top(stack *Stack) string {
	return stack.head.chislo
}
func Print(flist *Node) {
	for flist != nil {
		fmt.Println(flist.chislo)
		flist = flist.next
	}
}
func PrintAddresses(stack *Stack) {
	cur_node := stack.head
	for cur_node != nil {
		fmt.Printf("%s | адрес узла: %p | next: %p\n", cur_node.chislo, cur_node, cur_node.next)
		cur_node = cur_node.next
	}
	fmt.Println()
}

func main() {
	var stack Stack
	InitStack(&stack)
	PrintAddresses(&stack)
	push(&stack, "1")
	fmt.Println("После push 1 ")
	PrintAddresses(&stack)
	push(&stack, "2")
	fmt.Println("После push 2 ")
	PrintAddresses(&stack)
	push(&stack, "3")
	fmt.Println("После push 3 ")
	PrintAddresses(&stack)
	fmt.Print("Вершина ", top(&stack), "\n\n")
	pop(&stack)
	fmt.Println("После первого pop ")
	PrintAddresses(&stack)
	pop(&stack)
	fmt.Println("После второго pop ")
	PrintAddresses(&stack)
	pop(&stack)
	fmt.Println("После третьего pop ")
	PrintAddresses(&stack)
}
