package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Say(s string) {
	fmt.Println("->", s)
}
func Err(s string) {
	fmt.Println("Ошибка:", s)
}

// слова через пробел, "в кавычках" - одно слово с пробелами; пустые слова пропускаем
func Split(s string) []string {
	var out []string
	r := []rune(s)
	i, n := 0, len(r)
	for i < n {
		for i < n && unicode.IsSpace(r[i]) {
			i++
		}
		if i >= n {
			break
		}
		var word []rune
		if r[i] == '"' {
			i++
			for i < n && r[i] != '"' {
				word = append(word, r[i])
				i++
			}
			i++
		} else {
			for i < n && !unicode.IsSpace(r[i]) {
				word = append(word, r[i])
				i++
			}
		}
		if len(word) > 0 { // пустые слова пропускаем
			out = append(out, string(word))
		}
	}
	return out
}

func ToInt(s string) (int, bool) {
	x, err := strconv.Atoi(s)
	return x, err == nil
}

// тип структуры в файле по команде; "" - такой команды нет
func TypeOfCommand(cmd string) string {
	switch cmd {
	case "MPUSH", "MINSERT", "MGET", "MSET", "MDEL", "MLEN":
		return "massiv"
	case "FPUSH", "FDEL", "FGET", "FFIND", "FPRINTBACK":
		return "spisok1"
	case "LPUSH", "LDEL", "LGET", "LFIND", "LPRINTBACK":
		return "spisok2"
	case "QPUSH", "QPOP":
		return "ochered"
	case "SPUSH", "SPOP":
		return "stack"
	}
	return ""
}

func IsCreateCommand(cmd string) bool {
	switch cmd {
	case "MPUSH", "MINSERT", "FPUSH", "LPUSH", "QPUSH", "SPUSH":
		return true
	}
	return false
}

// ---------- перенос значений файл <-> структура ----------
func LoadM(values []string) *Array {
	var arr Array
	MInitArray(&arr, len(values))
	for _, v := range values {
		MAddEnd(&arr, v)
	}
	return &arr
}
func SaveM(arr *Array) []string {
	values := []string{}
	for i := 0; i < MLength(arr); i++ {
		values = append(values, MGet(arr, i))
	}
	return values
}
func LoadF(values []string) *FNode {
	var head *FNode
	for i := len(values) - 1; i >= 0; i-- { // с конца в голову - O(n)
		FAddFirst(&head, values[i])
	}
	return head
}
func SaveF(head *FNode) []string {
	values := []string{}
	for ; head != nil; head = head.next {
		values = append(values, head.chislo)
	}
	return values
}
func LoadL(values []string) *LList {
	list := &LList{}
	for _, v := range values {
		LAddEnd(list, v)
	}
	return list
}
func SaveL(list *LList) []string {
	values := []string{}
	for cur := list.head; cur != nil; cur = cur.next {
		values = append(values, cur.chislo)
	}
	return values
}
func LoadQ(values []string) *Queue1 {
	q := &Queue1{}
	for _, v := range values {
		Q1Push(q, v)
	}
	return q
}
func SaveQ(q *Queue1) []string {
	values := []string{}
	for cur := q.head; cur != nil; cur = cur.next {
		values = append(values, cur.chislo)
	}
	return values
}

// стек в файле от дна к вершине
func LoadS(values []string) *Stack {
	stack := &Stack{}
	for _, v := range values {
		SPush(stack, v)
	}
	return stack
}
func SaveS(stack *Stack) []string {
	values := []string{}
	for cur := stack.head; cur != nil; cur = cur.next {
		values = append([]string{cur.chislo}, values...)
	}
	return values
}

// ---------- M: массив ----------
func RunM(q, values []string) []string {
	arr := LoadM(values)
	cmd, n := q[0], len(q)
	switch {
	case cmd == "MPUSH" && n == 3:
		MAddEnd(arr, q[2])
		Say(q[2])
	case (cmd == "MINSERT" || cmd == "MSET") && n == 4:
		index, isInt := ToInt(q[2])
		if !isInt {
			Err("индекс должен быть числом")
		} else if cmd == "MINSERT" {
			MAddIndex(arr, q[3], index)
		} else {
			MReplace(arr, q[3], index)
		}
	case (cmd == "MGET" || cmd == "MDEL") && n == 3:
		index, isInt := ToInt(q[2])
		if !isInt {
			Err("индекс должен быть числом")
		} else if cmd == "MGET" {
			MGetIndex(arr, index)
		} else {
			MDeleteIndex(arr, index)
		}
	case cmd == "MLEN" && n == 2:
		Say(strconv.Itoa(MLength(arr)))
	default:
		Err("неверные аргументы, см. HELP")
	}
	return SaveM(arr)
}

// ---------- F: односвязный список ----------
func RunF(q, values []string) []string {
	head := LoadF(values)
	cmd, n := q[0], len(q)
	mode := ""
	if n > 2 {
		mode = strings.ToUpper(q[2])
	}
	switch {
	case cmd == "FPUSH" && (n == 3 || (n == 4 && mode == "TAIL")):
		FAddEnd(&head, q[n-1])
		Say(q[n-1])
	case cmd == "FPUSH" && n == 4 && mode == "HEAD":
		FAddFirst(&head, q[3])
		Say(q[3])
	case cmd == "FPUSH" && n == 5 && mode == "AFTER":
		FAddAfter(head, q[3], q[4])
	case cmd == "FPUSH" && n == 5 && mode == "BEFORE":
		FAddBefore(&head, q[3], q[4])
	case cmd == "FPUSH" && n == 5 && mode == "INDEX":
		if index, isInt := ToInt(q[3]); isInt {
			FAddIndex(&head, q[4], index)
		} else {
			Err("индекс должен быть числом")
		}
	case cmd == "FDEL" && n == 3 && mode == "HEAD":
		FDeleteFirst(&head)
	case cmd == "FDEL" && n == 3 && mode == "TAIL":
		FDeleteEnd(&head)
	case cmd == "FDEL" && n == 4 && mode == "AFTER":
		FDeleteAfter(head, q[3])
	case cmd == "FDEL" && n == 4 && mode == "BEFORE":
		FDeleteBefore(&head, q[3])
	case cmd == "FDEL" && n == 4 && mode == "VALUE":
		FDeleteValue(&head, q[3])
	case cmd == "FDEL" && n == 4 && mode == "INDEX":
		if index, isInt := ToInt(q[3]); isInt {
			FDeleteIndex(&head, index)
		} else {
			Err("индекс должен быть числом")
		}
	case cmd == "FGET" && n == 3:
		if index, isInt := ToInt(q[2]); isInt {
			FGet(head, index)
		} else {
			Err("индекс должен быть числом")
		}
	case cmd == "FFIND" && n == 3:
		FFind(head, q[2])
	case cmd == "FPRINTBACK" && n == 2:
		FPrintBack(head)
	default:
		Err("неверные аргументы, см. HELP")
	}
	return SaveF(head)
}

// ---------- L: двусвязный список ----------
func RunL(q, values []string) []string {
	list := LoadL(values)
	cmd, n := q[0], len(q)
	mode := ""
	if n > 2 {
		mode = strings.ToUpper(q[2])
	}
	switch {
	case cmd == "LPUSH" && (n == 3 || (n == 4 && mode == "TAIL")):
		LAddEnd(list, q[n-1])
		Say(q[n-1])
	case cmd == "LPUSH" && n == 4 && mode == "HEAD":
		LAddFirst(list, q[3])
		Say(q[3])
	case cmd == "LPUSH" && n == 5 && mode == "AFTER":
		LAddAfter(list, q[3], q[4])
	case cmd == "LPUSH" && n == 5 && mode == "BEFORE":
		LAddBefore(list, q[3], q[4])
	case cmd == "LPUSH" && n == 5 && mode == "INDEX":
		if index, isInt := ToInt(q[3]); isInt {
			LAddIndex(list, q[4], index)
		} else {
			Err("индекс должен быть числом")
		}
	case cmd == "LDEL" && n == 3 && mode == "HEAD":
		LDeleteFirst(list)
	case cmd == "LDEL" && n == 3 && mode == "TAIL":
		LDeleteEnd(list)
	case cmd == "LDEL" && n == 4 && mode == "AFTER":
		LDeleteAfter(list, q[3])
	case cmd == "LDEL" && n == 4 && mode == "BEFORE":
		LDeleteBefore(list, q[3])
	case cmd == "LDEL" && n == 4 && mode == "VALUE":
		LDeleteValue(list, q[3])
	case cmd == "LDEL" && n == 4 && mode == "INDEX":
		if index, isInt := ToInt(q[3]); isInt {
			LDeleteIndex(list, index)
		} else {
			Err("индекс должен быть числом")
		}
	case cmd == "LGET" && n == 3:
		if index, isInt := ToInt(q[2]); isInt {
			LGet(list, index)
		} else {
			Err("индекс должен быть числом")
		}
	case cmd == "LFIND" && n == 3:
		LFind(list, q[2])
	case cmd == "LPRINTBACK" && n == 2:
		LPrintBack(list)
	default:
		Err("неверные аргументы, см. HELP")
	}
	return SaveL(list)
}

// ---------- Q: очередь ----------
func RunQ(q, values []string) []string {
	queue := LoadQ(values)
	cmd, n := q[0], len(q)
	switch {
	case cmd == "QPUSH" && n == 3:
		Q1Push(queue, q[2])
		Say(q[2])
	case cmd == "QPOP" && n == 2:
		if chislo := Q1Pop(queue); chislo != "" { // pop сразу отдаёт значение
			Say(chislo)
		}
	default:
		Err("неверные аргументы, см. HELP")
	}
	return SaveQ(queue)
}

// ---------- S: стек ----------
func RunS(q, values []string) []string {
	stack := LoadS(values)
	cmd, n := q[0], len(q)
	switch {
	case cmd == "SPUSH" && n == 3:
		SPush(stack, q[2])
		Say(q[2])
	case cmd == "SPOP" && n == 2:
		if chislo := SPop(stack); chislo != "" { // pop сразу отдаёт значение
			Say(chislo)
		}
	default:
		Err("неверные аргументы, см. HELP")
	}
	return SaveS(stack)
}

func RunCommand(q, values []string) []string {
	switch q[0][0] {
	case 'M':
		return RunM(q, values)
	case 'F':
		return RunF(q, values)
	case 'L':
		return RunL(q, values)
	case 'Q':
		return RunQ(q, values)
	case 'S':
		return RunS(q, values)
	}
	return values
}

func SameValues(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// PRINT выводит все структуры с этим именем
func RunPrint(data Data, name string) {
	found := false
	for _, typ := range []string{"massiv", "spisok1", "spisok2", "ochered", "stack", "tree"} {
		values, ok := GetValues(data, typ, name)
		if !ok {
			continue
		}
		found = true
		switch typ {
		case "massiv":
			fmt.Print("Массив ", name, ": ")
			MPrint(LoadM(values))
		case "spisok1":
			fmt.Println("Односвязный список", name, "(от головы):")
			FPrint(LoadF(values))
		case "spisok2":
			fmt.Println("Двусвязный список", name, "(от головы):")
			LPrint(LoadL(values))
		case "ochered":
			fmt.Println("Очередь", name, "(от головы):")
			Q1Print(LoadQ(values))
		case "stack":
			fmt.Println("Стек", name, "(с вершины):")
			SPrint(LoadS(values))
		case "tree":
			fmt.Println("АВЛ-дерево", name, "(в ширину, работа с деревом - в C++ версии):", strings.Join(values, " "))
		}
	}
	if !found {
		Err("структура " + name + " не найдена")
	}
}

func RunHelp() {
	fmt.Print(`Запуск: ./dbms --file data.json --query 'КОМАНДА имя [аргументы]'
Массив:      MPUSH a x | MINSERT a i x | MGET a i | MSET a i x | MDEL a i | MLEN a
Односвязный: FPUSH f [HEAD|TAIL] x | FPUSH f AFTER|BEFORE t x | FPUSH f INDEX i x
             FDEL f HEAD|TAIL | FDEL f AFTER|BEFORE t | FDEL f VALUE x | FDEL f INDEX i
             FGET f i | FFIND f x | FPRINTBACK f
Двусвязный:  то же с буквой L: LPUSH, LDEL, LGET, LFIND, LPRINTBACK
Очередь:     QPUSH q x | QPOP q
Стек:        SPUSH s x | SPOP s
Общие:       PRINT имя | HELP
Деревья - только в C++ версии.
Значение с пробелами пишется в двойных кавычках: SPUSH s "hello world"
`)
}
