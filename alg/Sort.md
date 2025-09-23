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



