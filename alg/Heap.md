# Heap



cf `arch/go/heap/main.go`

[code here](../go/heap/minheap.go)





A complete binary tree 完全二叉树  （满二叉树）

每个节点都大于等于其节点的值

每个节点都小于等于其节点的值







for complete binary tree, node at index `m`, then its children at index `2m+1`, `2m+2`

For children `m`, its parent is at index `(m-1) / 2`



For more details, refer to `arch/go/heap/main.go`



root is usually referred to as level 0



min heap: for any node, its value is `<= `its children, elements may be the same in value

根始终最小，因此堆的array表示中，第一个数始终最小

![img](https://pic3.zhimg.com/v2-599ec8060b240a07e412972996617324_1440w.jpg)





### Insert



![img](https://pica.zhimg.com/v2-06486db4e3eae115581a0ec8917e693e_1440w.jpg)





### Delete

![heap_delete](/home/autentico/arch/alg/img/heap_delete.jpg)





## Build Heap

从数组中构建堆，一种方法是对每个元素进行insert操作，每次insert复杂度为`O(logn)`因此总复杂度为`O(nlogn)`， 如果数组的元素不是一次性全知道可以采用这种方式



但对于已经知道所有元素的数组，可以从（最后一个非叶子节点）`n/2-1`到`0`进行逐一sift down操作，复杂度`O(n)`



## Go

使用container/heap

```go

type Heap []int

// 这里heap本身就是数组，swap虽然对元素进行修改，但其指针不需要改变
func(h Heap) Less(i, j int) bool {
    return h[i] < h[j]
}
func(h Heap) Len() int {
    return len(h)
}

func (h Heap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

// 这里pop和push, 需要对len进行更改，所以需要使用*heap
func (h *Heap) Pop() any {
	tail := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tail
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func main() {
    h := &Heap{}
    heap.Init(h)
    
    heap.Push(h, 4)
    a1 := heap.Pop()
}

```



自己实现

```go
// min heap
type Heap struct {
	nums []int
}


// 第一种写法
func (h *Heap) Push(x int) {
	h.nums = append(h.nums, x)

	i := len(h.nums) - 1
	for i > 0 {
		p := (i - 1) / 2
		if h.nums[i] < h.nums[p] {
			h.nums[i], h.nums[p] = h.nums[p], h.nums[i]
			i = p
		} else {
			break
		}
	}
}
// 第二种写法
func (h *Heap) Push(x int) {
    h.nums = append(h.nums, x)
    i := len(h.nums)-1
    for i > 0 {
        p = (i - 1) / 2
        if h.nums[i] <= h.p[i] {
            break
        }
        h.nums[i], h.nums[p] = h.nums[p], h.nums[i]
        i = p
    }
}
// 第三种写法
func (h *Heap) Push(x int) {
    h.nums = append(h.nums, x)
    for i := len(h.nums) - 1; i > 0 {
        p := (i - 1) / 2
        if h.nums[i] <= h.nums[p] {
            break
        }
        h.nums[i], h.nums[p] = h.nums[p], h.nums[i]
        i = p
    }
}
// 第四种写法
func (h *Heap) Push(x int) {
    h.nums = append(h.nums, x)
    for i := len(nums) - 1; i > 0 && h.nums[i] < h.nums[(i-1)/2]; i = (i-1)/2 {
        h.nums[i], h.nums[(i-1) / 2] = h.nums[(i-1)/2], h.nums[i]
    }
}

func (h *Heap) Pop() int {
	top := h.nums[0]
	h.nums[0] = h.nums[len(h.nums)-1]
	h.nums = h.nums[:len(h.nums)-1]

	i := 0
	for {
		c1 := 2*i + 1
		c2 := 2*i + 2
		smaller := i
		if c1 < len(h.nums) && h.nums[c1] < h.nums[smaller] {
			smaller = c1
		}
		if c2 < len(h.nums) && h.nums[c2] < h.nums[smaller] {
			smaller = c2
		}
		if smaller == i {
			break
		}
		h.nums[i], h.nums[smaller] = h.nums[smaller], h.nums[i]
		i = smaller
	}
	return top
}

```





# Priority Queue

insert()

deleteMin()

getMin()
