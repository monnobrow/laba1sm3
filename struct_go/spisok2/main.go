package main

import "fmt"

type Node struct {
	chislo string
	next   *Node
	prev   *Node
}

type LList struct {
	head *Node
	tail *Node
}

func InitList(list *LList) {
	list.head = nil
	list.tail = nil
}

func Print(flist *Node) {
	for flist != nil {
		fmt.Println(flist.chislo)
		flist = flist.next
	}
}
func PrintBac(last *Node) {
	for last != nil {
		fmt.Println(last.chislo)
		last = last.prev
	}
}
func AddAfter(list *LList, flist *Node, chislo string) {
	new_node := &Node{chislo, flist.next, flist}
	if flist.next != nil {
		flist.next.prev = new_node
	} else {
		list.tail = new_node
	}
	flist.next = new_node
}
func AddFirst(list *LList, chislo string) {
	new_node := &Node{chislo, list.head, nil}
	if list.head != nil {
		list.head.prev = new_node
	} else {
		list.tail = new_node
	}
	list.head = new_node
}
func AddEnd(list *LList, chislo string) {
	new_node := &Node{chislo, nil, list.tail}
	if list.tail != nil {
		list.tail.next = new_node
	} else {
		list.head = new_node
	}
	list.tail = new_node
}
func AddBefore(list *LList, flist *Node, chislo string) {
	new_node := &Node{chislo, flist, flist.prev}
	if flist.prev != nil {
		flist.prev.next = new_node
	} else {
		list.head = new_node
	}
	flist.prev = new_node
}
func AddIndex(list *LList, chislo string, index int) {
	if index == 0 {
		AddFirst(list, chislo)
		return
	}
	cur_node := list.head
	for i := 0; i < index-1; i++ {
		cur_node = cur_node.next
	}
	AddAfter(list, cur_node, chislo)
}

func DeleteFirst(list *LList) {
	list.head = list.head.next
	if list.head != nil {
		list.head.prev = nil
	} else {
		list.tail = nil
	}
}
func DeleteEnd(list *LList) {
	list.tail = list.tail.prev
	if list.tail != nil {
		list.tail.next = nil
	} else {
		list.head = nil
	}
}
func DeleteAfter(list *LList, flist *Node) {
	delete_node := flist.next
	flist.next = delete_node.next
	if delete_node.next != nil {
		delete_node.next.prev = flist
	} else {
		list.tail = flist
	}
}
func DeleteBefore(list *LList, flist *Node) {
	delete_node := flist.prev
	flist.prev = delete_node.prev
	if delete_node.prev != nil {
		delete_node.prev.next = flist
	} else {
		list.head = flist
	}
}
func DeleteIndex(list *LList, index int) {
	cur_node := list.head
	for i := 0; i < index; i++ {
		cur_node = cur_node.next
	}
	DeleteNode(list, cur_node)
}
func DeleteNode(list *LList, flist *Node) {
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
func DeleteValue(list *LList, chislo string) {
	cur_node := list.head
	for cur_node != nil && cur_node.chislo != chislo {
		cur_node = cur_node.next
	}
	if cur_node == nil {
		return
	}
	DeleteNode(list, cur_node)
}
func Find(flist *Node, chislo string) {
	for flist != nil {
		if flist.chislo == chislo {
			fmt.Printf("%s найден, адрес: %p\n", chislo, flist)
			return
		}
		flist = flist.next
	}
	fmt.Println(chislo, "не найден")
}
func PrintAddresses(list *LList) {
	fmt.Printf("head: %p | tail: %p\n", list.head, list.tail)
	cur_node := list.head
	for cur_node != nil {
		fmt.Printf("%s | адрес: %p | next: %p | prev: %p\n", cur_node.chislo, cur_node, cur_node.next, cur_node.prev)
		cur_node = cur_node.next
	}
	fmt.Println()
}

func main() {
	var list LList
	InitList(&list)
	AddEnd(&list, "1")
	AddEnd(&list, "2")
	AddEnd(&list, "3")
	fmt.Println("Исходный список:")
	PrintAddresses(&list)
	AddFirst(&list, "4")
	fmt.Println("После AddFirst(4):")
	PrintAddresses(&list)
	AddEnd(&list, "6")
	fmt.Println("После AddEnd(6):")
	PrintAddresses(&list)
	AddAfter(&list, list.head.next, "5")
	fmt.Println("После AddAfter(после 1, 5):")
	PrintAddresses(&list)
	AddBefore(&list, list.head.next.next.next.next, "7")
	fmt.Println("После AddBefore(перед 3, 7):")
	PrintAddresses(&list)
	AddIndex(&list, "8", 2)
	fmt.Println("После AddIndex(8, 2):")
	PrintAddresses(&list)
	fmt.Println("Print:")
	Print(list.head)
	fmt.Println("\nPrintBac:")
	PrintBac(list.tail)
	fmt.Println("\nFind:")
	Find(list.head, "2")
	Find(list.head, "100")
	DeleteFirst(&list)
	fmt.Println("\nПосле DeleteFirst:")
	PrintAddresses(&list)
	DeleteEnd(&list)
	fmt.Println("После DeleteEnd:")
	PrintAddresses(&list)
	DeleteAfter(&list, list.head)
	fmt.Println("После DeleteAfter(после 1):")
	PrintAddresses(&list)
	DeleteBefore(&list, list.head.next.next)
	fmt.Println("После DeleteBefore(перед 2):")
	PrintAddresses(&list)
	DeleteNode(&list, list.head.next)
	fmt.Println("После DeleteNode(2):")
	PrintAddresses(&list)
	DeleteIndex(&list, 1)
	fmt.Println("После DeleteIndex(1):")
	PrintAddresses(&list)
	DeleteValue(&list, "3")
	fmt.Println("После DeleteValue(3):")
	PrintAddresses(&list)
}
