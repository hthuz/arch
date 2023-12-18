
#include <iostream>
#include <limits.h>

using namespace std;

int main()
{
    int a = INT_MAX;
    a = a + 2;
    // Overflow happends
    cout << a << " " << a + 1 << endl;
}
