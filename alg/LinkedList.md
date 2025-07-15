# Linked List



- Dummy Head to simplify operation
- Double pointers -> (get last N element of a list)
- Fast and slow pointers -> get middle of a linked list (one ptr two steps, one ptr one step, when one finishes, the other in the middle)
- Recursive

- Consider using dummy nodes if head may be removed

  

## removal of an elemnt


For example

dummy->a->b->c

If index is 1, 
iterate once, with cur being a, changing it to a->c, thereby removing b

```go 
func remove(index int, head *ListNode) {
    dummy := &ListNode{
        Val: 0,
        Next: head,
    }
    cur := dummy
    for range index {
        cur := cur.Next
    }

    cur.Next = cur.Next.Next
    return dummy.Next
}
```


## add of an elemnt

dummy->a->b->c
If index is 0


```go
func add(index int, head *ListNode, val int) {
    dummy := &ListNode {
        Val: 0,
        Next: head,
    }
    cur := dummy
    for range index {
        cur := cur.Next
    }

    node := &ListNode {
        Val: val,
        Next: cur.Next
    }
    cur.Next = node
    return dummy.Next
}
```



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

```cpp
ListNode* getMedian(ListNode* left, ListNode* right) {
    ListNode* fast = left;
    ListNode* slow = left;
    while (fast != right && fast->next != right) {
        fast = fast->next;
        fast = fast->next;
        slow = slow->next;
    }
    return slow;
}

```






