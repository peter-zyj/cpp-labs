#include <iostream>
#include <array>
#include <iostream>
#include <time.h>
using namespace std;

int main() {

    srand(time(0));
    int target=rand()%100;

    string resp;
    auto numof=1;
    int guess;
    string message;
    cout << "Target " << target << endl;
    while(numof < 6) {
        cout << "Guess " << numof++ << endl;
        cout << "Press <enter> to exit" << endl;
        getline(cin, resp);
        if(resp=="") {
            cout << "User stopped game";
            return 0;
        }

        guess=stol(resp);
        if(guess == target) {
            cout << "You won!!!";
            return 0;
        }
        message=guess>target?"over":"under";
        cout << "you are " << message << " target value" << endl;

    }

    cout << "You lose!";
    return 0;
}