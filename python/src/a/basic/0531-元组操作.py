"""
元组基本操作
"""

tup1 = ('Google', 'Alibaba', 'Facebook', 'Tencent')

# Facebook
print(tup1[2])

# ('Google', 'Alibaba', 'Facebook', 'Tencent')
print(tup1)

# ()
print(tup1[2:2])

# ('Facebook',)
print(tup1[2:3])

# ('Facebook', 'Tencent')
print(tup1[2:4])

# 如果这里定义时不加逗号，则元组拼接时就会报错（TypeError: can only concatenate tuple (not "str") to tuple）
tup2 = ('Microsoft',)

# Microsoft
print(tup2)

tup3 = tup1 + tup2

# ('Google', 'Alibaba', 'Facebook', 'Tencent', 'Microsoft')
print(tup3)

tup = (10, 'A', (55, 88), ["AS", ("WE", "ER")], True,)

# (55, 88)
print(tup[2])

# WE
print(tup[3][1][0])

# ('A', (55, 88), ['AS', ('WE', 'ER')])
print(tup[1:4])

# 10
# A
# (55, 88)
# ['AS', ('WE', 'ER')]
# True
for i in tup:
    print(i)

tup[3].append('欢迎')

# (10, 'A', (55, 88), ['AS', ('WE', 'ER'), '欢迎'], True)
print(tup)

tup[3][1] = False

# (10, 'A', (55, 88), ['AS', False, '欢迎'], True)
print(tup)

tup = tup[:1] + tup[2:]

# (10, (55, 88), ['AS', False, '欢迎'], True)
print(tup)

# 0
print(tup.count('AS'))

# 1
print(tup.count(10))

# 0
print(tup.count(55))

# 1
print(tup.count((55, 88)))

# 4
print(len(tup))

# 1
print(min(('1', '3', '5', '7', '9',)))

# 10
print(max((2, 4, 6, 8, 10,)))

# 3
print(tup.index(1))

del tup

# NameError: name 'tup' is not defined
# print(tup)

# 10 a
# 11 b
# 12 c
# 13 d
for k, v in enumerate(('a', 'b', 'c', 'd'), 10):
    print(k, v)


"""
将字符串转换为元组
"""

s = 'abcdefg'

# ('a', 'b', 'c', 'd', 'e', 'f', 'g')
print(tuple(s))


"""
将列表转换为元组
"""

items = ['a', 'b', 'c', 'd', 'e', 'f', 'g']

# ('a', 'b', 'c', 'd', 'e', 'f', 'g')
print(tuple(items))
