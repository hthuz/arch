
#include <stdlib.h>
#include <stdio.h>


void my_malloc(int* ptr)
{
    ptr = malloc(sizeof(int));
    *ptr = 3;
    printf("%d\n",*ptr);
}

int main()
{
    int* ptr;
    my_malloc(ptr);
    printf("%d\n",*ptr);
    // free(ptr);

}
