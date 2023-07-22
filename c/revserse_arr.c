


#include <stdio.h>

void reverse(int*, int);

void reverse(int* array, int array_len)
{
    int tmp;
    for(int i = 0; i < array_len / 2; i++){
        tmp = array[i];
        array[i] = array[array_len - 1 - i];
        array[array_len - 1 - i] = tmp;
    }

};


int main()
{

    int array[] = {3,5,1,3,6,7,9};
    int len = 7;
    reverse(array, len);

    for(int i = 0; i < len; i++)
        printf("%d ", array[i]);
    printf("\n");

};
