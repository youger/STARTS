#include<stdio.h>

char **all_sorted_string(char *str, int *count);
int main(){
    while(1){
        char str[];
        printf("input string:\n");
        scanf("%s", str);
        int count;
        char **str_comps = all_sorted_string(str, &count);
        for(int i = 0; i < count; i++){
            printf("%s\n", str_comps[i]);
        }
        return 0;
    }
}

char **all_sorted_string(char *str, int *count){


}



