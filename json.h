#pragma once
#include <iostream>
#include <string>
#include "struct_cplas/massiv.h"
using namespace std;

// одна именованная структура из файла: тип ("stack"), имя ("s") и значения
struct Entry{
    string type;
    string name;
    Array values;
    Entry* next;
};
struct Data{
    Entry* head;
    Entry* tail;
};

void InitData(Data& data);
Entry* FindEntry(Data& data, const string& type, const string& name);
Entry* AddEntry(Data& data, const string& type, const string& name);
void ClearData(Data& data);

// false - файл есть, но это не JSON нужного вида (тогда файл не трогаем)
bool LoadJson(Data& data, const string& filename);
void SaveJson(Data& data, const string& filename);
