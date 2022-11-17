int foo(int *p);
int main(){
    int *p = 0;
    foo(p);
}

int foo(int *p){
    int y = *p;
    return y;
}
