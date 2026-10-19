package main

import "fmt"

type Node struct {
	chislo string
	next   *Node
	prev   *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func InitQueue(q *Queue) {
	q.head = nil
	q.tail = nil
}
func Insert(q *Queue, chislo string) {
	new_node := &Node{chislo, nil, q.tail}
	if q.tail == nil {
		q.head = new_node
	} else {
		q.tail.next = new_node
	}
	q.tail = new_node
}
func Delete(q *Queue) {
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	} else {
		q.head.prev = nil
	}
}
func front(q *Queue) string {
	return q.head.chislo
}
func back(q *Queue) string {
	return q.tail.chislo
}
func Print(q *Queue) {
	cur_node := q.head
	for cur_node != nil {
		fmt.Println(cur_node.chislo)
		cur_node = cur_node.next
	}
}
func PrintBack(q *Queue) {
	cur_node := q.tail
	for cur_node != nil {
		fmt.Println(cur_node.chislo)
		cur_node = cur_node.prev
	}
}
func ClearQueue(q *Queue) {
	q.head = nil
	q.tail = nil
}
func isEmptyQueue(q *Queue) bool {
	return q.head == nil
}
func PrintAddresses(q *Queue) {
	cur_node := q.head
	for cur_node != nil {
		fmt.Printf("%s | адрес: %p | next: %p | prev: %p\n", cur_node.chislo, cur_node, cur_node.next, cur_node.prev)
		cur_node = cur_node.next
	}
	fmt.Println()
}

func main() {
	var q Queue
	InitQueue(&q)
	Insert(&q, "1")
	PrintAddresses(&q)
	Insert(&q, "2")
	PrintAddresses(&q)
	Insert(&q, "3")
	PrintAddresses(&q)
	fmt.Println("От головы к хвосту ")
	Print(&q)
	fmt.Println("\nОт хвоста к голове ")
	PrintBack(&q)
	fmt.Println("\nПервый", front(&q))
	fmt.Println("Последний", back(&q))
	Delete(&q)
	PrintAddresses(&q)
	fmt.Println("\nПосле удаления ")
	Print(&q)
	ClearQueue(&q)
	fmt.Println("\nОчередь пуста ", isEmptyQueue(&q))
}
