# MySQL

https://cloud.tencent.com/developer/article/1981543



![architectural](https://dev.mysql.com/doc/refman/8.4/en/images/mysql-architecture.png)







## 发送 SELECT 过程

connect + cache + parse + optimize + execute(call API of storage engine)



### connect 部分

长连接：客户端持续有请求，使用同一个连接

短链接：执行完很少几次查询就断开请求

一般有timeout,超过timeout自动断连



建立连接的过程通常是比较复杂的，建议在使用中要尽量减少建立连接的动作，尽量使用长连接。但是全部使用长连接后，有时候 MySQL  占用内存涨得特别快，这是因为 MySQL  在执行过程中临时使用的内存是管理在连接对象里面的。这些资源会在连接断开的时候才释放。所以如果长连接累积下来，可能导致内存占用太大，被系统强行杀掉（OOM），从现象看就是 MySQL 异常重启了。



更多使用**连接池**

### cache （removed in MySQL 8.0)

将最近查询结果缓存，下次查询前将查询语句比对，如果不一样，才回去解析执行，否则直接cache缓存

但也可能导致性能下降，比如一个新查询，除了返回结果，还需要把数据进行缓存，进行缓存本身就需要时间

### parse

### optimize

### execute

通过调用存储引擎的接口实现





## 存储引擎InnoDB



![innodb](https://dev.mysql.com/doc/refman/8.4/en/images/innodb-architecture-8-0.png)

### In-Memory Architecture



### On-Disk Structure

file-per-table tablespaces by default



https://dev.mysql.com/doc/refman/8.4/en/innodb-row-format.html

row format 

- redundant
- compact
- dynamic (by default) （对于溢出的列，只在原始行中存储一个 **20 字节的指针**，**完整数据都存放在溢出页**。这使原始页能存放更多行数据，**全表扫描效率更高**。支持更好的索引键前缀限制）
- compressed



When new records are inserted into an `InnoDB`  [clustered index](https://dev.mysql.com/doc/refman/8.4/en/glossary.html#glos_clustered_index), InnoDB` tries to leave 1/16 of the page free for future insertions and updates of the index records. If index records are inserted in a sequential order (ascending or descending), the resulting index pages are about 15/16 full. If records are inserted in a random order, the pages are from 1/2 to 15/16 full.

https://dev.mysql.com/doc/refman/8.4/en/innodb-physical-structure.html







## 日志

更新数据时，会涉及到日志的更改

redo log： InnoDB storage engine log （物理日志，存储数据被修改的值）

undo log: InnoDB storage engine log (如何undo transaction)

binlog： MySQL server log （逻辑日志，存储逻辑SQL修改语句）



有了 redo log，InnoDB 就可以保证即使数据库发生异常重启，之前提交的记录都不会丢失，InnoDB存储引擎会使用redo log恢复到掉宕机前的时刻，以此来保证数据的完整性。这个能力称为crash-safe。



binlog主要是为了复制（Master-Slave 主从同步），恢复和审计， 自身无法做到crash safe





## MISC



page: 16KB (innodb_page_size)

extent： 1MB

segment：

https://dev.mysql.com/doc/refman/8.4/en/innodb-file-space.html

double write buffer:  store page flushed from buffer pool, 用于recovery

If a row does not exceed the maximum row length, all of it is stored locally within the page. If a row exceeds the maximum row length, variable-length columns are chosen for external off-page storage until the row fits within the maximum row length limit. External off-page storage for variable-length columns differs by row format