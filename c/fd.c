
#include <stdio.h>
#include <fcntl.h>
#include <unistd.h>

int main()
{
    int fd1,fd2, fd3;

    fd1 = open("./foo.txt",O_RDONLY);
    fd2 = open("./bar.txt",O_RDONLY);
    printf("%d\n",fd1);
    close(fd1);
    printf("%d\n",fd2);
    fd3 = open("./foo.txt",O_RDONLY);
    printf("%d\n",fd3);


}
