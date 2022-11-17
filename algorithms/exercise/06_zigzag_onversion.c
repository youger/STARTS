//Z 字形变换
//将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//
//比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

#include<stdio.h>
#include<stdlib.h>
#include<string.h>
#include<math.h>

char * convert(char * s, int numRows);
int main(){
    char s[] = {};
    while(1){
        printf("input string:");
        scanf("%s", s);
        printf("input row:");
        int row = 1;
        scanf("%d", &row);
        char *result = convert(s, row);
        printf("%s\n", result);
    }
    return 0;
}

char * convert(char * s, int numRows){
    if(numRows <= 1) return s;
    unsigned long len = strlen(s);
    int clumn = (int)ceil(len*1.0/2.0);
    char *str = (char*)malloc((len+1)*sizeof(char));
    str[len] = '\0';
    int interval = (numRows-1)*2;
    int i = 0;
    for(int r = numRows; r > 0; r--){
        int index = numRows - r;
        for(; index < len;){
           str[i++] = s[index];
           if(r != numRows && r != 1){
               int siblingIndex = index + (r - 1)*2;
               if(siblingIndex < len){
                   str[i++] = s[siblingIndex];
               }
           }
           index += interval; 
        }
    }
    return str;
}
