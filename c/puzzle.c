
#include <stdio.h>
#include <stdint.h>
#include <limits.h>

int main()
{

    int8_t a;
    a = INT8_MIN;
    printf("a is %d\n", a);
    printf("-a is %d\n", -a);
    printf("a == -a is %d\n", (a == -a));

    int16_t b;
    b = INT16_MIN;
    printf("b is %d\n", b);
    printf("-b is %d\n", -b);
    printf("b == -b is %d\n", (b == -b));

    int32_t c;
    c = INT32_MIN;
    uint32_t d = (uint32_t) c;
    printf("c is %d\n", c);
    printf("-c is %d\n", -c);
    printf("c == -c is %d\n", (c == -c));

    printf("d is %ud\n", d);
    printf("c == d is %d\n", (c == d));

    

}


