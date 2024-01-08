
# defi ranger
Ry = 1000
Rx = 1000
# uniswap v1 price formula
def sell_x_for_y(x):
    global Ry
    global Rx
    res = 997 * x * Ry / (997 * x + 1000 * Rx)
    # res = x * Ry / (x + Rx)
    Ry -= res
    Rx += x
    return res

def sell_y_for_x(y):
    global Ry
    global Rx
    res = 997 * y * Rx / (997 * y + 1000 * Ry)
    # res = y * Rx / (y + Ry)
    Ry += y
    Rx -= res
    return res


y_i_have = sell_x_for_y(900)
print(y_i_have)

print(sell_y_for_x(y_i_have))

