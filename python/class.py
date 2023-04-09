


class Cookie:
    def __init__(self):
        self.name = None;
        self.size = None;

    def bake(self):
        self.size = 1;
        self.size += 1;

cookie = Cookie()
cookie.bake()

print(cookie.name)
print(cookie.size)
