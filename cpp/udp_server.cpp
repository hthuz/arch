

/* 
 * Udp server. It receives message from client
 *
 *
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <iostream>
#include <sys/socket.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <fcntl.h>

using namespace std;

#define PORT 1200
#define BUFLEN 1024

int main(){

    int s;
    // Step1: Create socket
    if((s = socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP)) == -1){
        perror("socket creation failed");
        exit(EXIT_FAILURE);
    }

    struct sockaddr_in servaddr;
    struct sockaddr_in cliaddr;
    // Step2: Fill in server information, client info not filled
    memset((char*) &servaddr, 0, sizeof(servaddr));
    servaddr.sin_family = AF_INET; // IPV4
    servaddr.sin_addr.s_addr = INADDR_ANY;
    servaddr.sin_port = htons(PORT);

    // Step3: Binds mean to associate a socket with an IP and port
    if(bind(s,(struct sockaddr *)&servaddr, sizeof(servaddr)) == -1) {
        perror("bind failed");
        exit(EXIT_FAILURE);
    }

    // Optional 1: set option for timeout
    // timeval timeout;
    // timeout.tv_sec = 1;
    // timeout.tv_usec = 0;
    // setsockopt(s, SOL_SOCKET, SO_RCVTIMEO, &timeout,sizeof(timeout)) ;

    // Optional 2: set non-blocking socket
    // fcntl(s, F_SETFL, O_NONBLOCK);

    socklen_t len;
    int n; // Size read from client
    char buf[BUFLEN];
    // Step4: receive message from client
    cout << "Server started" << endl;
    while(true)
    {
        if((n = recvfrom(s, &buf, BUFLEN,0, (struct sockaddr* )&cliaddr, &len)) == -1){
            perror("recv fail");
            exit(EXIT_FAILURE);
        }
        buf[n] = '\0';
        cout << "Recv> " << buf << endl;
    }

}
