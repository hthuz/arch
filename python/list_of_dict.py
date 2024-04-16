
list_of_dict = [{}] * 3
for item in list_of_dict:
    print(id(item))
list_of_dict[1]['key'] = 2
print(list_of_dict)

list_of_dict = [{} for _ in range(3)]
for item in list_of_dict:
    print(id(item))
list_of_dict[1]['key'] = 2
print(list_of_dict)
