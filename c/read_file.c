
// Read contents of a file to a char array
#include <stdio.h>
#include <stdlib.h>

int main()
{
    FILE* f = fopen("test.html","r");
    fseek(f,0, SEEK_END);
    unsigned long len = (unsigned long) ftell(f); 
    fseek(f,0, SEEK_SET);
    // Length counts the last 
    printf("File Length is %lu\n", len);

    // The last character will be \n
    char* s = (char *)malloc(len * sizeof(char));
    fread(s, len, 1, f);

    printf("%s\n",s);

}
