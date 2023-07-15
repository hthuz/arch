
#include <iostream>

using namespace std;

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

int main(){
    BST* tree = new BST();
    tree->insert(2);
    tree->insert(1);
    tree->insert(3);
    cout << tree->root->val << endl;
    cout << tree->root->left->val << endl;
    cout << tree->root->right->val << endl;

}































