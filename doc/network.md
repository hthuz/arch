# Network


## Topology

### Application Layer

HTTP, cookies, proxy server, DNS

### Transport Layer

### Network Layer
device: router
routing
IP address

### Link Layer
device: switch
MAC address
ARP (address resolution protocol)
LAN
 - Wired: Ethernet
 - Wireless: WLAN

### Physical Layer

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
