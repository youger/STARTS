//最长回文子串
//给你一个字符串 s，找到 s 中最长的回文子串。
//
#include<stdlib.h>
#include<stdio.h>
#include<string.h>

char * longestPalindrome(char *s); 
int main(){
    while(1){
        char s[] = {};
        printf("input string:\n");
        scanf("%s", s);
        char *result = longestPalindrome(s);
        printf("%s\n", result);
    }
    return 1;
}

//执行用时：8 ms
//内存消耗：6.1 MB
char * longestPalindrome(char *s){
    unsigned long len = strlen(s);
    if(len <= 1) return s;
    char *p, *q;
    char *start = s;
    int maxLength = 1;
    for(int i = 0; i < len - 1; i++){
        if (maxLength >= (len - i - 1) * 2 + 1 ) {
            break;
        }
        p = s+i;
        if (s[i] == s[i+1]) {
            q = s+i+1;
            int j = 0;
            int anchor = i;
            while(j <= anchor && anchor + j < len){
                if(*(p-j) != *(q+j)){
                    break;
                }
                j++;
            }
            int curLength = j * 2;
            if(curLength > maxLength){
                start = p - j + 1;
                maxLength = curLength;
            }
        }
        if(i + 2 < len && s[i] == s[i+2]){
            q = s+i+2;
            int anchor = i + 1;
            int j = 0;
            while(j < anchor && anchor + j < len){
                if(*(p-j) != *(q+j)){
                    break;
                }
                j++;
            }
            int curLength = j * 2 + 1;
            if(curLength > maxLength){
                start = p - j + 1;
                maxLength = curLength;
            }
        }
    }
    printf("length %d\n", maxLength);
    char *res = NULL;
    if (maxLength >= 1) {
        res = (char *)malloc((maxLength + 1) * sizeof(char));
        strncpy(res, start, maxLength);
        res[maxLength] = '\0';
    }
    return res;
}

//opt
//执行用时：12 ms
//内存消耗：6 MB
char * longestPalindrome1(char * s){
    unsigned long len = strlen(s);
    if(len <= 1) return s;
    char *start = s;
    int maxLength = 1;
    for(int i = 0; i < len - 1; i++){
        if (maxLength >= (len - i - 1) * 2 + 1 ) {
            break;
        }
        if (s[i] == s[i+1]) {
            int j = 0;
            while(j <= i && i + j < len){
                if(s[i-j] != s[i+1+j]){
                    break;
                }
                j++;
            }
            int curLength = j * 2;
            if(curLength > maxLength){
                start = s + i - j + 1;
                maxLength = curLength;
            }
        }
        if(i + 2 < len && s[i] == s[i+2]){
            int anchor = i + 1;
            int j = 0;
            while(j < anchor && anchor + j < len){
                if(s[i-j] != s[i+2+j]){
                    break;
                }
                j++;
            }
            int curLength = j * 2 + 1;
            if(curLength > maxLength){
                start = s + i - j + 1;
                maxLength = curLength;
            }
        }
    }
    printf("length %d\n", maxLength);
    start[maxLength] = '\0';
    printf("%s\n", start);
    return start;
}


//opt
//执行用时：8 ms
//内存消耗：5.9 MB
char * longestPalindrome2(char * s){
    unsigned long len = strlen(s);
    if(len <= 1) return s;
    char *start = s;
    int maxLength = 1;
    for(int i = 0; i < len - 1; i++){
        if (maxLength >= (len - i - 1) * 2 + 1 ) {
            break;
        }
        int sibling = i+1;
        if (s[i] == s[sibling]) {
            int j = 0;
            while(j <= i && i + j < len){
                if(*(s+i-j) != *(s+sibling+j)){
                    break;
                }
                j++;
            }
            int curLength = j * 2;
            if(curLength > maxLength){
                start = s + i - j + 1;
                maxLength = curLength;
            }
        }
        int interval = i + 2;
        if(interval < len && s[i] == s[interval]){
            int j = 0;
            while(j < sibling && sibling + j < len){
                if(s[i-j] != s[interval+j]){
                    break;
                }
                j++;
            }
            int curLength = j * 2 + 1;
            if(curLength > maxLength){
                start = s + i - j + 1;
                maxLength = curLength;
            }
        }
    }
    printf("length %d\n", maxLength);
    start[maxLength] = '\0';
    printf("%s\n", start);
    return start;
}
