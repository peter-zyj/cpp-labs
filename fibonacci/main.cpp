#include <iostream>
using namespace std;

class Fibonacci {

public:
    Fibonacci(int _start, int _end) {
        start=_start;
        end=_end;
    }

    void Display() {
        auto n1=0;
        auto n2=1;
        while(true) {
            auto next=n1+n2;
            if (next>end){
                break;
            }
            if(next>=start) {
                cout << next << endl;
            }
            n1=n2;
            n2=next;
        }

    }

private:
    int start;
    int end;
};

int main(int argc, char *argv[]) {

    if(argc!=3) {
        std::cout << "Invalid number of parameters" << endl;
        return 0;
    }

    auto start=atol(argv[1]);
    auto end=atol(argv[2]);

    if(start > end) {
        std::cout << "Invalid parameters" << endl;
        return 0;
    }

    Fibonacci obj(start, end);
    obj.Display();

    return 0;
}