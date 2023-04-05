

#include <stdio.h>
#include <stdint.h>


int main()
{
    uint8_t x;
    uint8_t y;
    uint8_t z;

    x = 255;
    y = x + 1;
    z = x + 255;

    printf("%d %d \n", y, z);

}


