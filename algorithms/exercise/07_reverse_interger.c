//整数反转
//给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
//假设环境不允许存储 64 位整数（有符号或无符号）。
//
#include<stdio.h>
#include<stdlib.h>
#include<math.h>

int reverse(int x);
int main(){
    while(1){
        int a;
        printf("input number:");
        scanf("%d", &a);
        a = reverse(a);
        printf("output number:%d\n", a);
    }
    return 0;
}

int reverse(int x){
    int signSymbol = x > 0? 1 : -1;
    int result = 0;
    int value = abs(x);
    int maxInt = -pow(2, 31) - 2;
    printf("%d\n", maxInt);
    while(value){
        int mod = value%10;
        result = result*10 + mod;
        if(result < 0){
            result = 0;
            return result;
        }
        value = (value - mod)/10; 
    }
    result *= signSymbol;
    return result;
}
