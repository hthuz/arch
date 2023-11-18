

#include <iostream>
using namespace std;

enum {
    EMPTY,
    HALF_FULL,
    FULL
};

enum mytype{
    AMA,
    WOW
};

int main()
{
    cout << "EMPTY is "  << EMPTY << endl;
    cout << "FULL is " << FULL << endl;
    mytype a;
    a = AMA;
    cout << a << endl;

}
