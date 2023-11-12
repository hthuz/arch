#include <stdio.h>

int main()
{
    int set_msb = 1 << 31;
    int beyond_msb = 1 << 32;
    int beyond_msb2 = 2 << 31;
    printf("%d\n", set_msb);
    printf("%d\n", beyond_msb);
    printf("%d\n", beyond_msb2);

}
