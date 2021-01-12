import heapq


class PriorityQueue:

    def __init__(self) -> None:
        self.queue = []
        self.index = 0

    def push(self, item, priority: int) -> None:
        # priority 取负值是为了让队列能够按元素的优先级从高到低排列
        heapq.heappush(self.queue, (-priority, self.index, item))
        self.index += 1

    def pop(self):
        return heapq.heappop(self.queue)[-1]


class Item:

    def __init__(self, name):
        self.name = name

    def __repr__(self):
        return 'Item({!r})'.format(self.name)


if __name__ == "__main__":
    pq = PriorityQueue()
    pq.push(Item('A'), 1)
    pq.push(Item('B'), 4)
    pq.push(Item('C'), 5)
    pq.push(Item('D'), 3)
    pq.push(Item('E'), 1)

    # Item('C') Item('B') Item('D') Item('A')
    print(pq.pop(), pq.pop(), pq.pop(), pq.pop())

    a, b, c = (1, Item('A')), (2, Item('B')), (2, Item('C'))

    # True False
    print(a < b, b == c)

    # TypeError: '>' not supported between instances of 'Item' and 'Item'
    # print(b > c)

    # TypeError: '<' not supported between instances of 'Item' and 'Item'
    # print(b < c)

    o, p, q = (1, 0, Item('A')), (5, 1, Item('B')), (1, 2, Item('C'))

    # True True False
    print(a < b, a < c, b == c)

    # TypeError: '>' not supported between instances of 'Item' and 'Item'
    # print(b > c)

    # TypeError: '<' not supported between instances of 'Item' and 'Item'
    # print(b < c)
