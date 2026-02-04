

## IO放大/读写放大

Read Amplification/ Write Amplification

应用层只是需要读取/写入少量数据，但在底层存储中，产生更多的磁盘/网络IO等

写放大：需要写入跟多的数据在磁盘/日志等等

典型例子如leveldb 先写WAL, 再写memtable， 后续compactioin又会多次重写同一数据

对于HDFS block 64MB来说，读写16MB, 不需要读写满64MB, 这个64MB更多是逻辑单位

对磁盘/SSD 的page 4KB来说，

如果读1KB, 通常需要把4KB整个读入内存

如果写1KB, 

- 如果是覆盖写，修改原有文件，则需要先读4KB, 然后内存中修改1KB, 再把4KB写回，
- 如果是追加写，不需要先读旧页，可能在内存里，先把1KB写入到4KB page中，先把空间留着，真正落到磁盘时仍是以page 4KB写出

这种情况下flush: 把1KB的脏页，按照页的单位，真正写到磁盘

对策就是把多次小写合并成大写，降低flush频率





## 小文件问题

数据被且分成大量体积很小的文件，导致存储和计算效率下降

- 造成一个后果是对元数据的更新维护压力更大，使NameNode等负载飙升，成为数据瓶颈

- 另一方面，在mapreduce中，会产生大量的task, 对资源的调度需求，instance启动，等等都会上升，
- 另外，对小文件大量读写，会涉及到大量的随机IO和读写操作，进一步增加了计算的开销

常常可能因为频繁的写入，或者分区过细，使每个分区的数据量都很少导致

对于小文件问题，需要定期进行compaction, 以及避免过细的分区



## MapReduce



输入切分InputSplit => Map => Shuffle => Reduce



1 Input Split: 

按照block大小对输入进行切分，每个split对应一个map task，产出(k1,v1) 作为map输入



2 Map

对每个记录**调用map(k1, v1)**， 输出0-多条结果 (k2,v2)

当数据溢出至磁盘时，通过hash(key) % numReducer 确认这个record属于哪个reducer (partition)， 然后对每个partition内部按照key排序，也会记录下自己的partitionId

这样子reducer在进行拉取数据是，可以通过k-way merge， 一直顺序读，只要key变了，上一个key的数据就结束了，效率更高



3 Shuffle 把同一个key的数据汇聚到同一个reducer

每个reduce 从各个map 节点中抓取自己分区的中间文件片段

reduce拉过来的数据，进行归并，形成一个按key有序排的数据流, 同key的value被聚在一起





4 Reduce

**调用reduce(k2, iterable[v2])**， 输出最终结果 (k3, v3)



5 输出

每个reduce task通常生成一个结果文件



partitioner 决定到那个reducer： hash(key) % numReducers







```markdown
简单说
map：对每条输入记录
map(r) => [(k,v), (k,v)]

shuffle: 
[(k,v), (k,v) ...] => (k,[v1,v2,v3])
并且保证每个key的这一组只交给一个reducer处理

reduce:
reduce(k, [v...]) => [(k', v'), ...]

```

### select count(*) from large_table

```markdown
map: 
输入：所有record
输出 (k=ALL, v=1), (k=ALL, v=1)...

shuffle:
输出 (ALL: [1,1,1,...])

reduce:
输出 (All: totalCount)

```



### select name, count(*) from large_table where xxx group by name

```markdown
map: 
输入：所有record
输出： (alice, 1), (bob, 1), (alice, 1)...

shuffle:
输出：(alice: [1,1,1,1], bob: [1,1])

reduce:
输出: (alice: count) (bob: count)
```

