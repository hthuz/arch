
#include <iostream>
using namespace std;
class A {
public:
	int name;
	A* copy() {return this;}
};

class B : public A {
};


int main() {
	B* b = new B;

	b->name = 3;

	A* obj = b->copy();
	cout << b->copy()->name << endl;
}
