package main

import "fmt"

type FNode struct {
	chislo string
	next   *FNode
}

// поиск target делается внутри операции, за один проход по списку
// результат ("-> x") или ошибку операция печатает сама

func FAddFirst(flist **FNode, chislo string) {
	*flist = &FNode{chislo, *flist}
}
func FAddEnd(flist **FNode, chislo string) {
	new_node := &FNode{chislo, nil}
	if *flist == nil {
		*flist = new_node
		return
	}
	cur_node := *flist
	for cur_node.next != nil {
		cur_node = cur_node.next
	}
	cur_node.next = new_node
}

// нашли target и сразу вставили, второго прохода нет
func FAddAfter(flist *FNode, target, chislo string) {
	for flist != nil && flist.chislo != target {
		flist = flist.next
	}
	if flist == nil {
		fmt.Println("Ошибка:", target, "не найден")
		return
	}
	flist.next = &FNode{chislo, flist.next}
	fmt.Println("->", chislo)
}

// идём на шаг позади: cur_node.next - кандидат на target
func FAddBefore(flist **FNode, target, chislo string) {
	if *flist != nil && (*flist).chislo == target {
		FAddFirst(flist, chislo)
		fmt.Println("->", chislo)
		return
	}
	cur_node := *flist
	for cur_node != nil && cur_node.next != nil && cur_node.next.chislo != target {
		cur_node = cur_node.next
	}
	if cur_node == nil || cur_node.next == nil {
		fmt.Println("Ошибка:", target, "не найден")
		return
	}
	cur_node.next = &FNode{chislo, cur_node.next}
	fmt.Println("->", chislo)
}
func FAddIndex(flist **FNode, chislo string, index int) {
	if index == 0 {
		FAddFirst(flist, chislo)
		fmt.Println("->", chislo)
		return
	}
	cur_node := *flist
	for i := 0; i < index-1 && cur_node != nil; i++ {
		cur_node = cur_node.next
	}
	if index < 0 || cur_node == nil {
		fmt.Println("Ошибка: нет такого индекса")
		return
	}
	cur_node.next = &FNode{chislo, cur_node.next}
	fmt.Println("->", chislo)
}

// вырезает узел после prev
func FDeleteNext(prev *FNode) {
	fmt.Println("->", prev.next.chislo)
	prev.next = prev.next.next
}
func FDeleteFirst(flist **FNode) {
	if *flist == nil {
		fmt.Println("Ошибка: список пуст")
		return
	}
	fmt.Println("->", (*flist).chislo)
	*flist = (*flist).next
}
func FDeleteEnd(flist **FNode) {
	if *flist == nil || (*flist).next == nil {
		FDeleteFirst(flist)
		return
	}
	cur_node := *flist
	for cur_node.next.next != nil {
		cur_node = cur_node.next
	}
	FDeleteNext(cur_node)
}
func FDeleteAfter(flist *FNode, target string) {
	for flist != nil && flist.chislo != target {
		flist = flist.next
	}
	if flist == nil || flist.next == nil {
		fmt.Println("Ошибка: после", target, "нечего удалять")
		return
	}
	FDeleteNext(flist)
}

// удаляем узел перед target: идём на два шага позади
func FDeleteBefore(flist **FNode, target string) {
	if *flist == nil || (*flist).chislo == target {
		fmt.Println("Ошибка: перед", target, "нечего удалять")
		return
	}
	if (*flist).next != nil && (*flist).next.chislo == target {
		FDeleteFirst(flist)
		return
	}
	cur_node := *flist
	for cur_node.next != nil && cur_node.next.next != nil && cur_node.next.next.chislo != target {
		cur_node = cur_node.next
	}
	if cur_node.next == nil || cur_node.next.next == nil {
		fmt.Println("Ошибка: перед", target, "нечего удалять")
		return
	}
	FDeleteNext(cur_node)
}
func FDeleteIndex(flist **FNode, index int) {
	if index == 0 {
		FDeleteFirst(flist)
		return
	}
	cur_node := *flist
	for i := 0; i < index-1 && cur_node != nil; i++ {
		cur_node = cur_node.next
	}
	if index < 0 || cur_node == nil || cur_node.next == nil {
		fmt.Println("Ошибка: нет такого индекса")
		return
	}
	FDeleteNext(cur_node)
}
func FDeleteValue(flist **FNode, chislo string) {
	if *flist != nil && (*flist).chislo == chislo {
		FDeleteFirst(flist)
		return
	}
	cur_node := *flist
	for cur_node != nil && cur_node.next != nil && cur_node.next.chislo != chislo {
		cur_node = cur_node.next
	}
	if cur_node == nil || cur_node.next == nil {
		fmt.Println("Ошибка:", chislo, "не найден")
		return
	}
	FDeleteNext(cur_node)
}
func FGet(flist *FNode, index int) {
	for i := 0; i < index && flist != nil; i++ {
		flist = flist.next
	}
	if index < 0 || flist == nil {
		fmt.Println("Ошибка: нет такого индекса")
		return
	}
	fmt.Println("->", flist.chislo)
}
func FFind(flist *FNode, chislo string) {
	for ; flist != nil; flist = flist.next {
		if flist.chislo == chislo {
			fmt.Println("-> TRUE")
			return
		}
	}
	fmt.Println("-> FALSE")
}
func FPrintBack(flist *FNode) {
	if flist == nil {
		return
	}
	FPrintBack(flist.next)
	fmt.Println(flist.chislo)
}
func FPrint(flist *FNode) {
	for ; flist != nil; flist = flist.next {
		fmt.Println(flist.chislo)
	}
}
