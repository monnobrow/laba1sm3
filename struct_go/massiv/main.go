package main

import "fmt"

type Array struct {
	data  []string
	size  int
	mesto int
}

func InitArray(arr *Array, mesto int) {
	arr.data = make([]string, mesto)
	arr.size = 0
	arr.mesto = mesto
}
func Resize(arr *Array) {
	new_data := make([]string, arr.mesto*2)
	for i := 0; i < arr.size; i++ {
		new_data[i] = arr.data[i]
	}
	arr.data = new_data
	arr.mesto = arr.mesto * 2
}
func AddEnd(arr *Array, chislo string) {
	if arr.size == arr.mesto {
		Resize(arr)
	}
	arr.data[arr.size] = chislo
	arr.size++
}
func AddIndex(arr *Array, chislo string, index int) {
	if index < 0 || index > arr.size {
		fmt.Println("no")
		return
	}
	if arr.size == arr.mesto {
		Resize(arr)
	}
	for i := arr.size; i > index; i-- {
		arr.data[i] = arr.data[i-1]
	}
	arr.data[index] = chislo
	arr.size++
}
func Get(arr *Array, index int) string {
	return arr.data[index]
}
func DeleteIndex(arr *Array, index int) {
	if index < 0 || index >= arr.size {
		fmt.Println("no")
		return
	}
	for i := index; i < arr.size-1; i++ {
		arr.data[i] = arr.data[i+1]
	}
	arr.size--
}
func Replace(arr *Array, chislo string, index int) {
	if index < 0 || index >= arr.size {
		fmt.Println("no")
		return
	}
	arr.data[index] = chislo
}
func Length(arr *Array) int {
	return arr.size
}
func Print(arr *Array) {
	for i := 0; i < arr.size; i++ {
		fmt.Print(arr.data[i], " ")
	}
	fmt.Println()
}
func PrintAddresses(arr *Array) {
	for i := 0; i < arr.size; i++ {
		fmt.Printf("[%d] %s | адрес: %p\n", i, arr.data[i], &arr.data[i])
	}
	fmt.Printf("size = %d mesto = %d\n\n", arr.size, arr.mesto)
}

func main() {
	var arr Array
	InitArray(&arr, 2)
	fmt.Println("Создали массив:")
	PrintAddresses(&arr)
	AddEnd(&arr, "1")
	AddEnd(&arr, "2")
	fmt.Println("После AddEnd(1) и AddEnd(2):")
	PrintAddresses(&arr)
	AddEnd(&arr, "3")
	fmt.Println("После AddEnd(3), массив расширился:")
	PrintAddresses(&arr)
	AddIndex(&arr, "4", 1)
	fmt.Println("После AddIndex(4, 1):")
	PrintAddresses(&arr)
	AddIndex(&arr, "5", 0)
	fmt.Println("После AddIndex(5, 0):")
	PrintAddresses(&arr)
	fmt.Println("Get(2) =", Get(&arr, 2))
	fmt.Println("Length =", Length(&arr))
	fmt.Println()
	Replace(&arr, "7", 2)
	fmt.Println("После Replace(7, 2):")
	PrintAddresses(&arr)
	DeleteIndex(&arr, 0)
	fmt.Println("После DeleteIndex(0):")
	PrintAddresses(&arr)
	DeleteIndex(&arr, 2)
	fmt.Println("После DeleteIndex(2):")
	PrintAddresses(&arr)
	fmt.Println("Print:")
	Print(&arr)
}
