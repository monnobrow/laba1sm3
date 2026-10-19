// ./dbms --file data.json --query 'КОМАНДА имя [аргументы]'
// линейные структуры на Go, деревья по заданию только в C++
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filename, query := "", ""
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "--file" && i+1 < len(os.Args) {
			i++
			filename = os.Args[i]
		} else if os.Args[i] == "--query" && i+1 < len(os.Args) {
			i++
			query = os.Args[i]
		}
	}
	if filename == "" || query == "" {
		fmt.Println("Запуск: ./dbms --file data.json --query 'КОМАНДА имя [аргументы]'")
		os.Exit(1)
	}
	q := Split(query)
	if len(q) == 0 {
		Err("пустой запрос")
		os.Exit(1)
	}
	q[0] = strings.ToUpper(q[0])
	cmd := q[0]
	if cmd == "HELP" {
		RunHelp()
		return
	}
	typ := TypeOfCommand(cmd)
	if cmd != "PRINT" && typ == "" {
		if strings.HasPrefix(cmd, "T") || cmd == "ISMEMBER" {
			Err("деревья только в C++ версии")
		} else {
			Err("неизвестная команда " + cmd + ", см. HELP")
		}
		os.Exit(1)
	}
	if len(q) < 2 {
		Err("не указано имя структуры")
		os.Exit(1)
	}
	name := q[1]

	data, err := LoadJson(filename)
	if err != nil {
		Err(filename + " - не JSON нужного вида, файл не изменён (" + err.Error() + ")")
		os.Exit(1)
	}

	if cmd == "PRINT" {
		RunPrint(data, name)
		return
	}
	values, ok := GetValues(data, typ, name)
	if !ok && !IsCreateCommand(cmd) {
		Err("структура " + name + " не найдена")
		os.Exit(1)
	}
	after := RunCommand(q, values)
	if !SameValues(values, after) { // перезаписываем только при изменении
		SetValues(data, typ, name, after)
		if err := SaveJson(data, filename); err != nil {
			Err(err.Error())
			os.Exit(1)
		}
	}
}
