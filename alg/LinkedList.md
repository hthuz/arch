# Linked List



- Dummy Head to simplify operation
- Double pointers -> (get last N element of a list)
- Fast and slow pointers -> get middle of a linked list (one ptr two steps, one ptr one step, when one finishes, the other in the middle)
- Recursive

- Consider using dummy nodes if head may be removed

  



## Reverse List

```go

// iterative way
prev := nil
cur := head
for cur != nil {
    next = cur.Next
    cur.next = prev
    
    // update
    prev = cur
    cur = next
}
return prev

// post recursive way
func reverse(head) {
    if head == nil || head.Next = nil {
        return head
    }
    
    reversed := reverse(head)
    
    head.Next.Next = head
    head.Next = nil
    return reversed
}

// pre recursive way
func reverse(head, prev) {
    if head == nil {
        return prev
    }
    next := head.next
    head.next = prev
    reverse(next, head)
    
}
reverse(head, nil)
```



## Median of List

```go
func middleNode(head *ListNode) *ListNode {
    first := head
    second := head
    for second != nil {
        second = second.Next
        if second == nil {
            break
        }
        second = second.Next
        first = first.Next
    }
    return first
}
1->2->3->4->5 return 3
1->2->3->4->5->6 return 4
```







