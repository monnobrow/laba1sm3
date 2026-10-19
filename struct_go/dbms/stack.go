package main

import "fmt"

type SNode struct {
	chislo string
	next   *SNode
}

type Stack struct {
	head *SNode
}

func SIsEmptyStack(stack *Stack) bool {
	return stack.head == nil
}
func SPush(stack *Stack, chislo string) {
	stack.head = &SNode{chislo, stack.head}
}

// снимает вершину и сразу отдаёт её значение
func SPop(stack *Stack) string {
	if SIsEmptyStack(stack) {
		fmt.Println("Ошибка: стек пуст")
		return ""
	}
	chislo := stack.head.chislo
	stack.head = stack.head.next
	return chislo
}
func SPrint(stack *Stack) {
	for cur := stack.head; cur != nil; cur = cur.next {
		fmt.Println(cur.chislo)
	}
}
