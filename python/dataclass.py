

from dataclasses import dataclass, field

@dataclass
class User:

    category: str = field(default="amazing")

    name: str = field(default="unknown")
    # items: list = field(init=False, default_factory=list)
    items: list = field(default_factory=list)


user1 = User("alice")
user2 = User("bob")
user3 = User()

user1.items.append(3)

# user2.category = "teacher"
# User.category = "teacher"
# user1.category = "student"
User.name = "fu"
print(user1.category)
print(user2.category)

print(user1)
print(user2)
print(user3)


# class Person:
#     def __init__(self, name, age):
#         self.name = name
#         self.age = age
#     def __repr__(self):
#         return f"{self.age}, {self.name}"
#
#     
#     @classmethod
#     def from_birth_year(cls, name, birth_year):
#         """通过出生年份创建 Person 实例"""
#         from datetime import datetime
#         age = datetime.now().year - birth_year
#         return cls(name, age)  # 相当于调用 Person(name, age)
#     
#     @classmethod
#     def from_dict(cls, data_dict):
#         """通过字典创建 Person 实例"""
#         return cls(data_dict['name'], data_dict['age'])
#
# # 使用不同的构造方式
# person1 = Person("Alice", 25)  # 标准构造
# person2 = Person.from_birth_year("Bob", 1990)  # 类方法构造
# person3 = Person.from_dict({"name": "Charlie", "age": 30})  # 类方法构造
#
#
# print(person1)
# print(person2)
# print(person3)
