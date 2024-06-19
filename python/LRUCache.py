class LRUCache:

    def __init__(self, capacity: int):
        self.cache = {} # key: value
        self.lru = [] # a stack record lru key, right is top of stack, most recently used
        self.num = 0
        self.capacity = capacity


    def get(self, key: int) -> int:
        # key exists
        try:
            value = self.cache[key]
            self.lru.remove(key)
            self.lru.append(key)
            print("Failure")
            return value
        except:
            print("Success")
            return -1

    def put(self, key: int, value: int) -> None:
        if self.num < self.capacity:
            self.cache[key] = value
            self.lru.append(key)
            self.num += 1
            return 

        # key exists
        try:
            self.cache[key] = value
            self.lru.remove(key)
            self.lru.append(key)
        # eviction
        except:
            lru_key = self.lru.pop(0)
            del self.cache[lru_key]
            self.cache[key] = value
            self.lru.append(key)

            

if __name__ == "__main__":
    lrucache = LRUCache(2)
    lrucache.put(1,1)
    lrucache.put(2,2)
    print(lrucache.cache)
    print(lrucache.lru)
    print(lrucache.get(1))

