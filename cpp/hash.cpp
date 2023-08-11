
#include <iostream>
#include <functional>
#include <unordered_map>

using namespace std;

int main()
{
    int num = 3;
    hash<int> myhash1;
    hash<int> myhash2;
    cout << myhash1(num) << endl;
    cout << myhash2(num) << endl;

    unordered_map<int, int> mymap;
    mymap[1] = 100;
    mymap[2] = 200;
    mymap[3] = 300;

    cout << mymap[1] << endl;
    cout << mymap[4] << endl;


}
