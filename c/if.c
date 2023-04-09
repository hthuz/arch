

#include <stdio.h>

int test(int a)
{
    return a * a;
}


int main()
{
    int a;
    int b;
    int c;
    a = 0;
    c = test(a);

    b = a ? 100 : 50;
    printf("%d \n",b);

}
