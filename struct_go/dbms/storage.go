package main

// формат файла тот же, что у C++ версии:
// { "тип": { "имя": [значения] } }
// дерево ("tree") Go не изменяет, но сохраняет как было

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Data map[string]map[string][]any

// false-ошибка, если файл есть, но это не JSON нужного вида (тогда файл не трогаем)
func LoadJson(filename string) (Data, error) {
	data := Data{}
	text, err := os.ReadFile(filename)
	if errors.Is(err, os.ErrNotExist) || (err == nil && strings.TrimSpace(string(text)) == "") {
		return data, nil // создадим при первом изменении
	}
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(bytes.NewReader(text))
	dec.UseNumber() // числа дерева оставляем числами
	if err := dec.Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func SaveJson(data Data, filename string) error {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "    ")
	if err := enc.Encode(data); err != nil {
		return err
	}
	return os.WriteFile(filename, buf.Bytes(), 0644)
}

// значения из файла как строки
func GetValues(data Data, typ, name string) ([]string, bool) {
	raw, ok := data[typ][name]
	if !ok {
		return nil, false
	}
	values := make([]string, len(raw))
	for i, v := range raw {
		if s, isString := v.(string); isString {
			values[i] = s
		} else {
			values[i] = fmt.Sprint(v)
		}
	}
	return values, true
}

func SetValues(data Data, typ, name string, values []string) {
	if data[typ] == nil {
		data[typ] = map[string][]any{}
	}
	raw := make([]any, len(values))
	for i, v := range values {
		raw[i] = v
	}
	data[typ][name] = raw
}
