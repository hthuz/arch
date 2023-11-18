
/* 
 * Udp client. It sends message from user input to server
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

using namespace std;

#define PORT 1200
#define BUFLEN 1024

int main()
{
    int s;
    // Step1: Create socket
    if((s = socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP)) == -1){
        perror("socket creation failed");
        exit(EXIT_FAILURE);
    }
    
    // Step2: Fill in server information
    struct sockaddr_in servaddr;
    memset(&servaddr, 0, sizeof(servaddr));
    servaddr.sin_family = AF_INET; // IPV4
    servaddr.sin_addr.s_addr = INADDR_ANY;
    servaddr.sin_port = htons(PORT);

    socklen_t len;
    char buf[BUFLEN];

    // Step3: Send messages to server
    while(1)
    {
        cout << "Client> ";
        cin >> buf;
        if(sendto(s, buf, sizeof(buf), 0, (struct sockaddr*)&servaddr, sizeof(servaddr)) == -1){
            perror("send failed");
            exit(EXIT_FAILURE);
        }
    }
}
