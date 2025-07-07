# GraphSearch



Basic

```python

u: node
Explore(G, u):
	visited[u] = True
	toExplore.push(u)
	while (toExplore is not empty):
        x <- toExplore.pop()
        for y in Adj(x):
            if !visited[y]:
                visited[y] = True
                toExplore.push(y)
                
# toExplore stack -> DFS
# toExplore queue -> BFD

# Recursive way

Explore(G,u):
    visited[u] = True
   	for y in Adj(u):
        if !visited[y]:
            visited[y] = True
            Explore(G,y)
	
```

