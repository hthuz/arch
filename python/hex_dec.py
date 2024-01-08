
def hex2dec_signed(x: str,num_bits=16):
    ans = int(x,16)
    if ans > (1 << (num_bits- 1)):
        ans -= (1 << num_bits)
    print(ans)

def hex2dec_unsigned(x: str):
    print(int(x,16))


hex2dec_signed("fffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8e9dc3",256)




