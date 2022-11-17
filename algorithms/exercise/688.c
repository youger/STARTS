#include<stdio.h>
#include<stdlib.h>

double pointInChessBoard(int n, int row, int column);
double currentProbability(int n, int row, int column);
double knightProbability(int n, int k, int row, int column);
int main(){
    int n, k, row, column;
    printf("input chessboard n:");
    scanf("%d", &n);
    printf("input step k:");
    scanf("%d", &k);
    printf("input row:");
    scanf("%d", &row);
    printf("input column:");
    scanf("%d", &column);
    double p = knightProbability(n, k, row, column);
    printf("%lf\n", p);
    return 0;
}

double knightProbability(int n, int k, int row, int column){
    if(k <= 0) return pointInChessBoard(n, row, column);
    double p = currentProbability(n, row, column);
    for(int i = 0; i < k; i++){
        p *= 0.125;
        for(int j = 0; j < 8; j++){
            
        }
    }
}

double currentProbability(int n, int row, int column){
    double p = 0;
    p += pointInChessBoard(n, row - 1, column - 2) * 0.125;
    p += pointInChessBoard(n, row - 1, column + 2) * 0.125;
    p += pointInChessBoard(n, row + 1, column - 2) * 0.125;
    p += pointInChessBoard(n, row + 1, column + 2) * 0.125;
    p += pointInChessBoard(n, row - 2, column - 1) * 0.125;
    p += pointInChessBoard(n, row + 2, column - 1) * 0.125;
    p += pointInChessBoard(n, row - 2, column + 1) * 0.125;
    p += pointInChessBoard(n, row + 2, column + 1) * 0.125;
    return p;
}

double pointInChessBoard(int n, int row, int column){
    if(row < n && row >= 0 && column < n && column >= 0) {
        return 1.0;
    }else{
        return 0;
    }
}
