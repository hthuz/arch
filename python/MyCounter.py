

class MyCounter(dict):

    def __init__(self):
        return

    def add(self, key):
        if key not in self.keys():
            self[key] = 0
        self[key] += 1

    def most_common(self):

counter = MyCounter()


counter.add("a")
counter.add("b")
counter.add("a")
print(counter)
print(counter)

