#pragma once
#include <iostream>
#include <string>
using namespace std;

struct FNode{
    string chislo;
    FNode* next;
};
struct FList{
    FNode* head;
};

void FInitList(FList& list);

// поиск target делается внутри операции, за один проход по списку
// результат ("-> x") или ошибку операция печатает сама
void FAddFirst(FNode*& flist, string chislo);
void FAddEnd(FNode*& flist, string chislo);
void FAddAfter(FNode* flist, string target, string chislo);
void FAddBefore(FNode*& flist, string target, string chislo);
void FAddIndex(FNode*& flist, string chislo, int index);
void FDeleteNext(FNode* prev);
void FDeleteFirst(FNode*& flist);
void FDeleteEnd(FNode*& flist);
void FDeleteAfter(FNode* flist, string target);
void FDeleteBefore(FNode*& flist, string target);
void FDeleteIndex(FNode*& flist, int index);
void FDeleteValue(FNode*& flist, string chislo);
void FGet(FNode* flist, int index);
void FFind(FNode* flist, string chislo);
void FPrintBack(FNode* flist);
void FPrint(FNode* flist);
void FPrintAddresses(FNode* flist);
void FClear(FNode*& flist);
