package main

import "fmt"

type Array struct {
	data  []string
	size  int
	mesto int
}

func MInitArray(arr *Array, mesto int) {
	if mesto < 1 {
		mesto = 1
	}
	arr.data = make([]string, mesto)
	arr.size = 0
	arr.mesto = mesto
}
func MResize(arr *Array) {
	new_data := make([]string, arr.mesto*2)
	for i := 0; i < arr.size; i++ {
		new_data[i] = arr.data[i]
	}
	arr.data = new_data
	arr.mesto = arr.mesto * 2
}
func MAddEnd(arr *Array, chislo string) {
	if arr.size == arr.mesto {
		MResize(arr)
	}
	arr.data[arr.size] = chislo
	arr.size++
}

// проверка индекса внутри структуры, результат печатается тут же
func MAddIndex(arr *Array, chislo string, index int) {
	if index < 0 || index > arr.size {
		fmt.Println("Ошибка: нет такого индекса")
		return
	}
	if arr.size == arr.mesto {
		MResize(arr)
	}
	for i := arr.size; i > index; i-- {
		arr.data[i] = arr.data[i-1]
	}
	arr.data[index] = chislo
	arr.size++
	fmt.Println("->", chislo)
}
func MGet(arr *Array, index int) string {
	return arr.data[index] // без проверки
}
func MGetIndex(arr *Array, index int) {
	if index < 0 || index >= arr.size {
		fmt.Println("Ошибка: нет такого индекса")
		return
	}
	fmt.Println("->", arr.data[index])
}
func MDeleteIndex(arr *Array, index int) {
	if index < 0 || index >= arr.size {
		fmt.Println("Ошибка: нет такого индекса")
		return
	}
	fmt.Println("->", arr.data[index])
	for i := index; i < arr.size-1; i++ {
		arr.data[i] = arr.data[i+1]
	}
	arr.size--
}
func MReplace(arr *Array, chislo string, index int) {
	if index < 0 || index >= arr.size {
		fmt.Println("Ошибка: нет такого индекса")
		return
	}
	arr.data[index] = chislo
	fmt.Println("->", chislo)
}
func MLength(arr *Array) int {
	return arr.size
}
func MPrint(arr *Array) {
	for i := 0; i < arr.size; i++ {
		fmt.Print(arr.data[i], " ")
	}
	fmt.Println()
}
