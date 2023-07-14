
#include <stdio.h>
#include <sys/types.h>

int main(){

    pid_t pid;

    // On success, pid of child is returned in parent. 0 is return in child,
    // On failure, -1 is returned in parent
    pid = fork();

    if (pid > 0){
        printf("PARENT\n");
        printf("Parent's pid is %d\n", getppid());
        printf("My's pid is %d\n", getpid());
        printf("Returned value is %d\n", pid);
    }

    if (pid == 0){
        printf("CHILD\n");
        printf("Parent's pid is %d\n", getppid());
        printf("My's pid is %d\n", getpid());
        printf("Returned value is %d\n", pid);
    }
        // printf("Parent's pid is %d\n", getppid());
        // printf("My's pid is %d\n", getpid());
        // printf("Returned value is %d\n", pid);

}
