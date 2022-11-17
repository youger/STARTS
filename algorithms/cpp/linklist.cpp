#include<iostream>

using namespace std;

template<typename T>
struct LinkList
{    
    T value;
    struct LinkList *next;
    
    LinkList(T val){
        value = val;
        next = nullptr;
    }
    
    ~LinkList(){
        delete next;
    }
    
    void insert(T value);
    
    void remove(T value);
    
    void travel();
};

//typedef struct LinkList LinkList;

template<typename T>
void LinkList<T> :: insert(T value)
{
    LinkList *p = this->next;
    LinkList *tail = this;
    while(p != NULL){
        tail = p;
        p = p->next;
    }
    tail->next = new LinkList(value);
}

template<typename T>
void LinkList<T> :: remove(T value)
{
    LinkList *p = this->next;
    LinkList *pre = this;
    while(p != NULL){
        if (p->value == value)
        {
            pre->next = p->next;
            p->next = nullptr;
            delete p;
            break;
        }
        pre = p;
        p = p->next;
    }
}

template<typename T>
void LinkList<T> :: travel() 
{
    LinkList *p = this;
    while(p != NULL){
        cout << p->value << "->";
        p = p->next;
    }
    cout << "null" << endl;
}

int main()
{
    LinkList<int> lkls = LinkList<int>(0);
    lkls.insert(1);
    lkls.insert(2);
    lkls.insert(3);
    lkls.insert(4);
    
    cout << "Integer linklist: "<<endl;
    lkls.travel();

    cout<<"remove value=3\n";
    lkls.remove(3);
    lkls.travel();
    cout<<"remove value=0\n";
    lkls.remove(1);
    lkls.travel();

    LinkList<string> lkls2 = LinkList<string>("head");
    lkls2.insert("aba");
    lkls2.insert("bd");
    lkls2.insert("cfe");
    lkls2.insert("out");
    
    cout << endl;
    cout << "String linklist: "<<endl;
    lkls2.travel();
    cout<<"remove value=\"b\"\n";
    lkls2.remove("b");

    lkls2.travel();
    return 0;
}
