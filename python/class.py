


class Cookie:
    def __init__(self, name, size):
        self.name = name;
        self.size = size;

    def bake(self):
        self.size += 1;

cookie = Cookie("Chocolate",4)
cookie.bake()

print(cookie.name)
print(cookie.size)
