
# Union Find

https://zhuanlan.zhihu.com/p/93647900

also called disjoint set (dividing elements into subsets that have no common elements)

Elements are unique in union find

In principle, a forest
e.g.
1<-2<-3<-4
7<-5<-3


Supports two operation
- Union: combine two sets into one
- Find: judge if two elements are in the same set (return root of a elemnt)



## Find Operation
```go
type UnionFind[T comparable] struct {
    parent map[T]T
    rank  map[T]int
}

// return the parent of a element
func (uf *UnionFind) Find(x T) {
    if uf.parent[x] == x {
        return x
    } else {
        return uf.Find(uf.parent[x])
    }
}

// a more efficient way is path compression when finding
// when find the root, the it as direct parent
func (uf *UnionFind) Find(x T) {
    if uf.parent[x] == x {
        return uf.parent[x]
    } else {
        uf.parent[x] = uf.Find(uf.parent[x])
        return uf.parent[x]
    }
}
// or in a more simpler form
func (uf *UnionFind) Find(x T) {
    if uf.parent[x] != x {
        uf.parent[x] = uf.Find(uf.parent[x])
    }
    return uf.parent[x]
}



```

## Union Operation
```go
func (uf *UnionFind) Union(x, y T) {
    uf.parent[x] = uf.Find(y)
}
```

define rank[x] as (upper bound) of height of subtree with x as root
// in practice, you should combine the simpler tree into the more complex tree
// for example. combine a large tree into a single node increases the path to root

```go
func (uf *UnionFind) Union(x, y T) {
    root_x := uf.Find(x)
    root_y := uf.Find(y)
    if uf.rank[root_x] < uf.rank[root_y] {
        uf.parent[root_x] = root_y
    } else if uf.rank[root_x] > uf.rank[root_y]{ 
        uf.parent[root_y] = root_x
    } else {
        // the rank of new root increases by one
        uf.parent[root_y] = root_x
        uf.rank[root_x]++
    }
}

```


## Other extended operations, Merge two union find
```go
func (uf *UnionFind) Merge(other *UnionFind) {
    for x := range other.parent {
        rootX := other.Find(x)
        uf.Add(x)
        uf.Add(rootX)
        uf.Union(x, rootX)
    }
}

```



