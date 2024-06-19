



num_elts = 10


def init(num_elts):
    return [i for i in range(1,num_elts + 1)]

parent = init(num_elts)
# Find root of a set
def find(x):
    if parent[x] == x:
        return x
    parent[x] = find(parent[x])
    return parent[x]

def merge(parent,x,y):
    parent[find(x)] = parent[find(y)]



