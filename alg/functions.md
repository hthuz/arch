# Common used functions in Go



```go

func main() {
    i, err := strconv.Atoi("-42")
    s := strconv.Itoa(-42)
    
    strings.Contains("abcedfg", "a")
    strings.Split("who,are,you", ",")
    
    math.MaxInt32
    math.MaxInt64
    math.MinInt32
    math.MinInt64
    
    slices.Sort(arr)
    names := []string{"Bob", "alice", "VERA"}
    slices.SortFunc(names, func(a, b string) int {
        return cmp.Compare(strings.ToLower(a), strings.ToLower(b))
    })
    
    
    sort.Ints(nums)
    sort.Strings([]string{})
   	// 自定义slice排序
    sort.Slice(s, func(i, j int) bool {
        return s[i] < s[j]
    })
    sort.Slice(s, func(i,j int) bool {
        return people[i].age < people[j].age
    })


    // rand
    rand.IntN(100) // [0,100)
    rand.Float64() // [0,1)

    nums := []int{1, 2, 3, 4, 5}
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
}	
```

