#include <iostream>
#include <string>
#include "avl_tree.h"
using namespace std;

void TInitTree(Tree& tree){
    tree.root = nullptr;
}

int THeight(TNode* root){
    if (root == nullptr){
        return 0;
    }
    return root->height;
}

// > 1 перевес влево, < -1 перевес вправо
int TBalanceFactor(TNode* root){
    if (root == nullptr){
        return 0;
    }
    return THeight(root->left) - THeight(root->right);
}

void TUpdateHeight(TNode* root){
    int hl = THeight(root->left);
    int hr = THeight(root->right);
    if (hl > hr){
        root->height = hl + 1;
    } else {
        root->height = hr + 1;
    }
}

void TRotateLeft(TNode*& root){
    TNode* newRoot = root->right;
    root->right = newRoot->left;
    newRoot->left = root;
    TUpdateHeight(root); // старый корень теперь ниже, считаем его первым
    TUpdateHeight(newRoot);
    root = newRoot;
}

void TRotateRight(TNode*& root){
    TNode* newRoot = root->left;
    root->left = newRoot->right;
    newRoot->right = root;
    TUpdateHeight(root);
    TUpdateHeight(newRoot);
    root = newRoot;
}

// 4 случая: LL, LR, RR, RL
void TBalance(TNode*& root){
    if (root == nullptr){
        return;
    }
    TUpdateHeight(root);
    int bf = TBalanceFactor(root);
    if (bf > 1){ // левое поддерево выше
        if (TBalanceFactor(root->left) < 0){ // LR: сначала левый поворот у сына
            TRotateLeft(root->left);
        }
        TRotateRight(root); // LL
    } else if (bf < -1){ // правое поддерево выше
        if (TBalanceFactor(root->right) > 0){ // RL: сначала правый поворот у сына
            TRotateRight(root->right);
        }
        TRotateLeft(root); // RR
    }
}

// обычная вставка в BST, на обратном ходе рекурсии балансируем каждый узел пути
void TAdd(TNode*& root, int chislo){
    if (root == nullptr){
        root = new TNode{chislo, 1, nullptr, nullptr};
        return;
    }
    if (chislo < root->data){
        TAdd(root->left, chislo);
    } else if (chislo > root->data){
        TAdd(root->right, chislo);
    } else {
        cout << "Ошибка: " << chislo << " уже есть в дереве\n"; // повторы не храним
        return;
    }
    TBalance(root);
}

// вырезает минимальный узел поддерева, его значение отдаёт в chislo
void TDeleteMin(TNode*& root, int& chislo){
    if (root->left == nullptr){
        TNode* old = root;
        chislo = old->data;
        root = root->right;
        delete old;
        return;
    }
    TDeleteMin(root->left, chislo);
    TBalance(root);
}

// поиск идёт внутри удаления, отдельный TFind перед ним не нужен
void TDelete(TNode*& root, int chislo){
    if (root == nullptr){ // дошли до пустого места - числа нет
        cout << "Ошибка: " << chislo << " не найден\n";
        return;
    }
    if (chislo < root->data){
        TDelete(root->left, chislo);
    } else if (chislo > root->data){
        TDelete(root->right, chislo);
    } else {
        cout << "-> " << chislo << "\n";
        if (root->left != nullptr && root->right != nullptr){
            // два сына: на место узла ставим минимум правого поддерева
            int min_chislo;
            TDeleteMin(root->right, min_chislo);
            root->data = min_chislo;
        } else {
            // 0 или 1 сын: подвешиваем сына на место удаляемого
            TNode* old = root;
            if (root->left != nullptr){
                root = root->left;
            } else {
                root = root->right;
            }
            delete old;
            return; // сын уже сбалансирован
        }
    }
    TBalance(root);
}

void TFind(TNode* root, int chislo){
    while (root != nullptr){
        if (chislo == root->data){
            cout << "-> TRUE\n";
            return;
        }
        if (chislo < root->data){
            root = root->left;
        } else {
            root = root->right;
        }
    }
    cout << "-> FALSE\n";
}

// preorder: корень, левое, правое
void TPrintPryam(TNode* root){
    if (root == nullptr) {return;}
    cout << root->data << " ";
    TPrintPryam(root->left);
    TPrintPryam(root->right);
}
// inorder: левое, корень, правое (по возрастанию)
void TPrintSim(TNode* root){
    if (root == nullptr) {return;}
    TPrintSim(root->left);
    cout << root->data << " ";
    TPrintSim(root->right);
}
// postorder: левое, правое, корень
void TPrintObr(TNode* root){
    if (root == nullptr) {return;}
    TPrintObr(root->left);
    TPrintObr(root->right);
    cout << root->data << " ";
}

// дописывает в out все узлы уровня level слева направо
void TPrintLevel(TNode* root, int level, string& out){
    if (root == nullptr) {return;}
    if (level == 0){
        out += to_string(root->data) + " ";
        return;
    }
    TPrintLevel(root->left, level - 1, out);
    TPrintLevel(root->right, level - 1, out);
}
// обход в ширину по уровням
void TPrintBfs(TNode* root){
    string out;
    for (int level = 0; level < THeight(root); ++level){
        TPrintLevel(root, level, out);
    }
    cout << out;
}

// дерево боком: правое поддерево сверху, корень слева
void TPrintTree(TNode* root, int otstup){
    if (root == nullptr) {return;}
    TPrintTree(root->right, otstup + 1);
    for (int i = 0; i < otstup; ++i){
        cout << "        ";
    }
    cout << root->data << "(h=" << root->height << ")\n";
    TPrintTree(root->left, otstup + 1);
}

void TPrintAddresses(TNode* root){
    if (root == nullptr) return;
    cout << "Значение: " << root->data << " | h: " << root->height << " | адрес: " << root << " | left: " << root->left << " | right: " << root->right << '\n';
    TPrintAddresses(root->left);
    TPrintAddresses(root->right);
}

void TClear(TNode*& root){
    if (root == nullptr) return;
    TClear(root->left);
    TClear(root->right);
    delete root;
    root = nullptr;
}
