
#include <netdb.h>
#include <stdio.h>
#include <stdlib.h>
#include <arpa/inet.h>


/*
* struct in_addr {
*   unsigend int s_addr // ip addr
* };
*
*/

// aton, application: 128.2.194.242
// network: 0x8002c2f2 (reversed)
int main(int argc, char* argv[])
{
    struct in_addr addr;
    struct hostent *hostp;

    if (argc != 2) {
        printf("ArgLen != 2\n");
        exit(1);
    }

    if(inet_aton(argv[1], &addr) != 0){
        // printf("%x\n",addr);
        hostp = gethostbyaddr((const char *)&addr, sizeof(addr), AF_INET);
            
        if(hostp == NULL){
            printf("official name: not found\n");
            exit(1);
        }
        printf("official name: %s\n", hostp->h_name);
        for(char** pp = hostp->h_aliases; *pp != NULL; pp++)
            printf("alias name: %s\n", *pp);
    }
}




