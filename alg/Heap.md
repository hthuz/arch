# Heap



cf `arch/go/heap/main.go`

[code here](../go/heap/main.go)





A complete binary tree 完全二叉树  （满二叉树）

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





