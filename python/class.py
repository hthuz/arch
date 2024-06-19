


class Cookie:
    def __init__(self):
        self.name = None;
        self.size = None;
    def __str__(self):
        return self.name

    def bake(self):
        self.size = 1;
        self.size += 1;

cookie = Cookie()
cookie.name = "HAHAHA"
cookie.bake()
print(cookie)
print(cookie.name)

print(cookie.size)


print(FrozenSet())
