#include <iostream>
#include <string>
#include "ochered1.h"
using namespace std;

void Q1InitQueue(Queue1& q){
    q.head = nullptr;
    q.tail = nullptr;
}

void Q1PrintQueue(Queue1* q){
    Q1Node* cur_node =q->head;
    while(cur_node != nullptr){
        cout << cur_node->chislo<<"\n";
        cur_node = cur_node->next;
    }
}
void Q1Push(Queue1* q, string chislo){
    Q1Node* new_node = new Q1Node{chislo,nullptr};
    if (q->head == nullptr) {
        q->head = new_node;
        q->tail = new_node;
        return;
    }
    q->tail->next = new_node; // ссылку но новый
    q->tail = new_node;
}
string Q1Pop(Queue1* q){
    if (q->head == nullptr){
        cout << "Ошибка: очередь пуста\n";
        return "";
    }
    Q1Node* delete_node = q->head;
    string chislo = delete_node->chislo;
    q->head = q->head->next;
    delete delete_node;
    if (q->head ==nullptr){
        q->tail =nullptr;
    } // раскомент с лекции
    return chislo;
}
string Q1Front(const Queue1& q){
    return q.head->chislo;
}

string Q1Back(const Queue1& q){
    return q.tail->chislo;
}
void Q1PrintAddresses(const Queue1* q) {
    Q1Node* cur_node = q->head;
    while (cur_node != nullptr) {
        cout << cur_node->chislo << " | адрес узла: " << cur_node << " | next: " << cur_node->next << '\n';
        cur_node = cur_node->next;
    }
    cout << '\n';
}
void Q1Clear(Queue1* q){
    while (q->head != nullptr){
        Q1Pop(q);
    }
}
