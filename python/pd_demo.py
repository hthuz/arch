
import pandas as pd

def show(msg):
    print(f"==={msg}===")

show("list to Series")
lst = [10,20,30,50,40]
s = pd.Series(lst)
print(s)
print(type(s))

show("indexing series")
print(s[1:3])
print(s > 20)
print(s[s > 20])
print(s.sum())

show("dict to Series. Key will be index")
dic = {"a": 1, "b": 3, "c": 5}
s = pd.Series(dic)
print(s)

dic = {"a": [1,2], "b": [5], "c": [10]}
s = pd.Series(dic)
print(s)
