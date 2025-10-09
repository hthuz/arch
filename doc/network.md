# Network


## Topology

### Application Layer

HTTP, cookies, proxy server, DNS

### Transport Layer

### Network Layer
device: router
routing
IP address
ARP (address resolution protocol)

### Link Layer
device: switch
MAC address
LAN

 - Wired: Ethernet
 - Wireless: WLAN
 CSMA/CD

### Physical Layer


## TCP

SYN 
SYN-ACK
ACK


## HTTP

1. HTTP/0.9
Has only GET method
2. HTTP/1.0 
Add: **Header, Status code, Content-Type, POST and HEAD method**
HEAD: only get metadata about a document
POST: transfer data from client to server
Each req/resp requires opening a new connection
3. HTTP/1.1
**Persistent connection**
Add methods: PUT, PATCH, DELETE, CONNECT, TRACE, OPTIONS
4. HTTP/2.0
**Request multiplexing**: in 1.1 single request at a time. in 2.0, multiple requests at a time and receive resp asynchronously
**Automatic Gzip compressing**
**Binary protocol instead of plain text**

## (IP) Internel Protocol

`sudo pacman -S net-tools bind`

Private IP range:  
10.0.0.0 to 10.255.255.255  
172.16.0.0 to 172.31.255.255  
192.168.0.0 to 192.168.255.255  

IP: 32 bits, MAC: 48 bits

port: 65535 ports in total (16 bits).   
Restircted: 0 - 1023

Common ports:  
22: secure shell  
53: DNS
80: HTTP
179: border gateway protocol
443: HTTPS
8080: web server
3360: TCP/IP


## Check IP

| Cmd                                           | Remark                        |
|-----------------------------------------------|-------------------------------|
| `ip addr`                                     | Private IP                    |
| `ifconfig`                                    | Private IP                    |
| `host myip.opendns.com resolver1.opendns.com` | Public IP, shown at last line |

IP of router (access point) that you are connected to
In windows: default gateway
In linux: `route -n`

## Change IP

| Cmd                                          | Remark         |
|----------------------------------------------|----------------|
| `sudo ip addr add [ip_addr] dev [interface]` | Add private IP |
| `sudo ip addr del [ip_addr] dev [interface]` | Del private IP |
| Not easy to change public ip                 |                |

## NIC

To check NIC devices, ip and mac.

| Cmd            |
|----------------|
| `netstat -i`   |
| `ip link` |
| `ifconfig`     |

mtu: max transfer unit


For example, I have three NICs  
1. enp0s31f6, en. it means ethernet, bus number 0 (p0) and slot number 31 (s31)
and function nuber 6 (f6), This is ethernet interface.
2. lo, local. it means loopback device with ip addr 127.0.0.1. If you run a webserver on your machine
and browse it using browser with the same machine, it uses this device.
3. wlp2s0, wl. it means wlan(wireless local area network).  bus nubmer 2 and slot nubmer 0

## LAN

LAN usually spans a building or campus

- LAN
    - Wired: ethernet, using coaxial cable
        - bus : popular last century
        - switched: popular today
    - Wireless: WLAN, wifi

Check devices that are in the same LAN as you
1. `arp-scan --localnet`
2. `nmap <ip_range>`, e.g. `nmap 192.168.100.0-254`
3. use `ipscan`
4. use advanced ip scanner

It's better to use several tools together to find info, as some results are incomplete.  

Common Ethernet interface type: RJ45.

## IW

manipulate wireless devices

| CMD                       | Description                    |
|---------------------------|--------------------------------|
| `iw dev`                  | get name of wireless interface |
| `iw dev <interface> link` | show link status               |
| `iw dev <interface> scan` | search access points<messy info>           |


## ip cmd

| CMD                          | Description                        |
|------------------------------|------------------------------------|
| `ip link show`               | similar to ifconfig                |
| `ip link show <interface>`   | show status of a certain interface |
| `sudo ip link set <interface> up/down` | activate/deactiviate interface                 |



## Network Debug

`ping/traceroute `

`nslookup/dig` => if domain name can be resolved as ip



### DNS record

DNS resource records (RR): `(name, value, type, ttl)`

| type  | desc                                                    |
| ----- | ------------------------------------------------------- |
| A     | name: domain name, value: IPv4 address                  |
| AAAA  | name: domain name, value: IPv6 address                  |
| CNAME | name: domain name, value: canonical name of domain name |
|       |                                                         |



### nslookup

default query type: A record

```bash
nslookup www.xxx.com

# response
Server:         130.126.2.131    # DNS server used to query
Address:        130.126.2.131#53 # DNS server + port

Non-authoritative answer:		 # result from cache
www.baidu.com   canonical name = www.a.shifen.com.
www.a.shifen.com        canonical name = www.wshifen.com.
Name:   www.wshifen.com
Address: 103.235.46.102
Name:   www.wshifen.com
Address: 103.235.46.115 		# multiple servers, possibly for load babalancing
```



| Cmd                                 | Desc                    |
| ----------------------------------- | ----------------------- |
| `nslookup www.baidu.com`            | common use              |
| `nslookup <domain> <dns server ip>` | query use specified ip  |
| `nslookup -type=mx <domain>`        | query mx record         |
| `nslookup -type=ptr 8.8.8.8`        | find domain name use ip |



### dig

```bash
dig <domain> <DNS server> <type>
default type: A
```

```bash
# response
; <<>> DiG 9.20.4 <<>> www.baidu.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 30657
;; flags: qr rd ra; QUERY: 1, ANSWER: 4, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 1232
; COOKIE: 8134d6d04d65906b0100000068da24592a7898103a801b18 (good)
;; QUESTION SECTION:
;www.baidu.com.                 IN      A

;; ANSWER SECTION:
www.baidu.com.          441     IN      CNAME   www.a.shifen.com.
www.a.shifen.com.       24      IN      CNAME   www.wshifen.com.
www.wshifen.com.        233     IN      A       103.235.46.102
www.wshifen.com.        233     IN      A       103.235.46.115

;; Query time: 361 msec
;; SERVER: 130.126.2.131#53(130.126.2.131) (UDP)
;; WHEN: Mon Sep 29 14:16:56 CST 2025
;; MSG SIZE  rcvd: 161
```



### traceroute

```bash
traceroute [options] <destination>
```

| Flags          | Desc                                |
| -------------- | ----------------------------------- |
| `-n`           | skip DNS resolution                 |
| `-m <max_ttl>` | max number of hops, default 30      |
| `-q <queries>` | number of probes per hop, default 3 |
| `-I`           | use ICMP                            |
| `-T`           | use TCP                             |
| `-w <timeout>` | timeout for responses in seconds    |

### ping

send ICMP echo request



| CMD                         | Desc                                         |
| --------------------------- | -------------------------------------------- |
| `ping -c 4 www.baidu.com`   | send  4 packets                              |
| `ping -i 0.2 www.baidu.com` | time interval between packets is 0.2 seconds |
| `ping -w 5 www.baidu.com`   | set timeout to 5 seconds                     |
| `ping -s 128 www.baidu.com` | set packet size to 128 bytes                 |
| `ping -t 3 www.baidu.com`   | set TTL to 3                                 |
| `ping -4 www.baidu.com`     | force use of IPv4                            |
| `ping -6 www.baidu.com`     | force use of IPv6                            |



### telnet

```bash
telnet [domain name/ip address] port
```

replaced by ssh, typically only used for connectivity test











## Geth

geth default port: 8545  
simulation mode

| CMD                                                                                            | Comment                               |
|------------------------------------------------------------------------------------------------|---------------------------------------|
| `geth --datadir <SOME_DIR> account new`                                                        | create new account, key in ./keystore |
| `geth --datadir <SOME_DIR> account new`                                                        | create another one                    |
| `geth --datadir <SOME_DIR> --dev --password ./secret.txt`                                      |                                       |
| On anther terminal `geth attach <ipc_addr>`                                                    | ipc_addr show according to above cmd  |
| in console, `eth`                                                                              | show relevant functions               |
| `eth.accounts`                                                                                 | display accounts                      |
| `eth.coinbase`                                                                                 |                                       |
| `eth.getBalance(eth.accounts[0])`                                                              | show balance                          |
| `eth.sendTransaction({from:eth.accounts[0], to:eth.accounts[1], value:web3.toWei(1.5,"Ether")})` | transfer,node will show tx hash       |

Test to sepolia network

`curl https://raw.githubusercontent.com/prysmaticlabs/prysm/master/prysm.sh --output prysm.sh && chmod +x prysm.sh`

| CMD                                                                       | Comment              |
|---------------------------------------------------------------------------|----------------------|
| `geth --datadir <SOME_DIR>  --sepolia --http -http.api eth,net,web,admin` |                      |
| `eth.syncing`                                                             | check if syncing     |
| `eth.blockNumber`                                                         | lastest block number |
| `curl http://localhost:3500/eth/v1/node/syncing`                    | check beacon node sync status          |

JSON-RPC
```
curl -X POST -H "Content-Type: application/json"
    127.0.0.1:8545 --data \
    '{"jsonrpc":"2.0", \
    "method":eth_accounts}, \
    "params": [], \
    "id": 1 }'  \
```

curl -H "Content-Type: application/json" -X POST localhost:8545 --data '{"jsonrpc":"2.0","method":"debug_traceTransaction","params":["0x2059dd53ecac9827faad14d364f9e04b1d5fe5b506e3acc886eff7a6f88a696a"],"id":1}'

curl -H "Content-Type: application/json" -X POST localhost:8545 --data '{"jsonrpc":"2.0","method":"debug_traceBlockByNumber","params":["0x11fce84",{"disableStorage":false,"enableMemory":true,"enableReturnData":true}],"id":1}'
