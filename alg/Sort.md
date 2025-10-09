​	

# Sort



## bucket sort

divide and conquer

把数组分到m个桶里，每个桶再进行排序，然后将各个桶排序的结果合并起来

two keys：

1. 桶的数目可以尽可能多
2. 尽量将数组均匀分配到m个桶中



```go

func BucketSort(nums []int) []int {
    numBuckets := 5
    buckets := make([][]int, numBuckets)
    interval := max(nums) / numBuckets
    if interval := 0 {
        interval = 1
    }

    
    for _, num := range nums {
        idx = num / interval
        // num is 100
        if idx >= numBuckets {
            idx = numBuckets - 1
        }
        buckets[idx] = append(buckets[idx], num)
    }
    
    for _, bucket := range buckets {
        sort.Ints(bucket)
    }
    
    result := make([]int, 0, len(nums))
    for _, bucket := range buckets {
        result = append(result, bucket...)
    }
    return result
}


```

如果刚好每个桶分布只有一两个，那么复杂度就接近于O(n + k) (k为桶的数目)

如果全部都在一个桶中，复杂度为O(nlogn)， 但一般使用插入排序，这种情况下复杂度为O(n^2)



一个比较通用的映射关系
$$
bucket\ idx = \lfloor{\frac{num - min}{max-min+1}*numBucket}\rfloor
$$

## radix sort 

基数排序

LSD least significant digit

通过逐位对数字进行排序



```go
func RadixSort(nums []int ) []int {
    if len(nums)  <= 1 {
        return nums
    }
    maxVal := max(nums)
    
    // LSD，低位优先
    for exp := 1; maxVal / exp > 10; exp *= 10 {
        CountingSort(nums, exp)
    }
    return nums
}

// 每次对第几位数进行排序，稳定排序，一般使用计数排序或桶排序
func CountingSort(nums []int, exp int) {
    
}
```

从**低位到高位**依次进行**稳定排序**，
 这样在处理更高位时，不会破坏低位已经排好的相对顺序。

假设x = d2d1d0,

先对d0排序，再对d1排序，再对d2排序





## counting sort

非比较排序算法

计数排序

线性复杂度，将数组转化为额外开辟的数组空间中



- 确认最大值和最小值后，

- 创建一个max-min+ 1的counter数组，

- count[i] 记录i + min出现的次数

- 再根据count，按顺序返回一个有序数组

`n`: number of elements

`k`: max - min

时间复杂度： O(n + k), 空间复杂度 O(k)

有限范围，k应该小，才比较适用



```go

func CountSort(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }
    
    minVal := nums[0]
    maxVal := nums[0]
    for _, num := range nums {
        minVal = min(minVal, num)
        maxVal = max(maxVal, num)
    }
    counter := make([]int, maxVal - minVal + 1)
    for _, num := range nums {
        counter[num - minVal] += 1
    }
    
    res := make([]int, 0, len(nums))
    for idx, count := range counter {
        for range count {
            res = append(res, idx + minVal)
        }
    }
    return res
}
```



以上方法可以实现排序，但无法做到**稳定排序**

**稳定排序**：如果数组中有多个值相等的元素，**排序后这些相等元素的相对先后顺序与排序前保持一致**

如25 Alice, 25 Bob 按年龄排序后仍保持25 Alice, 25 Bob 而不是 25 Bob, 25 Alice



To have stable sort, use the following



```go
// nums = [1, 5, 1, 1, 9, 2, 3, 5, 1]
// nums = [1 1 1 1 2 3 5 5 9]
func CountSort(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }
    
    minVal := nums[0]
    maxVal := nums[0]
    for _, num := range nums {
        minVal = min(minVal, num)
        maxVal = max(maxVal, num)
    }
    counter := make([]int, maxVal - minVal + 1)
    for _, num := range nums {
        counter[num - minVal] += 1
    }
    
    // prefix sum
    for i := 1; i < len(counter); i++ {
        counter[i] += counter[i-1]
    }
    // counter = [4,1,1,0,2,0,0,0,1]
    // prefixs = [4,5,6,6,8,8,8,8,9]
    // prefixs[i]: 值i最后一个元素应放到的位置
    
    res := make([]int, len(nums))
    for i := len(nums)-1; i >= 0 ;i-- {
        index := nums[i] - minVal
        counter[index] -= 1
        res[counter[index]] = nums[i]
    }

}
```

计数排序是桶排序在“桶宽 = 1 且值域离散已知”条件下的特化版本



## insertion sort

扑克牌排序





从第二个元素开始，依次将当前元素插入到已排好序的左侧序列中。

通过不断向左比较并移动元素，直到找到合适位置。

```go
func InsertionSort(nums []int) {
    for i := 1; i < len(nums); i++ {
        key := nums[i]
        j := i - 1
        // 向左找到合适位置
        for j >= 0 && nums[j] > key {
            nums[j+1] = nums[j]
            j--
        }
        // there's one additional minus so you need to add here
        nums[j+1] = key
    }
}



func InsertionSort(nums []int ) []int {
    for i := 1; i < len(nums); i++ {
        val := nums[i]
        j := i
        for j >= 1 && nums[j-1] > val {
            nums[j] = nums[j-1]
            j--
        }
        nums[j] = val
    }
    return nums
}


```







## shell Sort

希尔排序

插入排序在元素距离最终位置很远时，需要一次次移动，效率低。

希尔排序通过**“分组 + 逐步缩小间隔”**，让元素先进行**远距离的交换**，使数组整体更接近有序，然后在最后一次（gap=1）时插入排序的工作量就很小。

```go
func ShellSort(nums []int) {
    n := len(nums)
    // 1. 选择初始 gap
    for gap := n / 2; gap > 0; gap /= 2 {
        // 2. 对每个子序列做插入排序
        for i := gap; i < n; i++ {
            temp := nums[i]
            j := i
            // 按 gap 进行“分组插入排序”
            for j >= gap && nums[j-gap] > temp {
                nums[j] = nums[j-gap]
                j -= gap
            }
            nums[j] = temp
        }
    }
}
// [23, 12, 1, 8, 34, 54, 2, 3]
// gap = 4, [23,34], [12,54], [1,2], [8,3]
// result = [23,12,1,3,34,54,2,8]
// gap = 2, [23,1,34,2], [12,3,54,8]
// result = [1,3,2,8,23,12,34,54]
// gap = 1, [1,3,2,8,23,12,34,54]
// result = [1,2,3,8,12,23,34,54]
```

平均情况
$$
O(n^{1.5})
$$
最坏情况 O(n^2)

