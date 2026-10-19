#pragma once
#include <iostream>
#include <string>
#include "json.h"
using namespace std;

void Split(const string& s, Array& out);
string Upper(string s);
void CopyValues(const Array& from, Array& to);
bool SameValues(const Array& a, const Array& b);

// тип структуры в файле по команде: SPUSH -> "stack"; "" - такой команды нет
string TypeOfCommand(const string& cmd);
bool IsCreateCommand(const string& cmd);
// q - слова запроса, values - значения структуры из файла (меняются на месте)
void RunCommand(const Array& q, Array& values);
void RunPrint(Data& data, const string& name);
void RunHelp();
