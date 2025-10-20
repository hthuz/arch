# Stack



## min stack

use one stack to record the smallest element 

对于栈来说，如果一个元素 a 在入栈时，栈里有其它的元素 b, c, d，那么无论这个栈在之后经历了什么操作，只要 a 在栈中，b, c, d 就一定在栈中，因为在 a 被弹出之前，b, c, d 不会被弹出。

因此，在操作过程中的任意一个时刻，只要栈顶的元素是 a，那么我们就可以确定栈里面现在的元素一定是 a, b, c, d。

那么，我们可以在每个元素 a 入栈时把当前栈的最小值 m 存储起来。在这之后无论何时，如果栈顶元素是 a，我们就可以直接返回存储的最小值 m。



```go
type MinStack struct {
    stack []int
    minstack []int
}


func Constructor() MinStack {
    return MinStack{
        stack: make([]int, 0),
        minstack: make([]int, 0),
    }
}


func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val)
    if len(this.stack) == 1 {
        this.minstack = append(this.minstack, val)
    } else {
        this.minstack = append(this.minstack, min(this.minstack[len(this.minstack)-1], val))
    }

}


func (this *MinStack) Pop()  {
    if len(this.stack) == 0 {
        return
    }
    this.stack = this.stack[:len(this.stack) - 1]
    this.minstack = this.minstack[:len(this.minstack) - 1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    return this.minstack[len(this.minstack)-1]
}

```



## two stacks to queue

当stack2为空时，stack1记录的是[1,2,3] （后为堆顶）,这个时候把堆一个一个pop然后push到stack2,这个时候顺序就换了

```
type MyQueue struct {
    stack1 []int
    stack2 []int
    
}


func Constructor() MyQueue {
    return MyQueue {
        stack1: make([]int, 0),
        stack2: make([]int, 0),
    }
}


func (this *MyQueue) Push(x int)  {
    this.stack1 = append(this.stack1, x)
}


func (this *MyQueue) Pop() int {
    if len(this.stack2) == 0 {
        for len(this.stack1) != 0 {
            elt := this.stack1[len(this.stack1)-1]
            this.stack1 = this.stack1[:len(this.stack1)-1]
            this.stack2 = append(this.stack2, elt)
        }
    }

    elt := this.stack2[len(this.stack2)-1]
    this.stack2 = this.stack2[:len(this.stack2)-1]
    return elt 

}


func (this *MyQueue) Peek() int {
    if len(this.stack2) == 0 {
        for len(this.stack1) != 0 {
            elt := this.stack1[len(this.stack1)-1]
            this.stack1 = this.stack1[:len(this.stack1)-1]
            this.stack2 = append(this.stack2, elt)
        }
    }

    return this.stack2[len(this.stack2)-1]
    
}


func (this *MyQueue) Empty() bool {
    return len(this.stack1) == 0 && len(this.stack2) == 0
}

```

