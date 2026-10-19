#pragma once
#include <iostream>
#include <string>
using namespace std;

struct Q1Node{
    string chislo;
    Q1Node* next;
};
struct Queue1{
    Q1Node* head;
    Q1Node* tail;

};

void Q1InitQueue(Queue1& q);
void Q1PrintQueue(Queue1* q);
void Q1Push(Queue1* q, string chislo);
string Q1Pop(Queue1* q); // снимает голову и сразу отдаёт её значение
string Q1Front(const Queue1& q);
string Q1Back(const Queue1& q);
void Q1PrintAddresses(const Queue1* q);
void Q1Clear(Queue1* q);
