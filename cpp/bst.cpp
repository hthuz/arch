
#include <iostream>

using namespace std;
#define SPACE 2

class Node{
public:
    int val;
    Node* left;
    Node* right;
    Node(int val){
        this->val = val;
        this->left = NULL;
        this->right = NULL;
    }

};

class BST{
public:
    Node* root;
    BST(){
        this->root = NULL;
    }
    void insert(int);
    void _insert(Node**, Node*);
    void print();
    void _print(Node*, int);
};

void BST::insert(int val){
    Node* node = new Node(val);
    _insert(&this->root, node);
};

void BST::_insert(Node** parent, Node* node){
    if (!*parent){
        *parent = node;
        return;
    }
    if (node->val <= (*parent)->val)
        _insert(&(*parent)->left, node);
    if (node->val > (*parent)->val)
        _insert(&(*parent)->right, node);
};

void BST::print(){
    _print(this->root,SPACE);
}

void BST::_print(Node* node, int space){
    if (!node)
        return;
    space += SPACE;
    _print(node->right, space);
    cout << endl;
    for (int i = SPACE; i < space; i++)
        cout << " ";
    cout << node->val << endl;
    _print(node->left, space);
}

int main(){
    BST* tree = new BST();
    tree->insert(5);
    tree->insert(1);
    tree->insert(3);
    tree->insert(2);
    tree->insert(6);
    tree->insert(4);
    // cout << tree->root->val << endl;
    // cout << tree->root->left->val << endl;
    tree->print();

}


