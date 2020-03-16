#include <iostream>

using namespace std;

int a;

class Bubba {
public:

    ~Bubba() { cout << "d'tor" << endl;}
} ;

int main() {
    auto p=new Bubba[5];
    delete p;

}