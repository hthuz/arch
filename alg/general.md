## Database

事务ACID

原子性

一致性， 事务前后数据库都满足完整性约束

隔离性 isolation

持久性 durability



如果不设置事务

- 可能table 1 changed，但table 2 unchanged, 造成问题
- race condition, 查询时两线程都发现库存 = 1, 因此两次操作，导致库存更新为-1
- 执行了一半，发现问题了，但无法回滚，只能自己手动补偿，但很易错



事务并发可能的问题：

- 更新丢失：两个tx, 后tx覆盖先tx的结果

- 脏读：一个tx读了另一个tx未提交的数据， tx1修改数据，tx2读取数据，此时tx1未结束但数据已被修改，此时tx1回滚，导致tx2读的数据和数据库的数据不一致

- 不可重复读：one tx读两次相同数据，结果不一致，因为后tx进行了更新操作

- 幻象读（phantom read）





参考完整性 reference integrity

索引创建在

- 经常需要搜索的列
- 作为主键的列
- 常用于join的列（foreign key）
- 常需要范围搜索的列 （如block num）
- 常需要排序的列，加快排序
- 常用于where的列
- 





## HTTPS (SSL/TLS)

SSL: Secure Sockets Layer

TLS: Transport Layer Security



Client symmetric key X)

Server (private key A, public key A, 

对X进行非加密传输后，Client Server都知道X的信息，之后的内容都用X加密进行传输

Client -> request -> server	

Server -> public key A -> client

client -> pubA(X) -> server

server  decrypt, get X



需要认证公钥A是CA认证的（数字证书），才能防止man in the middle attack,因为middle man可以欺骗伪造其自己的公钥和私钥



但是数字证书也可能被篡改，因此需要digital signature = privateCA(hash(plain text))

digital certificate = 明文 + 数字签名， 用数字签名验证明文hash是否与数字签名解密后的hash一致

浏览器保有CA的公钥

![https](/home/autentico/arch/alg/img/https.jpg)

整体流程

1. Client向server发送https请求
2. server向client发送数字证书（公钥，企业信息等）
3. 客户端验证数字证书（用公钥解密并对比hash）
4. server生成会话密钥（symmetric key）
5. client server加密会话





| Layer       | Items                                                        | Unit                                |
| ----------- | ------------------------------------------------------------ | ----------------------------------- |
| Application |                                                              | Message/Data                        |
| Transport   | 端口，网关                                                   | segment段 - TCP；datagram数据报-UDP |
| Network     | IP, 路由器，ARP (Address Resolution Protocol translate from IP to MAC )， ICMP （用于ping） | packet包；数据报datagram            |
| Link        | MAC.  交换器， , LAN， CSMA/CD， MTU                         | Frame 帧                            |
| Physics     | 集线器hub,中继器repeater                                     | 比特bit                             |

路由器：NAT，访问控制，连接不同网络，寻址



- Distance Vector 距离矢量  （bellman ford）
  - RIP, BGP， IGRP

- Link State 链路状态 （dijkstra）

  - OSPF，IS-IS

  



选择重传：selective repeat protocols



## 交换变量

```
a = 5, b = 15
a = a + b
b = a - b
a = a - b


a = a * b
b = a / b
a = a / b

a = a ^ b
b = a ^ b
a = a ^ b

(a ^ a = 0, a ^ 0 = a)

```

