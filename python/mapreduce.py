
from functools import reduce

def square(x):
    return x * x

def fn(x,y):
    return x * 10 + y

items = [_ for _ in range(5)]

# Map( function, iterable)
mapresult = map(square, items)
print(list(mapresult))


# Reduce(function, [x1,x2,x3,x4]) = f(f(f(x1,x2),x3),x4)
result = reduce(fn, items)
print(result)
