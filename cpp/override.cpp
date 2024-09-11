
#include <iostream>


class Expression_class {
public:
	// Without virtual, the program will gives 0; => Based on declared type
	// With virtual => Based on derived classes
	virtual bool is_no_expr() {return false;}
};

typedef Expression_class* Expression;

class no_expr_class : public Expression_class {
public:
	bool is_no_expr() {return true;}
};

int main() {
	// Compiler makes decision based on static type information available at compile time
	Expression expr = new no_expr_class();
	std::cout << expr->is_no_expr() << std::endl;
}
