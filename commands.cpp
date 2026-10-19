// фронт: разбирает аргументы и вызывает функции структур
// поиск по target, проверки и вывод результата делают сами структуры
#include <iostream>
#include <sstream>
#include <string>
#include "commands.h"
#include "struct_cplas/spisok1.h"
#include "struct_cplas/spisok2.h"
#include "struct_cplas/ochered1.h"
#include "struct_cplas/stack.h"
#include "struct_cplas/avl_tree.h"
using namespace std;

void Say(const string& s){
    cout << "-> " << s << "\n";
}
void Err(const string& s){
    cout << "Ошибка: " << s << "\n";
}

// слова через пробел, "в кавычках" - одно слово с пробелами; пустые слова пропускаем
void Split(const string& s, Array& out){
    MInitArray(out, 4);
    int i = 0, n = s.size();
    while (i < n){
        while (i < n && isspace((unsigned char)s[i])) i++;
        if (i >= n) break;
        string word;
        if (s[i] == '"'){
            i++;
            while (i < n && s[i] != '"') word += s[i++];
            i++; // закрывающая кавычка
        } else {
            while (i < n && !isspace((unsigned char)s[i])) word += s[i++];
        }
        if (!word.empty()){
            MAddEnd(out, word);
        }
    }
}

string Upper(string s){
    for (int i = 0; i < (int)s.size(); ++i){
        s[i] = toupper((unsigned char)s[i]);
    }
    return s;
}

bool ToInt(const string& s, int& x){
    try {
        size_t pos;
        x = stoi(s, &pos);
        return pos == s.size();
    } catch (...) {
        return false;
    }
}

void CopyValues(const Array& from, Array& to){
    MInitArray(to, MLength(from));
    for (int i = 0; i < MLength(from); ++i){
        MAddEnd(to, MGet(from, i));
    }
}
bool SameValues(const Array& a, const Array& b){
    if (MLength(a) != MLength(b)){
        return false;
    }
    for (int i = 0; i < MLength(a); ++i){
        if (MGet(a, i) != MGet(b, i)){
            return false;
        }
    }
    return true;
}

string TypeOfCommand(const string& cmd){
    const char* commands[] = {
        "MPUSH", "MINSERT", "MGET", "MSET", "MDEL", "MLEN",
        "FPUSH", "FDEL", "FGET", "FFIND", "FPRINTBACK",
        "LPUSH", "LDEL", "LGET", "LFIND", "LPRINTBACK",
        "QPUSH", "QPOP", "SPUSH", "SPOP",
        "TINSERT", "TDEL", "TGET", "TPRINT"
    };
    for (const char* c : commands){
        if (cmd == c){
            switch (cmd[0]){
                case 'M': return "massiv";
                case 'F': return "spisok1";
                case 'L': return "spisok2";
                case 'Q': return "ochered";
                case 'S': return "stack";
                case 'T': return "tree";
            }
        }
    }
    return "";
}

bool IsCreateCommand(const string& cmd){
    return cmd == "MPUSH" || cmd == "MINSERT" || cmd == "FPUSH" || cmd == "LPUSH"
        || cmd == "QPUSH" || cmd == "SPUSH" || cmd == "TINSERT";
}

// ---------- перенос значений файл <-> структура ----------
void ResetValues(Array& values){
    MClear(values);
    MInitArray(values, 4);
}
void LoadF(const Array& values, FList& list){
    FInitList(list);
    for (int i = MLength(values) - 1; i >= 0; --i){ // с конца в голову - O(n)
        FAddFirst(list.head, MGet(values, i));
    }
}
void SaveF(FNode* flist, Array& values){
    ResetValues(values);
    for (; flist != nullptr; flist = flist->next){
        MAddEnd(values, flist->chislo);
    }
}
void LoadL(const Array& values, LList& list){
    LInitList(list);
    for (int i = 0; i < MLength(values); ++i){
        LAddEnd(list, MGet(values, i));
    }
}
void SaveL(const LList& list, Array& values){
    ResetValues(values);
    for (LNode* cur = list.head; cur != nullptr; cur = cur->next){
        MAddEnd(values, cur->chislo);
    }
}
void LoadQ(const Array& values, Queue1& q){
    Q1InitQueue(q);
    for (int i = 0; i < MLength(values); ++i){
        Q1Push(&q, MGet(values, i));
    }
}
void SaveQ(const Queue1& q, Array& values){
    ResetValues(values);
    for (Q1Node* cur = q.head; cur != nullptr; cur = cur->next){
        MAddEnd(values, cur->chislo);
    }
}
// стек в файле от дна к вершине: при загрузке просто кладём по порядку
void LoadS(const Array& values, Stack& stack){
    SInitStack(&stack);
    for (int i = 0; i < MLength(values); ++i){
        SPush(stack, MGet(values, i));
    }
}
void SaveSNode(SNode* node, Array& values){
    if (node == nullptr) return;
    SaveSNode(node->next, values);
    MAddEnd(values, node->chislo);
}
void SaveS(const Stack& stack, Array& values){
    ResetValues(values);
    SaveSNode(stack.head, values);
}
void LoadT(const Array& values, Tree& tree){
    TInitTree(tree);
    int chislo;
    for (int i = 0; i < MLength(values); ++i){
        if (ToInt(MGet(values, i), chislo)){
            TAdd(tree.root, chislo);
        }
    }
}
// дерево в файле в ширину: при вставке в этом порядке получится то же дерево
void SaveT(TNode* root, Array& values){
    ResetValues(values);
    string levels;
    for (int level = 0; level < THeight(root); ++level){
        TPrintLevel(root, level, levels);
    }
    istringstream in(levels);
    string word;
    while (in >> word){
        MAddEnd(values, word);
    }
}

// ---------- M: массив (values и есть массив) ----------
void RunM(const Array& q, Array& arr){
    string cmd = MGet(q, 0);
    int n = MLength(q);
    int index = 0;
    if (cmd == "MPUSH" && n == 3){
        MAddEnd(arr, MGet(q, 2));
        Say(MGet(q, 2));
    } else if ((cmd == "MINSERT" || cmd == "MSET") && n == 4){
        if (!ToInt(MGet(q, 2), index)){
            Err("индекс должен быть числом");
        } else if (cmd == "MINSERT"){
            MAddIndex(arr, MGet(q, 3), index);
        } else {
            MReplace(arr, MGet(q, 3), index);
        }
    } else if ((cmd == "MGET" || cmd == "MDEL") && n == 3){
        if (!ToInt(MGet(q, 2), index)){
            Err("индекс должен быть числом");
        } else if (cmd == "MGET"){
            MGetIndex(arr, index);
        } else {
            MDeleteIndex(arr, index);
        }
    } else if (cmd == "MLEN" && n == 2){
        Say(to_string(MLength(arr)));
    } else {
        Err("неверные аргументы, см. HELP");
    }
}

// ---------- F: односвязный список ----------
void RunF(const Array& q, Array& values){
    FList list;
    LoadF(values, list);
    string cmd = MGet(q, 0);
    int n = MLength(q);
    string mode = n > 2 ? Upper(MGet(q, 2)) : "";
    int index = 0;
    if (cmd == "FPUSH" && (n == 3 || (n == 4 && mode == "TAIL"))){
        FAddEnd(list.head, MGet(q, n - 1));
        Say(MGet(q, n - 1));
    } else if (cmd == "FPUSH" && n == 4 && mode == "HEAD"){
        FAddFirst(list.head, MGet(q, 3));
        Say(MGet(q, 3));
    } else if (cmd == "FPUSH" && n == 5 && mode == "AFTER"){
        FAddAfter(list.head, MGet(q, 3), MGet(q, 4));
    } else if (cmd == "FPUSH" && n == 5 && mode == "BEFORE"){
        FAddBefore(list.head, MGet(q, 3), MGet(q, 4));
    } else if (cmd == "FPUSH" && n == 5 && mode == "INDEX"){
        if (ToInt(MGet(q, 3), index)) FAddIndex(list.head, MGet(q, 4), index);
        else Err("индекс должен быть числом");
    } else if (cmd == "FDEL" && n == 3 && mode == "HEAD"){
        FDeleteFirst(list.head);
    } else if (cmd == "FDEL" && n == 3 && mode == "TAIL"){
        FDeleteEnd(list.head);
    } else if (cmd == "FDEL" && n == 4 && mode == "AFTER"){
        FDeleteAfter(list.head, MGet(q, 3));
    } else if (cmd == "FDEL" && n == 4 && mode == "BEFORE"){
        FDeleteBefore(list.head, MGet(q, 3));
    } else if (cmd == "FDEL" && n == 4 && mode == "VALUE"){
        FDeleteValue(list.head, MGet(q, 3));
    } else if (cmd == "FDEL" && n == 4 && mode == "INDEX"){
        if (ToInt(MGet(q, 3), index)) FDeleteIndex(list.head, index);
        else Err("индекс должен быть числом");
    } else if (cmd == "FGET" && n == 3){
        if (ToInt(MGet(q, 2), index)) FGet(list.head, index);
        else Err("индекс должен быть числом");
    } else if (cmd == "FFIND" && n == 3){
        FFind(list.head, MGet(q, 2));
    } else if (cmd == "FPRINTBACK" && n == 2){
        FPrintBack(list.head);
    } else {
        Err("неверные аргументы, см. HELP");
    }
    SaveF(list.head, values);
    FClear(list.head);
}

// ---------- L: двусвязный список ----------
void RunL(const Array& q, Array& values){
    LList list;
    LoadL(values, list);
    string cmd = MGet(q, 0);
    int n = MLength(q);
    string mode = n > 2 ? Upper(MGet(q, 2)) : "";
    int index = 0;
    if (cmd == "LPUSH" && (n == 3 || (n == 4 && mode == "TAIL"))){
        LAddEnd(list, MGet(q, n - 1));
        Say(MGet(q, n - 1));
    } else if (cmd == "LPUSH" && n == 4 && mode == "HEAD"){
        LAddFirst(list, MGet(q, 3));
        Say(MGet(q, 3));
    } else if (cmd == "LPUSH" && n == 5 && mode == "AFTER"){
        LAddAfter(list, MGet(q, 3), MGet(q, 4));
    } else if (cmd == "LPUSH" && n == 5 && mode == "BEFORE"){
        LAddBefore(list, MGet(q, 3), MGet(q, 4));
    } else if (cmd == "LPUSH" && n == 5 && mode == "INDEX"){
        if (ToInt(MGet(q, 3), index)) LAddIndex(list, MGet(q, 4), index);
        else Err("индекс должен быть числом");
    } else if (cmd == "LDEL" && n == 3 && mode == "HEAD"){
        LDeleteFirst(list);
    } else if (cmd == "LDEL" && n == 3 && mode == "TAIL"){
        LDeleteEnd(list);
    } else if (cmd == "LDEL" && n == 4 && mode == "AFTER"){
        LDeleteAfter(list, MGet(q, 3));
    } else if (cmd == "LDEL" && n == 4 && mode == "BEFORE"){
        LDeleteBefore(list, MGet(q, 3));
    } else if (cmd == "LDEL" && n == 4 && mode == "VALUE"){
        LDeleteValue(list, MGet(q, 3));
    } else if (cmd == "LDEL" && n == 4 && mode == "INDEX"){
        if (ToInt(MGet(q, 3), index)) LDeleteIndex(list, index);
        else Err("индекс должен быть числом");
    } else if (cmd == "LGET" && n == 3){
        if (ToInt(MGet(q, 2), index)) LGet(list, index);
        else Err("индекс должен быть числом");
    } else if (cmd == "LFIND" && n == 3){
        LFind(list, MGet(q, 2));
    } else if (cmd == "LPRINTBACK" && n == 2){
        LPrintBack(list);
    } else {
        Err("неверные аргументы, см. HELP");
    }
    SaveL(list, values);
    LClear(list);
}

// ---------- Q: очередь ----------
void RunQ(const Array& q, Array& values){
    Queue1 queue;
    LoadQ(values, queue);
    string cmd = MGet(q, 0);
    int n = MLength(q);
    if (cmd == "QPUSH" && n == 3){
        Q1Push(&queue, MGet(q, 2));
        Say(MGet(q, 2));
    } else if (cmd == "QPOP" && n == 2){
        string chislo = Q1Pop(&queue); // pop сразу отдаёт значение, пустую очередь проверяет сам
        if (!chislo.empty()) Say(chislo);
    } else {
        Err("неверные аргументы, см. HELP");
    }
    SaveQ(queue, values);
    Q1Clear(&queue);
}

// ---------- S: стек ----------
void RunS(const Array& q, Array& values){
    Stack stack;
    LoadS(values, stack);
    string cmd = MGet(q, 0);
    int n = MLength(q);
    if (cmd == "SPUSH" && n == 3){
        SPush(stack, MGet(q, 2));
        Say(MGet(q, 2));
    } else if (cmd == "SPOP" && n == 2){
        string chislo = SPop(stack); // pop сразу отдаёт значение, пустой стек проверяет сам
        if (!chislo.empty()) Say(chislo);
    } else {
        Err("неверные аргументы, см. HELP");
    }
    SaveS(stack, values);
    SClear(stack);
}

// ---------- T: АВЛ-дерево ----------
void RunT(const Array& q, Array& values){
    Tree tree;
    LoadT(values, tree);
    string cmd = MGet(q, 0);
    int n = MLength(q);
    int chislo = 0;
    int count = MLength(values);
    if ((cmd == "TINSERT" || cmd == "TDEL" || cmd == "TGET") && n == 3){
        if (!ToInt(MGet(q, 2), chislo)){
            Err("в дереве хранятся целые числа");
        } else if (cmd == "TINSERT"){
            TAdd(tree.root, chislo); // повтор TAdd проверяет и сообщает сам
        } else if (cmd == "TDEL"){
            TDelete(tree.root, chislo); // поиск внутри удаления
        } else {
            TFind(tree.root, chislo);
        }
    } else if (cmd == "TPRINT" && n == 3){
        string mode = Upper(MGet(q, 2));
        if (mode == "PRE") TPrintPryam(tree.root);
        else if (mode == "IN") TPrintSim(tree.root);
        else if (mode == "POST") TPrintObr(tree.root);
        else if (mode == "BFS") TPrintBfs(tree.root);
        else Err("TPRINT имя PRE|IN|POST|BFS");
        cout << "\n";
    } else {
        Err("неверные аргументы, см. HELP");
    }
    SaveT(tree.root, values);
    if (cmd == "TINSERT" && MLength(values) > count){ // узел добавился
        Say(to_string(chislo));
    }
    TClear(tree.root);
}

void RunCommand(const Array& q, Array& values){
    switch (MGet(q, 0)[0]){
        case 'M': RunM(q, values); break;
        case 'F': RunF(q, values); break;
        case 'L': RunL(q, values); break;
        case 'Q': RunQ(q, values); break;
        case 'S': RunS(q, values); break;
        case 'T': RunT(q, values); break;
    }
}

// PRINT выводит все структуры с этим именем (у разных типов имя может совпадать)
void RunPrint(Data& data, const string& name){
    const char* types[] = {"massiv", "spisok1", "spisok2", "ochered", "stack", "tree"};
    bool found = false;
    for (const char* type : types){
        Entry* e = FindEntry(data, type, name);
        if (e == nullptr){
            continue;
        }
        found = true;
        if (e->type == "massiv"){
            cout << "Массив " << name << ": ";
            MPrint(e->values);
        } else if (e->type == "spisok1"){
            FList list;
            LoadF(e->values, list);
            cout << "Односвязный список " << name << " (от головы):\n";
            FPrint(list.head);
            FClear(list.head);
        } else if (e->type == "spisok2"){
            LList list;
            LoadL(e->values, list);
            cout << "Двусвязный список " << name << " (от головы):\n";
            LPrint(list);
            LClear(list);
        } else if (e->type == "ochered"){
            Queue1 queue;
            LoadQ(e->values, queue);
            cout << "Очередь " << name << " (от головы):\n";
            Q1PrintQueue(&queue);
            Q1Clear(&queue);
        } else if (e->type == "stack"){
            Stack stack;
            LoadS(e->values, stack);
            cout << "Стек " << name << " (с вершины):\n";
            SPrint(stack.head);
            SClear(stack);
        } else if (e->type == "tree"){
            Tree tree;
            LoadT(e->values, tree);
            cout << "АВЛ-дерево " << name << " (корень слева, h - высота):\n";
            TPrintTree(tree.root);
            TClear(tree.root);
        }
    }
    if (!found){
        Err("структура " + name + " не найдена");
    }
}

void RunHelp(){
    cout << "Запуск: ./dbms --file data.json --query 'КОМАНДА имя [аргументы]'\n"
         << "Массив:      MPUSH a x | MINSERT a i x | MGET a i | MSET a i x | MDEL a i | MLEN a\n"
         << "Односвязный: FPUSH f [HEAD|TAIL] x | FPUSH f AFTER|BEFORE t x | FPUSH f INDEX i x\n"
         << "             FDEL f HEAD|TAIL | FDEL f AFTER|BEFORE t | FDEL f VALUE x | FDEL f INDEX i\n"
         << "             FGET f i | FFIND f x | FPRINTBACK f\n"
         << "Двусвязный:  то же с буквой L: LPUSH, LDEL, LGET, LFIND, LPRINTBACK\n"
         << "Очередь:     QPUSH q x | QPOP q\n"
         << "Стек:        SPUSH s x | SPOP s\n"
         << "АВЛ-дерево:  TINSERT t x | TDEL t x | TGET t x (или ISMEMBER t x) | TPRINT t PRE|IN|POST|BFS\n"
         << "Общие:       PRINT имя | HELP\n"
         << "Значение с пробелами пишется в двойных кавычках: SPUSH s \"hello world\"\n";
}
