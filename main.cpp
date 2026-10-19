#include <iostream>
#include <string>
#include "json.h"
#include "commands.h"
using namespace std;

int main(int argc, char* argv[]){
    string filename, query;
    for (int i = 1; i < argc; ++i){
        string arg = argv[i];
        if (arg == "--file" && i + 1 < argc){
            filename = argv[++i];
        } else if (arg == "--query" && i + 1 < argc){
            query = argv[++i];
        }
    }
    if (filename.empty() || query.empty()){
        cout << "Запуск: ./dbms --file data.json --query 'КОМАНДА имя [аргументы]'\n";
        return 1;
    }

    Array q;
    Split(query, q);
    if (MLength(q) == 0){
        cout << "Ошибка: пустой запрос\n";
        return 1;
    }
    string cmd = Upper(MGet(q, 0));
    if (cmd == "ISMEMBER"){ // как в примере из методички
        cmd = "TGET";
    }
    q.data[0] = cmd;
    if (cmd == "HELP"){
        RunHelp();
        return 0;
    }
    string type = TypeOfCommand(cmd);
    if (cmd != "PRINT" && type.empty()){
        cout << "Ошибка: неизвестная команда " << cmd << ", см. HELP\n";
        return 1;
    }
    if (MLength(q) < 2){
        cout << "Ошибка: не указано имя структуры\n";
        return 1;
    }
    string name = MGet(q, 1);

    Data data;
    InitData(data);
    if (!LoadJson(data, filename)){
        cout << "Ошибка: " << filename << " - не JSON нужного вида, файл не изменён\n";
        return 1;
    }

    if (cmd == "PRINT"){
        RunPrint(data, name);
    } else {
        Entry* entry = FindEntry(data, type, name);
        if (entry == nullptr && !IsCreateCommand(cmd)){
            cout << "Ошибка: структура " << name << " не найдена\n";
            ClearData(data);
            return 1;
        }
        if (entry == nullptr){ // новая структура, в файл попадёт только если команда прошла
            entry = AddEntry(data, type, name);
        }
        Array before;
        CopyValues(entry->values, before);
        RunCommand(q, entry->values);
        if (!SameValues(before, entry->values)){
            SaveJson(data, filename); // перезаписываем только при изменении
        }
        MClear(before);
    }
    ClearData(data);
    MClear(q);
    return 0;
}
