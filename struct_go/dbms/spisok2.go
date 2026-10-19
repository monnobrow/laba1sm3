package main

import "fmt"

type LNode struct {
	chislo string
	next   *LNode
	prev   *LNode
}

type LList struct {
	head *LNode
	tail *LNode
}

// поиск target делается внутри операции: один проход, потом перевязка за O(1)
// результат ("-> x") или ошибку операция печатает сама

func LSearch(flist *LNode, chislo string) *LNode {
	for flist != nil && flist.chislo != chislo {
		flist = flist.next
	}
	return flist
}
func LLinkBefore(list *LList, flist *LNode, chislo string) {
	new_node := &LNode{chislo, flist, flist.prev}
	if flist.prev != nil {
		flist.prev.next = new_node
	} else {
		list.head = new_node
	}
	flist.prev = new_node
}
func LLinkAfter(list *LList, flist *LNode, chislo string) {
	new_node := &LNode{chislo, flist.next, flist}
	if flist.next != nil {
		flist.next.prev = new_node
	} else {
		list.tail = new_node
	}
	flist.next = new_node
}
func LDeleteNode(list *LList, flist *LNode) {
	fmt.Println("->", flist.chislo)
	if flist.prev != nil {
		flist.prev.next = flist.next
	} else {
		list.head = flist.next
	}
	if flist.next != nil {
		flist.next.prev = flist.prev
	} else {
		list.tail = flist.prev
	}
}

func LAddFirst(list *LList, chislo string) {
	new_node := &LNode{chislo, list.head, nil}
	if list.head != nil {
		list.head.prev = new_node
	} else {
		list.tail = new_node
	}
	list.head = new_node
}
func LAddEnd(list *LList, chislo string) {
	new_node := &LNode{chislo, nil, list.tail}
	if list.tail != nil {
		list.tail.next = new_node
	} else {
		list.head = new_node
	}
	list.tail = new_node
}
func LAddAfter(list *LList, target, chislo string) {
	cur_node := LSearch(list.head, target)
	if cur_node == nil {
		fmt.Println("Ошибка:", target, "не найден")
		return
	}
	LLinkAfter(list, cur_node, chislo)
	fmt.Println("->", chislo)
}
func LAddBefore(list *LList, target, chislo string) {
	cur_node := LSearch(list.head, target)
	if cur_node == nil {
		fmt.Println("Ошибка:", target, "не найден")
		return
	}
	LLinkBefore(list, cur_node, chislo)
	fmt.Println("->", chislo)
}

// index == длина - вставка в конец
func LAddIndex(list *LList, chislo string, index int) {
	cur_node := list.head
	i := 0
	for i < index && cur_node != nil {
		cur_node = cur_node.next
		i++
	}
	if index < 0 || i < index {
		fmt.Println("Ошибка: нет такого индекса")
		return
	}
	if cur_node == nil {
		LAddEnd(list, chislo)
	} else {
		LLinkBefore(list, cur_node, chislo)
	}
	fmt.Println("->", chislo)
}
func LDeleteFirst(list *LList) {
	if list.head == nil {
		fmt.Println("Ошибка: список пуст")
		return
	}
	LDeleteNode(list, list.head)
}
func LDeleteEnd(list *LList) {
	if list.tail == nil {
		fmt.Println("Ошибка: список пуст")
		return
	}
	LDeleteNode(list, list.tail)
}
func LDeleteAfter(list *LList, target string) {
	cur_node := LSearch(list.head, target)
	if cur_node == nil || cur_node.next == nil {
		fmt.Println("Ошибка: после", target, "нечего удалять")
		return
	}
	LDeleteNode(list, cur_node.next)
}
func LDeleteBefore(list *LList, target string) {
	cur_node := LSearch(list.head, target)
	if cur_node == nil || cur_node.prev == nil {
		fmt.Println("Ошибка: перед", target, "нечего удалять")
		return
	}
	LDeleteNode(list, cur_node.prev)
}
func LDeleteIndex(list *LList, index int) {
	cur_node := list.head
	for i := 0; i < index && cur_node != nil; i++ {
		cur_node = cur_node.next
	}
	if index < 0 || cur_node == nil {
		fmt.Println("Ошибка: нет такого индекса")
		return
	}
	LDeleteNode(list, cur_node)
}
func LDeleteValue(list *LList, chislo string) {
	cur_node := LSearch(list.head, chislo)
	if cur_node == nil {
		fmt.Println("Ошибка:", chislo, "не найден")
		return
	}
	LDeleteNode(list, cur_node)
}
func LGet(list *LList, index int) {
	cur_node := list.head
	for i := 0; i < index && cur_node != nil; i++ {
		cur_node = cur_node.next
	}
	if index < 0 || cur_node == nil {
		fmt.Println("Ошибка: нет такого индекса")
		return
	}
	fmt.Println("->", cur_node.chislo)
}
func LFind(list *LList, chislo string) {
	if LSearch(list.head, chislo) != nil {
		fmt.Println("-> TRUE")
	} else {
		fmt.Println("-> FALSE")
	}
}
func LPrint(list *LList) {
	for cur := list.head; cur != nil; cur = cur.next {
		fmt.Println(cur.chislo)
	}
}
func LPrintBack(list *LList) {
	for cur := list.tail; cur != nil; cur = cur.prev {
		fmt.Println(cur.chislo)
	}
}
