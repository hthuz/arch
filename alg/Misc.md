# Misc

Fastest way to know if two arrays contain the same element => using hash, $O(m+n)$

Fastest way to know if an array contain repetitive elements => using hash, $O(n)$



## Traverse directions

a more convenient to go up, down, left, right, four directions in 2-D array

```go 
type pair struct {r, c int}

var directions = []pair{{-1,0}, {1,0}, {0,-1}, {0,1}}

func traverse(row, col int) {
    for _, d := directions {
        newR, newC := row + d.r, col + d.c; 
        if 0 <= newR && newR < height && 0 <= newC && newC < width {
            // do something
        }
    }
}


```

## Deletion of an element from an array

By recording the updated length

```go
 
func del(nums []int, index int, length *int) {
    for i := index; i < length - 1; i++ {
        nums[i] = nums[i+1]
    }
    (*length)--
}
```



## isPrime

判断一个数是否为质数

```go

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if (x/i)*i == x { /这里也可以等价为 if x % i == 0
			return false
		}
	}
	return true
}
```

因为(x/i)*i == x实际上是判断floor(x/i) == x/i

floor(x/i) = x/i 等价于x % i == -



## 判断数组是否有重复

判断数组是否有重复，并且比如说找到重复的元素，无特殊情况下最优方法就是用哈希集合，对所有元素遍历一边并确认元素是否已出现过，时间空间复杂度均为O(n)



但是一些特定情况下，可以优化，比如lc287





## 多个排序数组中得到最大的top N个

类似于设计推特的getnews feed

Q1: 给定K个已按升序排列的数组，请找出所有元素中最大的 N 个，并以降序形式返回这 N 个元素构成的数组。

最优思路：使用一个堆，从每个K个数组中均获取最后一个元素`num[k][-1]`，然后对heap进行pop, 同时将pop掉元素所在数组的次最大个插入heap中，直到有N个元素或者堆空了

```go
for len(heap) > 0 && len(res) < N {
    elt = heap.pop()
    res append elt
    
    num = nums[elt.k]
    
    if elt.idx > 0 {
        heap.push(num[elt.idx-1])
    }
}
```



Q2: 给定 K 个数组（可能未排序），请找出所有元素中最大的 N 个，并以降序（或升序）形式返回这 N 个元素构成的数组。



## 字典序排序 lexicographical order

相关题目：lc 31 下一个排序， lc 386  字典序排数， lc 2697 字典序最小回文串

从左到右逐位比较元素大小，若长度不同短者优先

比如1-12字典序: 1,10,11,12,2,3,4,5,6,7,8,9  



如lc 386 返回1-n数字的字典序排序

```go
func lexicalOrder(n int) []int {

	res := make([]int, n)
	num := 1
	for i := 0; i < n; i++ {
		res[i] = num
		if num*10 <= n {
			num = num * 10
		} else {
			for num%10 == 9 || num == n {
				num = num / 10
			}
			num += 1
		}
	}
	return res
}
```

更容易理解，也更抽象更形象的一种思维是递归dfs

```go
func lexicalOrder(n int) []int {
	res := make([]int, 0)

	var dfs func(int)

	dfs = func(start int) {
		if start > n {
			return
		}
		res = append(res, start)
		for i := range 10 {
			dfs(10*start + i)
		}

	}

	for i := 1; i < 10; i++ {
		dfs(i)
	}
	return res
}
```



## n-sum 和二分

两个虽然都有left, right但是需要注意的是，n-sum问题中，找到几个元素，和为n, 中left, right分别是left++, right--这样子变动的，本质上是左右指针，而二分是跳动的，两个不要混淆

两数之和的另一种解法：用hash记录每一个num, 然后遍历，寻找hash map中是否存在元素target-num，同时需要注意先判断hash是否存在，然后更新hash, 否则nums = [2], target = 4的情况也会被判对



## 两个数组的交集

最直观的方法，两个哈希，O(m + n)

```go
func intersection(nums1 []int, nums2 []int) (intersection []int) {
    set1 := map[int]struct{}{}
    for _, v := range nums1 {
        set1[v] = struct{}{}
    }
    set2 := map[int]struct{}{}
    for _, v := range nums2 {
        set2[v] = struct{}{}
    }
    if len(set1) > len(set2) {
        set1, set2 = set2, set1
    }
    for v := range set1 {
        if _, has := set2[v]; has {
            intersection = append(intersection, v)
        }
    }
    return
}


```

另一种方法，根据排序， O(mlog(m) + nlog(n))， 和merge两个数组类似

```go
func intersection(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)

    i := 0
    j := 0
    res := make([]int, 0)
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] == nums2[j] {
        	// 注意这里需要对res进行去重，因为nums1[i-1] 和 nums[i]可能相同
            if len(res) == 0 || res[len(res)-1] != nums1[i] {
                res = append(res, nums1[i])
            }

            i += 1
            j += 1
        } else if nums1[i] < nums2[j] {
            i += 1
        } else {
            j += 1
        }
    }
    return res
}
```



## 合并两个有序数组

```go
func merge(nums1 []int, m int, nums2 []int, n int) {
    sorted := make([]int, 0, m+n)
    p1, p2 := 0, 0
    for {
        if p1 == m {
            sorted = append(sorted, nums2[p2:]...)
            break
        }
        if p2 == n {
            sorted = append(sorted, nums1[p1:]...)
            break
        }
        if nums1[p1] < nums2[p2] {
            sorted = append(sorted, nums1[p1])
            p1++
        } else {
            sorted = append(sorted, nums2[p2])
            p2++
        }
    }
 
    
func merge(nums1 []int, m int, nums2 []int, n int) {
    sorted := make([]int, m + n)
    i, j := 0, 0
    
    for k := 0; k < m + n; k++ {
        if i == m {
            sorted[k] = nums2[j]
            j += 1
            continue
        }
        if j == n {
            sorted[k] = nums1[i]
            i += 1
            continue
        }
        if nums1[i] < nums2[j] {
            sorted[k] = nums1[i]
            i += 1
        } else {
            sorted[k] = nums2[j]
            j += 1
        }
    }

    copy(nums1, sorted)
}
    
    
// 更高级的解法，从后往前合并，这样不需要一个中间变量
    
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	for k := m + n - 1; k >= 0; k-- {
		if j == -1 || (i >= 0 && nums1[i] >= nums2[j]) {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
    
func merge(nums1 []int, m int, nums2 []int, n int) {
    i := m-1
    j := n-1
    // nums2结束的时候直接结束
    for j >= 0 {
        if i >= 0 && nums1[i] >= nums2[j] {
            nums1[k] = nums1[i]
            i--
        } else {
            nums1[k] = nums2[j]
        }
        k--
    }
}
    

```





## k bit 遍历



```go
func main() {
	for k, v := 0, 1; k <= 31; k, v = k+1, v<<1 {
		fmt.Println(k, v)
	}

	for k := 0; k <= 31; k++ {
		fmt.Println(k, 1<<k)
	}
}

```





## 前缀和

永远定义pre[0] = 0

pre[i] 表示前i个元素的和, 不包括nums[i]

```go
pre := make([]int, len(nums)+1)
for i := 0; i < len(nums); i++ {
    pre[i+1] = pre[i] + nums[i]
}

```

=> nums[i] = pre[i+1] - pre[i]

```go
for i := 0; i < len(nums); i++ {
    nums[i] = pre[i+1] - pre[i]
}

```

=> sum(l,r)  = pre[r+1] - pre[l]，l, r 是inclusive



## 给定数组，是否有n个/2个元素，满足某种性质

常用解法：hash, 排序+双指针，滑动窗口

比如，两个数的和为target, 使用hash, 可以将复杂度下降至O(n)

```go
for _, num := range nums {
    if set[target-num] {
        return true
    }
    set[num] = true
}
```

n个数的差为4的倍数 => 取模4相同

```go
for _, num := range nums {
    cnt[num % 4]++
    if cnt[num % 4] >= n {
        return true
    }
}
```

两个数的绝对值差小于k，如果用hash复杂度仍未O(n^2)

```go
for _, num := range nums {
    for _, v := range set {
        if abs(num - v) < k {
            return true
        }
    }
}
```

更好的方式是排序O(nlogn)

```go
sort.Ints(nums)
for i := 1; i < len(nums); i++ {
    if nums[i] - nums[i-1] < k {
        return true
    }
}
```

其他的hash, 比如一个数位另一个数的n倍，乘积为target，异或值为xxx等等

