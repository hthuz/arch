# Number


$$
123 \% 10 = 3  
$$

$$
123 / 10 = 12
$$

$$
-123 \% 10 = -3
$$

$$
-123 / 10 = -12
$$



```go
var maxInt32 int32 = 1 << 31 - 1
var minInt32 int32 = -1 << 31
```

check if some number will overflow: check if larger than `maxInt32` or smaller than `minInt32` before updating the value

