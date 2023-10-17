
#include <iostream>
using namespace std;

class Animal {
protected:
    int weight;
    int height;
    int age;

public:
    Animal() {
        weight = 0;
        height = 0;
        age = 0;
    }
    Animal(int i_weight, int i_height, int i_age) : weight(i_weight), height(i_height), age(i_age){} ;

    virtual void eat() {
        weight += 10;
        height += 10;
    }

    virtual void bark() {
        cout << "Making some sound " << endl;
    }

    virtual void print() { 
        cout << "Animal weight:" << weight << " height:" << height << " age: " << age << endl;
    }

    virtual ~Animal() {
        cout << "Animal Destructed" << endl;
    }
};

class Dog : public Animal{

private:
    int aggresive;
public:
    Dog(int i_aggresive) : aggresive(i_aggresive) {};
    void eat() {
        weight += 15;
        height += 15;
        age += 1;
    }
    void bark() {
        cout << "Woof!" << endl;
    }
    void print() { 
        cout << "Dog weight:" << weight << " height:" << height << " age: " << age << " aggresive: " << aggresive << endl;
    }
    ~Dog() {
        cout << "Dog Destructed" << endl;
    }

};

class Cat : public Animal {

public:
    void bark() {
        cout << "Meow!" << endl;
    }

};

int main() 
{
    Animal* animal = new Dog(3);
    // The following is forbidden
    // Dog* dog = new Animal();
    //
    // With virtual in Animal::print(), it calls print in Dog
    // Without virtual in Animal::print(), it will calls print in Animal
    // We may need to overwrite some functions in derived classes and virtual is used to achieve this.
    animal->print();

    Animal* dog = new Dog(5);
    Animal* cat = new Cat();
    dog->bark();
    cat->bark();

    delete animal;
    return 0;

}
