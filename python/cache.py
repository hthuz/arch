
# Given parameters of a cache, compute its (data) size
s_offset = 5
s_index = 4
s_ways = 2

line_size = 2**s_offset * 8 # in bits
num_index = 2**s_index
num_ways = 2**s_ways
size = line_size * num_ways * num_index
print(size,"bits")
print(size / 8, "bytes")
print(size / 8 / 1024, "KB")
