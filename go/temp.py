

nums = []
a = 0
b = 0
# type 1
for num in nums:
    a += num.profit
    b += num.cost

# type 2
for num in nums:
    a += num.profit
for num in nums:
    b += num.cost

# type 3
a = sum(num.profit for num in nums)
b = sum(num.cost ** 2 for num in nums)

