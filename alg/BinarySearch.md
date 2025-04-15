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
    }4
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