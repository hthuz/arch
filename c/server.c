
#include <stdio.h>
#include <sys/socket.h>
#include <netinet/in.h>

#define PORT
int main() 
{
    int sockfd;
    struct sockaddr_in addr; // require netinet/in.h
    
    struct addrinfo *

    sockfd = socket(AF_INET, SOCK_STREAM, 0);
    printf("%d\n",sockfd);

}
