## Database



### 回表查询

**回表的根源**：普通索引的叶子节点不存储完整数据，只存储主键ID。

就是先通过普通索引（也叫二级索引）找到主键值，再根据这个主键值回到主键索引（也叫聚簇索引）的叶子节点中，去查找完整数据行的一个过程。

假设我们有一张用户表 `user`：

sql

```sql
CREATE TABLE user (
  id INT PRIMARY KEY,          -- 主键
  name VARCHAR(50),            -- 姓名
  age INT,                     -- 年龄
  INDEX idx_name (name)        -- 为name字段创建了一个普通索引
);
```

执行如下查询：**

sql

```sql
SELECT * FROM user WHERE name = 'Alice';
```



这个查询的流程如下：

1. **在普通索引 `idx_name` 中查找：**
   - 数据库在 `idx_name` 这棵 B+Tree 里搜索 `name = 'Alice'`。
   - 在叶子节点找到了记录：`('Alice', 主键ID=18)`。
   - **注意：** 这里只拿到了 `name` 和 `id`，但 `SELECT *` 要求所有字段（比如 `age`）。
2. **“回表”到主键索引：**
   - 数据库拿着上一步得到的主键值 `id = 18`。
   - 回到主键索引的 B+Tree 中，去查找 `id = 18` 的这条记录。
3. **在主键索引中获取完整数据：**
   - 在主键索引的叶子节点中，找到了 `id = 18` 对应的**完整数据行** (`id=18, name='Alice', age=25`)。
   - 将这条完整的数据行返回。



避免回表需要**覆盖索引**

### 主从复制

### 读写分离



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

- **脏读 (dirty read)**：一个tx读了另一个tx未提交的数据， tx1修改数据，tx2读取数据，此时tx1未结束但数据已被修改，此时tx1回滚，导致tx2读的数据和数据库的数据不一致

- **不可重复读 (non-repeatable read)**：one tx读两次相同数据，结果不一致，因为后tx进行了更新操作

- **幻象读（phantom read）**:事务A在读取数据后，事务B向事务A多次读取的数据中插入了几条数据，事务A再次读取数据是发现多了几条数据，和之前读取的数据不一致



数据库隔离级别

- 未提交读 (read uncommitted)：一个事务在提交前，它的修改对其他事务也是可见的

- 提交读 (read committed)：一个事务提交之后，它的修改才能被其他事务看到

- 可重复读 (repeatable read)：在同一个事务中多次读取到的数据是一致的

- 串行化 (serializable)：需要加锁实现，会强制事务串行执行



| 隔离级别 | 脏读   | 不可重复读 | 幻读   |
| -------- | ------ | ---------- | ------ |
| 未提交读 | 允许   | 允许       | 允许   |
| 提交读   | 不允许 | 允许       | 允许   |
| 可重复读 | 不允许 | 不允许     | 允许   |
| 串行化   | 不允许 | 不允许     | 不允许 |

| Isolation Level      | Dirty Read  | Non-repeatable Read | Phantom Read | Concurrency | Performance |
| -------------------- | ----------- | ------------------- | ------------ | ----------- | ----------- |
| **READ UNCOMMITTED** | ❌ Allowed   | ❌ Allowed           | ❌ Allowed    | Highest     | Best        |
| **READ COMMITTED**   | ✅ Prevented | ❌ Allowed           | ❌ Allowed    | High        | Good        |
| **REPEATABLE READ**  | ✅ Prevented | ✅ Prevented         | ❌ Allowed*   | Medium      | Fair        |
| **SERIALIZABLE**     | ✅ Prevented | ✅ Prevented         | ✅ Prevented  | Lowest      | Worst       |

MySQL的默认隔离级别是可重复读

*MySQL's InnoDB default with REPEATABLE READ also prevents Phantom Reads.*

隔离级别是通过MVCC或者锁来实现

| 隔离级别 | 实现方式                                 |
| -------- | ---------------------------------------- |
| 未提交读 | 总是读取最新的数据，无需加锁             |
| 提交读   | 读取数据时加共享锁，读取数据后释放       |
| 可重复读 | 读取数据时加共享锁，事务结束后释放共享锁 |
| 串行化   | 锁定整个范围的键，一直持有锁直到事结束   |

### 



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

![https](img/https.jpg)

整体流程

1. Client向server发送https请求
2. server向client发送数字证书（公钥，企业信息等）
3. 客户端验证数字证书（用公钥解密并对比hash）
4. server生成会话密钥（symmetric key）
5. client server加密会话





| Layer       | Items                                                        | Unit                                |
| ----------- | ------------------------------------------------------------ | ----------------------------------- |
| Application |                                                              | Message/Data                        |
| Transport   | 端口，网关，端到端                                           | segment段 - TCP；datagram数据报-UDP |
| Network     | IP, 路由器，ARP (Address Resolution Protocol translate from IP to MAC )， ICMP （用于ping） | packet包；数据报datagram            |
| Link        | MAC.  交换器， , LAN， CSMA/CD， MTU, Etherenet, WiFi        | Frame 帧                            |
| Physics     | 集线器hub,中继器repeater                                     | 比特bit                             |

路由器：NAT，访问控制，连接不同网络，寻址

router: 

- Distance Vector 距离矢量  （bellman ford）
  - RIP, BGP,IGRP

- Link State 链路状态 （dijkstra）

  - OSPF，IS-IS

  



选择重传：selective repeat protocols



## 交换变量

```
a = 5, b = 15
a = a + b
b = a - b
a = a - b

and presentation editors
a = a * b
b = a / b
a = a / b

a = a ^ b
b = a ^ b
a = a ^ b

(a ^ a = 0, a ^ 0 = a)

```





## OSI

Open System Interconnection

application layer

presentation layer 表示层：不同终端的上层用户提供数据和信息正确的语法表示变换方法

session layer 会话层：opening, closing and managing a [session](https://en.wikipedia.org/wiki/Session_(computer_science)) between end-user application processes， 认证，权限

transport layer 



## What happens after you request sth

你在浏览器中访问 `http://example.com/index.html`

你的电脑 IP：`192.168.1.10`，源端口：`50000`

服务器 IP：`93.184.216.34`（example.com 真实 IP 之一），目的端口：`80`

MAC：

- 你的网卡 MAC：`AA:BB:CC:DD:EE:FF`
- 网关 MAC：`11:22:33:44:55:66`



**HTTP Payload**

```
GET /index.html HTTP/1.1\r\n
Host: example.com\r\n
User-Agent: curl/8.0\r\n
Accept: */*\r\n
\r\n
```

**TCP Segment: TCP header + [HTTP Payload]**

```
c3 50         # 源端口 0xC350 = 50000
00 50         # 目的端口 0x0050 = 80
00 00 00 01   # 序列号
00 00 00 00   # 确认号
50 18         # Data Offset=5, Flags=PSH+ACK
72 10         # 窗口大小
12 34         # 校验和
00 00         # 紧急指针
```

**IPv4 Packet: IP Header + [TCP Segment]**

```
45 00 00 7A   # 版本4+首部长度5，Total Length=0x007A
1c 46 40 00   # 标识与分片
40 06 a6 ec   # TTL=64, Protocol=6(TCP), Header Checksum
c0 a8 01 0a   # 源IP 192.168.1.10
5d b8 d8 22   # 目的IP 93.184.216.34
```

**WiFi/Ethernet Frame: Frame Header + IPv4 Packet**

```
08 01              # Frame Control: Type=Data, Subtype=Data, ToDS=1
00 00              # Duration
11 22 33 44 55 66  # Addr1: 接入点(AP)MAC — 接收方
aa bb cc dd ee ff  # Addr2: 你的设备MAC — 发送方
11 22 33 44 55 66  # Addr3: BSSID/路由器MAC
10 86              # Sequence Control
AA AA 03 00 00 00  # LLC/SNAP Header
08 00              # EtherType = IPv4
...                # 这里是整个 IPv4 包 (包含TCP+HTTP数据)
f1 e2 d3 c4        # Frame Check Sequence (CRC-32)
```





```
┌─Ethernet Header (Dst MAC | Src MAC | 0x0800)
│   ┌─IPv4 Header (src 192.168.1.10 | dst 93.184.216.34 | proto 6)
│   │   ┌─TCP Header (src port 50000 | dst port 80)
│   │   │   └─HTTP Request Payload ("GET /index.html ...")
│   │   └─TCP Options/Flags
│   └─IP checksum
└─Ethernet FCS
```











## GRPC



gRPC server: grpc server port

gRPB stub: grpc client port



proto buf:数据描述语言

数据构造称为 message



grpc四种服务提供方法

- Unary RPC， 

```protobuf
rpc SayHello(HelloRequest) returns (HelloResponse);
```

- server streaming rpc

客户端向服务器请求，客户端从流中一直读取返回信息

```
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);
```

- client streaming rpc

客户端写入一系列信息到服务器，返回服务器的响应

```
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse);
```

- bidirectional streaming rpc

两个流独立，可以交替读取消息然后写入消息

```protobuf
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse)
```

基于http2

通道 channel，RPC， 消息 message

metadata + channel



protobuf type: 

double, float, int32, int64, uint32, uint64, sint32, sint64, fixed32, fixed64, sfixed32, sfixed64, bool, string, bytes

field number: 1 到536,870,911的值，不应该被重用，低field number的应该占更少空间



protobuf wire format：





## ICMP(ping/traceroute)

ping， traceroute, both use ICMP protocol  (**Internet Control Message Protocol**)



ICMP

当网络层有问题，两个设备无法连接时，可以用ICMP用来诊断

TCP: connection-oriented，通过handshake建立连接，但 ICMP: connectionless protocol，不需要建立连接



**Packet: IP Header + [ICMP datagram]**

 ICMP datagram as follows: 

| 8 bit | 8 bit | 16 bit   | 32 bit          | variable length |
| ----- | ----- | -------- | --------------- | --------------- |
| Type  | Code  | CheckSum | Extended Header | Data/Payload    |



## ML DL

check notion





## Synchronization

- MuTex （Semaphore = 1）
- Semaphore
- SpinLock
- conditional variable
- Read Write Lock

进程

最短时间

最高优先级

**资源分配的基本单位** PCB （PID, 进程描述信息， 分配的内存，IO文件，CPU 寄存器信息） round robin

- 



线程

共享进程的资源

进程的创建和撤销，调度都需要较大的开销，为减小切换开销，使线程作为**调度的基本单位**

引入线程前，进程是资源分配和独立调度的基本单位。引入线程后，**进程是资源分配的基本单位，线程是独立调度的基本单位**。





## TCP

包可能有错误，包可能丢失

在unreliable的基础上，pkt**增加序列号， 增加ack，增加timeout**

每个包有seq_num, ack_num, msg_type (URG, ACK, PSH, RST, SYN, FIN)



**naive approach**

每个pkt都需要一个ack,序列号对应并且如果timeout了则重传，但效率会很低

因此两种优化



**Go-back-N:** 

sender每次可以发送至多N个packet without ack

recevier 如过发现pkg有gap, 则不会ack

sender设置timer给**oldest unacked packek**,如果 超时，重发所有unacked packet



**Selective Repeat **算法

sender每次可以发送至多N个packet without ack

recevier对每个packet独立发送ack

sender设置timer给**each unacked packet**, 如果超时，重发对应的unacked packet



**TCP**算法

滑动窗口， cwnd是

flow control: 流量控制

congestion control: 拥塞控制 - 在slow start, congestion avoidance, fast recovery状态中改变



**三握四挥**

握手：确定发送端和接收端具备收发信息的能力

客户端发送SYN

服务端发送SYN + ACK

客户端收到后，进入ESTABLISHED状态，发送给服务端ACK, 服务端也进入established



挥手: 确保通信双方都能通知对方释放连接

client: FIN, seq = u

server: ACK, ack = u + 1, seq = v

server: FIN + ACK ack = u + 1, seq = w

client: ACK, ack = w + 1, seq = u + 1



## 决策树





## goroutine泄露



## LRU算法

用hash + 双向链表实现





## RNN

RNN对具有序列特性的数据非常有效，它能挖掘数据中的时序信息以及语义信息

https://zhuanlan.zhihu.com/p/123211148





## RAG





