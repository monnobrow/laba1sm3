#include <iostream>
#include <string>
#include "spisok1.h"
using namespace std;

void FInitList(FList& list){
    list.head = nullptr;
}

void FAddFirst(FNode*& flist, string chislo){
    FNode* new_node = new FNode{chislo, flist};
    flist = new_node;
}
void FAddEnd(FNode*& flist, string chislo){
    FNode* new_node = new FNode{chislo, nullptr};
    if (flist == nullptr) {
        flist = new_node;
        return;
    }
    FNode* cur_node = flist;
    while(cur_node->next !=nullptr){
        cur_node = cur_node->next;
    }
    cur_node->next = new_node;
}
// нашли target и сразу вставили, второго прохода нет
void FAddAfter(FNode* flist, string target, string chislo){
    while (flist != nullptr && flist->chislo != target){
        flist = flist->next;
    }
    if (flist == nullptr){
        cout << "Ошибка: " << target << " не найден\n";
        return;
    }
    flist->next = new FNode{chislo, flist->next};
    cout << "-> " << chislo << "\n";
}
// идём на шаг позади: cur_node->next - кандидат на target
void FAddBefore(FNode*& flist, string target, string chislo){
    if (flist != nullptr && flist->chislo == target){ // перед головой = в начало
        FAddFirst(flist, chislo);
        cout << "-> " << chislo << "\n";
        return;
    }
    FNode* cur_node = flist;
    while (cur_node != nullptr && cur_node->next != nullptr && cur_node->next->chislo != target){
        cur_node = cur_node->next;
    }
    if (cur_node == nullptr || cur_node->next == nullptr){
        cout << "Ошибка: " << target << " не найден\n";
        return;
    }
    cur_node->next = new FNode{chislo, cur_node->next};
    cout << "-> " << chislo << "\n";
}
void FAddIndex(FNode*& flist, string chislo, int index){
    if (index == 0) {
        FAddFirst(flist, chislo);
        cout << "-> " << chislo << "\n";
        return;
    }
    FNode* cur_node = flist;
    for(int i=0; i<index-1 && cur_node != nullptr;++i){
        cur_node = cur_node->next;
    }
    if (index < 0 || cur_node == nullptr){
        cout << "Ошибка: нет такого индекса\n";
        return;
    }
    cur_node->next = new FNode{chislo, cur_node->next};
    cout << "-> " << chislo << "\n";
}

// вырезает узел, стоящий после prev
void FDeleteNext(FNode* prev){
    FNode* delete_node = prev->next;
    cout << "-> " << delete_node->chislo << "\n";
    prev->next = delete_node->next;
    delete delete_node;
}
void FDeleteFirst(FNode*& flist){
    if (flist == nullptr){
        cout << "Ошибка: список пуст\n";
        return;
    }
    FNode* delete_node = flist;
    cout << "-> " << delete_node->chislo << "\n";
    flist = flist->next;
    delete delete_node;
}
void FDeleteEnd(FNode*& flist){
    if (flist == nullptr || flist->next == nullptr){
        FDeleteFirst(flist);
        return;
    }
    FNode* cur_node = flist;
    while (cur_node->next->next != nullptr){
        cur_node = cur_node->next;
    }
    FDeleteNext(cur_node);
}
void FDeleteAfter(FNode* flist, string target){
    while (flist != nullptr && flist->chislo != target){
        flist = flist->next;
    }
    if (flist == nullptr || flist->next == nullptr){
        cout << "Ошибка: после " << target << " нечего удалять\n";
        return;
    }
    FDeleteNext(flist);
}
// удаляем узел перед target: идём на два шага позади
void FDeleteBefore(FNode*& flist, string target){
    if (flist == nullptr || flist->chislo == target){ // перед головой ничего нет
        cout << "Ошибка: перед " << target << " нечего удалять\n";
        return;
    }
    if (flist->next != nullptr && flist->next->chislo == target){
        FDeleteFirst(flist);
        return;
    }
    FNode* cur_node = flist;
    while (cur_node->next != nullptr && cur_node->next->next != nullptr
           && cur_node->next->next->chislo != target){
        cur_node = cur_node->next;
    }
    if (cur_node->next == nullptr || cur_node->next->next == nullptr){
        cout << "Ошибка: перед " << target << " нечего удалять\n";
        return;
    }
    FDeleteNext(cur_node);
}
void FDeleteIndex(FNode*& flist, int index){
    if (index == 0){
        FDeleteFirst(flist);
        return;
    }
    FNode* cur_node = flist;
    for(int i =0; i< index -1 && cur_node != nullptr;++i){
        cur_node = cur_node->next;
    }
    if (index < 0 || cur_node == nullptr || cur_node->next == nullptr){
        cout << "Ошибка: нет такого индекса\n";
        return;
    }
    FDeleteNext(cur_node);
}
void FDeleteValue(FNode*& flist, string chislo){
    if (flist != nullptr && flist->chislo == chislo){
        FDeleteFirst(flist);
        return;
    }
    FNode* cur_node = flist;
    while (cur_node != nullptr && cur_node->next != nullptr && cur_node->next->chislo != chislo){
        cur_node = cur_node->next;
    }
    if (cur_node == nullptr || cur_node->next == nullptr){
        cout << "Ошибка: " << chislo << " не найден\n";
        return;
    }
    FDeleteNext(cur_node);
}

void FGet(FNode* flist, int index){
    for (int i = 0; i < index && flist != nullptr; ++i){
        flist = flist->next;
    }
    if (index < 0 || flist == nullptr){
        cout << "Ошибка: нет такого индекса\n";
        return;
    }
    cout << "-> " << flist->chislo << "\n";
}
void FFind(FNode* flist, string chislo){
    while (flist != nullptr){
        if (flist->chislo == chislo){
            cout << "-> TRUE\n";
            return;
        }
        flist = flist->next;
    }
    cout << "-> FALSE\n";
}

void FPrintBack(FNode* flist){
    if (flist == nullptr) return;
    FPrintBack(flist->next);
    cout << flist->chislo << "\n";
}
void FPrint(FNode* flist){
    while(flist != nullptr){
        cout << flist->chislo<<"\n";
        flist = flist->next;
    }
}
void FPrintAddresses(FNode* flist) {
    while (flist != nullptr) {
        cout << "Узел: " << flist << ", chislo: " << flist->chislo << ", next: " << flist->next << '\n';
        flist = flist->next;
    }
}
void FClear(FNode*& flist){
    while (flist != nullptr){
        FNode* delete_node = flist;
        flist = flist->next;
        delete delete_node;
    }
}
