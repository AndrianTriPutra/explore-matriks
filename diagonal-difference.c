#include <stdio.h>
#include <stdint.h>
#include <stdbool.h>

int minmax(char matriks[100][100]);

uint8_t i, j;
int n;
int min, max;

int main(){
    printf("input nxn matriks, n>=2\n");
    scanf("%d", &n);
    printf("n=%d\n", n);

    if (n >= 2){
        int matrix[n][n];

        printf("\ninput the number and must be integer\n");
        printf("and give space betwen another element matriks\n");

        for (i = 0; i < n; i++){
            for (j = 0; j < n; j++){
                scanf("%d", &matrix[i][j]);
            }       
        }
        printf("\n");
        for (i = 0; i < n; i++){
            for (j = 0; j < n; j++){
                printf("%d ", matrix[i][j]);
            }
            printf("\n");
        }

        //get min max in matriks
        min = matrix[0][0];
        max = matrix[0][0];
        for (i = 0; i < n; i++){
            for (j = 0; j < n; j++){
                if (matrix[i][j] < min){
                    min = matrix[i][j];
                }
                if (matrix[i][j] > max){
                    max = matrix[i][j];
                }
            }
        }
        printf("\nmin:%d\n", min);
        printf("max:%d\n", max);

        //diagonal
        int x=0;
        int left=0;
        //left to right
        for (i = 0; i < n; i++){
            left +=  matrix[i][i];           
        }
        printf("\nsum left to right:%d\n", left);

        //right to left
        int right=0;
        j=n-1;
        for (i = 0; i < n; i++){
            right += matrix[i][j];
            j--;
        }
        printf("sum right to left:%d\n", right);
        int diffDiagonal=abs(left-right);
        printf("diffDiagonal:%d\n",diffDiagonal);

        //buttom-up
        //up
        int up =0;
        for (i = 0; i < n; i++){
            up += matrix[0][i];
        }
        printf("\nsum up:%d\n", up);
        //buttom
        int buttom =0;
        for (i = 0; i < n; i++){
            buttom += matrix[n-1][i];
        }
        printf("sum buttom:%d\n", buttom);
        int buttomup=abs(up-buttom);
        printf("diffbuttomup:%d\n",buttomup);

        //side
        //left
        int sleft=0;
        for (i = 0; i < n; i++){
            sleft += matrix[i][0];
        }
        printf("\nsleft:%d\n", sleft);
        //right
        int sright=0;
        for (i = 0; i < n; i++){
            sright += matrix[i][n-1];
        }   
        printf("sright:%d\n", sright); 
        int side=abs(sleft-sright);
        printf("diffside:%d\n\n",side);   
    }

    return 0;
}
