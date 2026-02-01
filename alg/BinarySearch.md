# Binary Search

## **Basic Template**

```c
int binarySearch(int[] nums, int target ) {
    int left = 0;
    int right = ...;

    while (...) {
        // int mid = left + (right - left) / 2; to avoid overflow
        int mid = (left + right) / 2;
        if (nums[mid] == target) {
            ...
        } else if (nums[mid] < target) {
            left = ...;
        } else if (nums[mid] > target ) {
            right = ...;
        }
    }

    return ...;
}
```

## **left < right or left <= right**

If search range is `[left, right]`, `rigth` will be initialized to `nums.length() - 1`, and you should stop when `left == right + 1.` The case `left == right` is also searched.

```c
int binarySearch(int[] nums, int target) {
    int left = 0;
    int right = nums.lenght() - 1;
    while (left <= right) {
        ...
    }
}
```

## **left = mid or left = mid + 1**

The search range is closed.

If you find that `mid` is not your target, you shouldn't include `mid` in your next search (since the range is closed), so you should avoid `[mid, right]` or `[left, mid]` and use `[mid + 1, right]` or `[left, mid - 1]` instead.

```c
        ...
        if (nums[mid] == target) {
            return mid
        } else if (nums[mid] < target) {
            left = mid + 1;
        } else if (nums[mid] > target) {
            right = mid - 1;
        }
```

## **Left bound search**

Search the leftmost element in the array that matches target. For example, return `1` for `nums = [1,2,2,2,3], target = 2`

```c
int binarySearch(int[] nums, target) {
    int left = 0;
    int right = nums.length(); // Search range is [left, right)

    // stop when left == right
    while (left < right) {
        int mid = (left + right) / 2;
        if (nums[mid] == target) {
            // search towards the left
            right = mid;
        } else if (nums[mid] < target) {
            left = mid + 1;
        } else if (nums[mid] > target) {
            right = mid;
        }
    }

    return left;
    // or more precisely, number of elements smaller than target, left is in range [0, nums.length()]
    // if return -1 on not found
    // return (left == nums.length() || nums[left] != target) ? -1 : left;
}
```

## **Right bound search**

```c
int binarySearch(int[] nums, target) {
    int left = 0;
    int right = nums.length();

    while(left < right) {
        int mid = (left + right) / 2;
        if (nums[mid] == target) {
            // search towards right
            left = mid + 1;
        } else if (nums[mid] < target) {
            left = mid + 1;
        } else if (nums[mid] > target) {
            right = mid;
        }
    }
    return left - 1; // since update on left is always mid + 1
    // if return -1 on not found
    // return (left == 0 || nums[left - 1] != target) ? -1 : left - 1;
}
```

2,5,6,0,0,1,2

0,1,2,3,4,5,6

0,0,1,2,2,5,6, k = 4

k前： +(l - k)

k后：- k

2,2,2,2,5,6,0,0,1,2 k=4

0,0,1,2,2,2,2,2,5,6

0,1,2,3,4,5,6,7,8,9







如果是左闭右开的搜索， [left, right)，那么请记住每次搜索的范围都是 [left, right)

对应的

```go

left := 0
right := len(nums) // 如果right = len(nums) - 1, 则会在len(nums) - 1 中漏过搜索
for left < right {
    
    if nums[mid] < target {
        left = mid + 1 // 如果还是 left = mid, 下次又会把mid包含在搜索里，但是这次已经把mid搜索过了
    }
    if nums[mid] > target {
        right = mid // 如果right = mid - 1, 因为右边界是不搜索的，因此会把mid - 1的情况漏了
    }
}
```

如果是两边都闭的搜索 [left, riight], 那么请记住每次搜索的范围都是[left, right]

```go
left := -
right := len(nums) - 1 // 如果right = len(nums), 会超界限

for left <= right {
    if nums[mid] < target {
        left = mid + 1 // mid不需要再被下一次在搜索了
    }
    if nums[mid] > target {
        right = mid - 1 // mid不需要在被下一次再搜索了
    }
}
```

## 通用模板

```go
找到第一个大于等于target的位置， 和寻找数组中插入位置相同
// 第一个 nums[mid] >= target, nums[mid] < target时，应该被排除
left := 0
right := len(nums)
for left < right {
    if nums[mid] < target { 
        left = mid + 1
    } else {
        right = mid
    }
    // 当nums[mid] = target时, 有可能你是后面的大于等于target的，因此对right进行左压
}

return left

```



### 变式一：找到数组中第一个等于target的位置

```go

逻辑和模板相同，但是返回时需要检查
if left < len(nums) && nums[left] == target {
    return left
}
return -1 // 不存在
```



### 变式二：找到数组中第一个大于target的位置

则需要找到的是第一个mid，满足nums[mid] > target, 当nums[mid] <= target 时，mid应该被排除

```go
for left < right {
    if nums[mid] <= target {
        left = mid + 1
    } else {
        right = mid
    }
}
return left
```



### 变式三：找到数组中最后一个等于target的位置

```go
for left < right {
    if nums[mid] <= target { // 需要继续找，直到右边界，
        					 // 本质上相当于此时left是第一个大于target的位置，然后往回走一步就是最后一个小于/等于的位置
        left = mid + 1
    } else {
        right = mid
    }
}

left = left - 1
if left >= 0 && nums[left] == target {
    return left
}
return -1 // 不存在
```



