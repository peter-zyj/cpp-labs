#include <iostream>
using namespace std;

typedef int (*ops)(int, int);

enum operation {
    a,
    d,
    m,
    s
};

int matrix[4][4]{   {operation::a, 4, 5, 0},
                    {operation::d, 0, 3, 0},
                    {operation::m, 4, 8, 0},
                    {operation::s, 1, 4, 0},
};

ops operations[4]={     [](int a, int b)->int{return a+b;},
                        [](int a, int b)->int{return a/b;},
                        [](int a, int b)->int{return a*b;},
                        [](int a, int b)->int{return a-b;}
};

int main() {
    for(auto row:matrix) {
        row[3] = (operations[row[0]])(row[1], row[2]);

        cout << row[0] << " " << row[1] << " " << row[2]
            << " " << row[3] << endl;
    }
    return 0;
}
