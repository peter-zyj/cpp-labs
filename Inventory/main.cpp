#include <iostream>
#include <queue>
using namespace std;

class Part {
public:
    Part(string _id, int _price) {
        id=_id;
        price=_price;
    }

    string GetId() {
        return id;
    }
    int GetPrice() {
        return price;
    }

private:

    string id;
    int price;
};


Part* GetMostExpensive(queue<Part*> inventory) {
    Part* mostExpensive=nullptr;
    while(!inventory.empty()) {

        Part *current = inventory.front();
        if (!mostExpensive) {
            mostExpensive = current;
            continue;
        }

        if (mostExpensive->GetPrice() <
            current->GetPrice()) {
            mostExpensive = current;
        }
        inventory.pop();
    }
    return mostExpensive;
}

int main() {
    queue<Part*> inventory;
    inventory.push(new Part{"1001", 50});
    inventory.push(new Part{"1002", 2000});
    inventory.push(new Part{"1003", 1750});
    inventory.push(new Part{"1004", 1000});

    Part* item=GetMostExpensive(inventory);
    cout << "most expensive item" << endl;
    cout << item->GetId() << endl;
    cout << item->GetPrice() << endl;
    return 0;
}