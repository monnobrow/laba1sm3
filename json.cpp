// формат файла:
// {
//     "massiv": { "arr": ["a", "b"] },
//     "stack":  { "s": ["дно", "вершина"] },
//     "tree":   { "t": [25, 20, 27] }
// }
// тип -> имя -> массив значений
#include <iostream>
#include <fstream>
#include <string>
#include "json.h"
using namespace std;

void InitData(Data& data){
    data.head = nullptr;
    data.tail = nullptr;
}

Entry* FindEntry(Data& data, const string& type, const string& name){
    for (Entry* cur = data.head; cur != nullptr; cur = cur->next){
        if (cur->type == type && cur->name == name){
            return cur;
        }
    }
    return nullptr;
}

Entry* AddEntry(Data& data, const string& type, const string& name){
    Entry* entry = new Entry;
    entry->type = type;
    entry->name = name;
    MInitArray(entry->values, 4);
    entry->next = nullptr;
    if (data.tail == nullptr){
        data.head = entry;
    } else {
        data.tail->next = entry;
    }
    data.tail = entry;
    return entry;
}

void ClearData(Data& data){
    while (data.head != nullptr){
        Entry* old = data.head;
        data.head = data.head->next;
        MClear(old->values);
        delete old;
    }
    data.tail = nullptr;
}

// ---------- чтение ----------
struct Parser{
    string text;
    int i;
};

void SkipSpaces(Parser& p){
    while (p.i < (int)p.text.size() && isspace((unsigned char)p.text[p.i])){
        p.i++;
    }
}
bool Expect(Parser& p, char c){
    SkipSpaces(p);
    if (p.i < (int)p.text.size() && p.text[p.i] == c){
        p.i++;
        return true;
    }
    return false;
}
bool Next(Parser& p, char c){ // следующий символ c? (не съедая его)
    SkipSpaces(p);
    return p.i < (int)p.text.size() && p.text[p.i] == c;
}

// код символа \uXXXX -> байты UTF-8
void AddUtf8(string& out, int code){
    if (code < 0x80){
        out += (char)code;
    } else if (code < 0x800){
        out += (char)(0xC0 | (code >> 6));
        out += (char)(0x80 | (code & 0x3F));
    } else {
        out += (char)(0xE0 | (code >> 12));
        out += (char)(0x80 | ((code >> 6) & 0x3F));
        out += (char)(0x80 | (code & 0x3F));
    }
}

bool ParseString(Parser& p, string& out){
    if (!Expect(p, '"')){
        return false;
    }
    out = "";
    while (p.i < (int)p.text.size()){
        char c = p.text[p.i++];
        if (c == '"'){
            return true;
        }
        if (c != '\\'){
            out += c;
            continue;
        }
        if (p.i >= (int)p.text.size()){
            return false;
        }
        char e = p.text[p.i++];
        if (e == 'n') out += '\n';
        else if (e == 't') out += '\t';
        else if (e == 'r') out += '\r';
        else if (e == 'b') out += '\b';
        else if (e == 'f') out += '\f';
        else if (e == 'u'){
            if (p.i + 4 > (int)p.text.size()){
                return false;
            }
            int code = 0;
            for (int k = 0; k < 4; ++k){
                char h = p.text[p.i++];
                code *= 16;
                if (h >= '0' && h <= '9') code += h - '0';
                else if (h >= 'a' && h <= 'f') code += h - 'a' + 10;
                else if (h >= 'A' && h <= 'F') code += h - 'A' + 10;
                else return false;
            }
            AddUtf8(out, code);
        }
        else out += e; // \" \\ \/
    }
    return false; // не нашли закрывающую кавычку
}

// значение в массиве: строка, число, true/false/null - всё храним как текст
bool ParseValue(Parser& p, string& out){
    SkipSpaces(p);
    if (Next(p, '"')){
        return ParseString(p, out);
    }
    out = "";
    while (p.i < (int)p.text.size()){
        char c = p.text[p.i];
        if (c == ',' || c == ']' || isspace((unsigned char)c)){
            break;
        }
        out += c;
        p.i++;
    }
    return !out.empty();
}

// [ "a", "b", ... ]
bool ParseArray(Parser& p, Entry* entry){
    if (!Expect(p, '[')){
        return false;
    }
    if (Expect(p, ']')){
        return true;
    }
    do {
        string value;
        if (!ParseValue(p, value)){
            return false;
        }
        MAddEnd(entry->values, value);
    } while (Expect(p, ','));
    return Expect(p, ']');
}

// { "имя": [...], ... }
bool ParseNames(Parser& p, Data& data, const string& type){
    if (!Expect(p, '{')){
        return false;
    }
    if (Expect(p, '}')){
        return true;
    }
    do {
        string name;
        if (!ParseString(p, name) || !Expect(p, ':')){
            return false;
        }
        Entry* entry = FindEntry(data, type, name);
        if (entry == nullptr){
            entry = AddEntry(data, type, name);
        }
        if (!ParseArray(p, entry)){
            return false;
        }
    } while (Expect(p, ','));
    return Expect(p, '}');
}

bool LoadJson(Data& data, const string& filename){
    ifstream in(filename);
    if (!in.is_open()){
        return true; // файла ещё нет - создадим при первом изменении
    }
    Parser p;
    string line;
    while (getline(in, line)){
        p.text += line + "\n";
    }
    p.i = 0;
    SkipSpaces(p);
    if (p.i == (int)p.text.size()){
        return true; // пустой файл
    }
    // { "тип": { ... }, ... }
    if (!Expect(p, '{')){
        return false;
    }
    if (!Expect(p, '}')){
        do {
            string type;
            if (!ParseString(p, type) || !Expect(p, ':') || !ParseNames(p, data, type)){
                return false;
            }
        } while (Expect(p, ','));
        if (!Expect(p, '}')){
            return false;
        }
    }
    SkipSpaces(p);
    return p.i == (int)p.text.size();
}

// ---------- запись ----------
void WriteString(ofstream& out, const string& s){
    out << '"';
    for (char c : s){
        if (c == '"') out << "\\\"";
        else if (c == '\\') out << "\\\\";
        else if (c == '\n') out << "\\n";
        else if (c == '\t') out << "\\t";
        else if (c == '\r') out << "\\r";
        else out << c;
    }
    out << '"';
}

void SaveJson(Data& data, const string& filename){
    ofstream out(filename);
    out << "{";
    bool first_type = true;
    // типы выводим в порядке появления, каждый тип один раз
    for (Entry* t = data.head; t != nullptr; t = t->next){
        bool seen = false;
        for (Entry* before = data.head; before != t; before = before->next){
            if (before->type == t->type){
                seen = true;
                break;
            }
        }
        if (seen){
            continue;
        }
        out << (first_type ? "\n" : ",\n") << "    ";
        WriteString(out, t->type);
        out << ": {";
        bool first_name = true;
        for (Entry* e = t; e != nullptr; e = e->next){
            if (e->type != t->type){
                continue;
            }
            out << (first_name ? "\n" : ",\n") << "        ";
            WriteString(out, e->name);
            out << ": [";
            for (int i = 0; i < MLength(e->values); ++i){
                if (i > 0){
                    out << ", ";
                }
                if (e->type == "tree"){ // в дереве числа, пишем без кавычек
                    out << MGet(e->values, i);
                } else {
                    WriteString(out, MGet(e->values, i));
                }
            }
            out << "]";
            first_name = false;
        }
        out << (first_name ? "}" : "\n    }");
        first_type = false;
    }
    out << (first_type ? "}\n" : "\n}\n");
}
