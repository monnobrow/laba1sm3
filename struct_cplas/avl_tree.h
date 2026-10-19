#pragma once
#include <iostream>
#include <string>
using namespace std;

struct TNode{
    int data;
    int height; // высота поддерева, лист = 1
    TNode* left;
    TNode* right;
};
struct Tree{
    TNode* root;
};

void TInitTree(Tree& tree);

int THeight(TNode* root);
int TBalanceFactor(TNode* root);
void TUpdateHeight(TNode* root);
void TRotateLeft(TNode*& root);
void TRotateRight(TNode*& root);
void TBalance(TNode*& root);
// поиск идёт внутри самой вставки/удаления, ошибки печатаются тут же
void TAdd(TNode*& root, int chislo);
void TDeleteMin(TNode*& root, int& chislo);
void TDelete(TNode*& root, int chislo);
void TFind(TNode* root, int chislo);
void TPrintPryam(TNode* root);
void TPrintSim(TNode* root);
void TPrintObr(TNode* root);
void TPrintLevel(TNode* root, int level, string& out);
void TPrintBfs(TNode* root);
void TPrintTree(TNode* root, int otstup = 0);
void TPrintAddresses(TNode* root);
void TClear(TNode*& root);
