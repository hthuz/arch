
import re
pattern = r'"[^"]*"'
item = 'Amazing: "Hello" world'
res = re.findall(pattern, item)
print(res)
