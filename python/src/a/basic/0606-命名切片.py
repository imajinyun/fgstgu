items = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
a = slice(2, 4)

# slice(2, 4, None)
print(a)

# [2, 3] [2, 3]
print(items[2:4], items[a])

items[a] = [88, 99]

# [0, 1, 88, 99, 4, 5, 6, 7, 8, 9]
print(items)

# 2 4 None
print(a.start, a.stop, a.step)

del items[a]

# [0, 1, 4, 5, 6, 7, 8, 9]
print(items)

s = 'HelloWorld'

# (2, 4, 1)
print(a.indices(len(s)))

"""
通过调用切片的 indices(size) 方法将它映射到一个已知大小的序列上

l
l
"""
for i in range(*a.indices(len(s))):
    print(s[i])
