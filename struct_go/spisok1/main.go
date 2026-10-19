package main

import "fmt"

type Node struct {
	chislo string
	next   *Node
}

func AddAfter(flist *Node, chislo string) {
	new_node := &Node{chislo, flist.next}
	flist.next = new_node
}
func AddFirst(flist **Node, chislo string) {
	new_node := &Node{chislo, *flist}
	*flist = new_node
}
func AddEnd(flist **Node, chislo string) {
	new_node := &Node{chislo, nil}
	cur_node := *flist
	if *flist == nil {
		*flist = new_node
		return
	}
	for cur_node.next != nil {
		cur_node = cur_node.next
	}
	cur_node.next = new_node
}
func AddBefore(flist **Node, target *Node, chislo string) {
	cur_node := *flist
	new_node := &Node{chislo, target}
	for cur_node.next != target {
		cur_node = cur_node.next
	}
	cur_node.next = new_node
}
func AddIndex(flist **Node, chislo string, index int) {
	cur_node := *flist
	if index == 0 {
		AddFirst(flist, chislo)
		return
	}
	for i := 0; i < index-1; i++ {
		cur_node = cur_node.next
	}
	new_node := &Node{chislo, cur_node.next}
	cur_node.next = new_node
}

func DeleteAfter(flist *Node) {
	delete_node := flist.next
	flist.next = delete_node.next
}
func DeleteFirst(flist **Node) {
	*flist = (*flist).next
}
func DeleteEnd(flist **Node) {
	if (*flist).next == nil {
		DeleteFirst(flist)
		return
	}
	cur_node := *flist
	for cur_node.next.next != nil {
		cur_node = cur_node.next
	}
	DeleteAfter(cur_node)
}
func DeleteBefore(flist **Node, target *Node) {
	if (*flist).next == target {
		DeleteFirst(flist)
		return
	}
	cur_node := *flist
	for cur_node.next.next != target {
		cur_node = cur_node.next
	}
	DeleteAfter(cur_node)
}
func DeleteIndex(flist **Node, index int) {
	cur_node := *flist
	for i := 0; i < index-1; i++ {
		cur_node = cur_node.next
	}
	DeleteAfter(cur_node)
}
func DeleteValue(flist **Node, chislo string) {
	if (*flist).chislo == chislo {
		DeleteFirst(flist)
		return
	}
	cur_node := *flist
	for cur_node.next != nil && cur_node.next.chislo != chislo {
		cur_node = cur_node.next
	}
	if cur_node.next != nil {
		DeleteAfter(cur_node)
	}
}

func PrintBack(flist *Node) {
	if flist == nil {
		return
	}
	PrintBack(flist.next)
	fmt.Println(flist.chislo)
}
func Print(flist *Node) {
	for flist != nil {
		fmt.Println(flist.chislo)
		flist = flist.next
	}
}
func PrintAddresses(flist *Node) {
	for flist != nil {
		fmt.Printf("Узел: %p, chislo: %s, next: %p\n", flist, flist.chislo, flist.next)
		flist = flist.next
	}
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

func main() {
	forwardList := &Node{"1", nil}
	forwardList.next = &Node{"2", nil}
	forwardList.next.next = &Node{"3", nil}
	fmt.Println("Исходный список:")
	PrintAddresses(forwardList)
	AddFirst(&forwardList, "4")
	fmt.Println("\nПосле AddFirst(4):")
	PrintAddresses(forwardList)
	AddEnd(&forwardList, "6")
	fmt.Println("\nПосле AddEnd(6):")
	PrintAddresses(forwardList)
	AddAfter(forwardList.next, "5")
	fmt.Println("\nПосле AddAfter(после 1, 5):")
	PrintAddresses(forwardList)
	AddBefore(&forwardList, forwardList.next.next.next.next, "7")
	fmt.Println("\nПосле AddBefore(перед 3, 7):")
	PrintAddresses(forwardList)
	AddIndex(&forwardList, "8", 2)
	fmt.Println("\nПосле AddIndex(8, 2):")
	PrintAddresses(forwardList)
	fmt.Println("\nPrint:")
	Print(forwardList)
	fmt.Println("\nPrintBack:")
	PrintBack(forwardList)
	fmt.Println("\nFind:")
	Find(forwardList, "2")
	Find(forwardList, "100")
	DeleteFirst(&forwardList)
	fmt.Println("\nПосле DeleteFirst:")
	PrintAddresses(forwardList)
	DeleteEnd(&forwardList)
	fmt.Println("\nПосле DeleteEnd:")
	PrintAddresses(forwardList)
	DeleteAfter(forwardList)
	fmt.Println("\nПосле DeleteAfter(после 1):")
	PrintAddresses(forwardList)
	DeleteBefore(&forwardList, forwardList.next.next)
	fmt.Println("\nПосле DeleteBefore(перед 2):")
	PrintAddresses(forwardList)
	DeleteIndex(&forwardList, 1)
	fmt.Println("\nПосле DeleteIndex(1):")
	PrintAddresses(forwardList)
	DeleteValue(&forwardList, "7")
	fmt.Println("\nПосле DeleteValue(7):")
	PrintAddresses(forwardList)
}
