#include<iostream>

using namespace std;

template<typename T>
struct BTree
{
    T value;
    BTree *left;
    BTree *right;

    BTree(T val){
        value = val; 
        left = nullptr;
        right = nullptr;
    }

    ~BTree(){
        delete left, right;
    }

    void travel();
    void removeLeaf(T value);
    unsigned int height();
}

template<typename T>
unsigned int BTree<T> :: height() 
{

}

template<typename T>
void BTree<T> :: travel() 
{
    
}

template<typename T>
void BTree<T> :: removeLeaf(T value) 
{
    
}

BTree<int> *createBinaryTree(vector<int> nums)
{
    
}

int main()
{
    BTree<int> *root = createBinaryTree();
    cout<<"binary tree height:"<<root->height()<<endl;
    root->travel();
    return 1;
}
