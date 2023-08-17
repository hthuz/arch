
#include <iostream>
#include <bitset>
using namespace std;



int main()
{
    int n = -2;
    bitset<64> b(n);
    bitset<64> comp(~n + 1);

    cout << n << endl;
    cout << b << endl;
    // Flip over every bits
    cout << ~n << endl;

    // 2'complement number (opposite number)
    cout << ~n + 1 << endl;
    cout << comp << endl;
    

    // negative number's operation
    cout << n / 2 << endl;    // -1
    cout << (n >> 1) << endl; // -1

    cout << n / 4 << endl;    // 0
    cout << (n >> 2) << endl; // -1
    //
    cout << n / 3 << endl;    // 0
    //
    cout << n % 2 << endl;   // 0
    cout << n % 3 << endl;   // -2
    cout << n % 4 << endl;  // -2
    cout << n % 5 << endl;   // -2

}

