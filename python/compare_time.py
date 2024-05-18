
import time
import numpy as np

'''
compare time for evaluating of elemnt is in an iterable
list takes time O(n)
dict, dict.keys(), set takes time O(1) and their time consumption is similar
dict and dict.keys() are similar but it seems that dict.keys() are a bit faster ?
np.array is also fast but not sure about time complexity
if you try to transform a list to a set first. it doesn't reduce time
'''

def timer(func):

    def wrapper(*args, **kwargs):
        times = 3
        ret_val = None
        begin = time.time()
        for _ in range(times):
            ret_val = func(*args, **kwargs)
        end = time.time()
        print(f"func: {func.__name__}, arg:{args[0]}, times: {times}, avg time: {(end - begin) / times}")
        return ret_val

    return wrapper


@timer
def isInList(item, mylist: list):
    return item in mylist

@timer
def isInDict(item, mydict: dict):
    return item in mydict

@timer
def isInDictKey(item, mydict: dict):
    return item in mydict.keys()

@timer
def isInSet(item, myset: set):
    return item in myset
    
@timer
def isInArray(item, myarray):
    return item in myarray

@timer
def isInListToSet(item, mylist: list):
    return item in set(mylist)

@timer 
def isInListToTrasformedSet(item, mylist: list):
    myset = set(mylist)
    return item in myset

size = 10000000

mylist = [i for i in range(size)]
isInList(999, mylist)
isInList(size - 199, mylist)
isInList(10888888,mylist)
isInListToSet(999, mylist)
isInListToSet(size - 199, mylist)
isInListToSet(10888888, mylist)
isInListToTrasformedSet(999, mylist)
isInListToTrasformedSet(size - 199, mylist)
isInListToTrasformedSet(10888888, mylist)
del mylist

mydict = {k:k * 2 for k in range(size)}
isInDict(999, mydict)
isInDict(size - 199, mydict)
isInDict(10888888, mydict)
isInDictKey(999, mydict)
isInDictKey(size - 199, mydict)
isInDictKey(10888888, mydict)
del mydict

myset = set(i for i in range(size))
isInSet(999, myset)
isInSet(size - 199, myset)
isInSet(10888888, myset)
del myset


myarray = np.array(i for i in range(size))
isInArray(999, myarray)
isInArray(size - 199, myarray)
isInArray(10888888, myarray)
del myarray





