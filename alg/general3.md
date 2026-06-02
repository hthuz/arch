
# Vector Database


输入某个input embedding后的vector， 和数据库中的对应vector field比较，返回score最高的

每次搜寻，指定index， metric, vector field

## 用于计算 similarity 的 metrics

Euclidean Distance (L2) = $||a-b||$ = $\sqrt{(ai-bi)^2}$

内积 Inner Product (IP) = $ab$ = $a1b1 + a2b2 +.. anbn$ = $||a||||b||cos\theta$
s
如果向量已经归一化，和cosine similarity 相同
$cos\theta=\frac{ab}{||a||||b||}$



## Recall

$$
Recall@k = \frac{|ANN\ TopK \cap True\ TopK|}{K}
$$

例如：True Top100, 有 100 个结果。

ANN Top100：找回了 97 个

那么：Recall@100 = 97%

## Index

Data structure + Quantization + Refiner

KNN：FLAT
ANN（Approximate Nearest Neighbour）的实现：IVF，HNSW，PQ等

### FLAT
最简单的，暴力搜索，query vector和库中每个vector比较，因此100% recall

### IVF_FLAT

Inverted File: 数据库中的向量通过k-means聚类，分成几个region， 每个region由centroid代表. 
index: map centroid => vectors in that cluster

Flat: 每个cluster的vector没有任何修改

对于一个query vector，首先和各个centroid对比，找到最近的，然后和centroid对应region的vector对比，找到最相似的。


![workflow](https://milvus-docs.s3.us-west-2.amazonaws.com/assets/IVF-FLAT-workflow.png)

但是如果找到的cluster，和目标vector不在同一个region，可能无法成功召回
![drawback](https://milvus-docs.s3.us-west-2.amazonaws.com/assets/IVF-FLAT-workflow-2.png)

### IVF_SQ8
Inverted File + Scalar Quantization 

SQ8: 32 bit float => 8-bit integer （0-255）

首先通过$normalized = \frac{val-min}{max-min}$归一化至0-1返回内，然后成255并四舍五入

### HNSW

Hierarchical Navigable Small World