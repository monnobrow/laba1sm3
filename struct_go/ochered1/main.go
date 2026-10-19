package main

import "fmt"

type Node struct {
	chislo string
	next   *Node
}
type Queue struct {
	head *Node
	tail *Node
}

func InitQueue(q *Queue) {
	q.head = nil
	q.tail = nil
}

func PrintQueue(q *Queue) {
	cur_node := q.head
	for cur_node != nil {
		fmt.Println(cur_node.chislo)
		cur_node = cur_node.next
	}
}
func Insert(q *Queue, chislo string) {
	new_node := &Node{chislo, nil}
	if q.head == nil {
		q.head = new_node
		q.tail = new_node
		return
	}
	q.tail.next = new_node
	q.tail = new_node
}
func Delete(q *Queue) {
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
}
func front(q *Queue) string {
	return q.head.chislo
}

func back(q *Queue) string {
	return q.tail.chislo
}
func PrintAddresses(q *Queue) {
	cur_node := q.head
	for cur_node != nil {
		fmt.Printf("%s | адрес узла: %p | next: %p\n", cur_node.chislo, cur_node, cur_node.next)
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
	Insert(&q, "4")
	PrintAddresses(&q)
	PrintQueue(&q)
	fmt.Println("Первый:", front(&q))
	fmt.Println("Последний:", back(&q))
	Delete(&q)
	PrintQueue(&q)
}
