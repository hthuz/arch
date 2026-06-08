# Random

## 实现随机获取某个链表元素

最简单方法，把链表存为数组，返回rand.IntN(n)

另一种方法

```go
func (this *Solution) GetRandom() int {
        res, n := 0, 0

        for cur := this.head; cur != nil; cur = cur.Next {
                n++
                if rand.IntN(n) == 0 {
                        res = cur.Val
                }
        }

        return res
}
```

返回第i个元素的概率, 当前元素为0， 后面元素都不为0

假如4个元素，返回第四个元素，明显是 1/4
返回第二个元素
$$
\frac{1}{2} \times \frac{2}{3} \times \frac{3}{4} 
$$


## 打乱数组

最简单：如果没有重复，对每个位置，可以随机从数组中获取一个元素，然后将其删去

删除可以暴力删 O(n)

或者把随机取到的元素，和最后一个元素交换，删去最后元素 O(1)
```go
func shuffle() {
    n := len(nums)
    shuffled := make([]int, len(nums))
    for i := 0; i < len(shuffled); i++ {
        idx := rand.IntN(len(nums))
        shuffled[i] = nums[idx]

        nums[idx], nums[len(nums)-1] = nums[len(nums-1)], nums[idx]
        nums = nums[:len(nums)-1]
        
    }
}
```

或者Fisher-Yates 洗牌算法，即在上面基础上，实现原地乱序

循环 n 次，在第 i 次循环中（0≤i<n）：
- 在 [i,n) 中随机抽取一个下标 j；
- 将第 i 个元素与第 j 个元素交换。

左边是shuffled过的，右边是待shuffle的

```go
func shuffle() {
    n := len(nums)
    for i := 0; i < len(nums); i++ {
        j := i + rand.IntN(n-i)
        nums[i], nums[j] = nums[j], nums[i]
    }
}
```

