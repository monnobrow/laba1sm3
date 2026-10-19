package main

import "fmt"

type Q1Node struct {
	chislo string
	next   *Q1Node
}
type Queue1 struct {
	head *Q1Node
	tail *Q1Node
}

func Q1Push(q *Queue1, chislo string) {
	new_node := &Q1Node{chislo, nil}
	if q.head == nil {
		q.head = new_node
		q.tail = new_node
		return
	}
	q.tail.next = new_node
	q.tail = new_node
}

// снимает голову и сразу отдаёт её значение
func Q1Pop(q *Queue1) string {
	if q.head == nil {
		fmt.Println("Ошибка: очередь пуста")
		return ""
	}
	chislo := q.head.chislo
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return chislo
}
func Q1Print(q *Queue1) {
	for cur := q.head; cur != nil; cur = cur.next {
		fmt.Println(cur.chislo)
	}
}
