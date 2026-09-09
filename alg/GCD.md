# 算法



## GCD

Great Common Divisor 最大公约数

gcd(12, 24) = 12

gcd(12, 18) = 6



求gcd(a,b)



欧几里德算法

```
gcd(a,b) = gcd(b, a % b)

  gcd(12, 24) 
= gcd(24, 12)
= gcd(12, 0)
= 12
```

```go
func gcd(a,b) {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}
```



a 和 b 的公约数集合，和b, a % b 的公约数集合相同





## LCM

Least Common Multiple 最小公倍数

LCM(4,6) = 12

