"""
分解一个序列的过程
"""

a = ('A', 99)
x, y = a

# A 99
print(x, y)

data = ['ABCD', 99, 88.99, (1, 2, 3)]

i, j, m, n = data

# ABCD 99 88.99 (1, 2, 3)
print(i, j, m, n)


"""
分解一个带标记元组序列的过程
"""

records = [
    ('AAA', 11, 22),
    ('BBB', 'hello'),
    ('CCC', 77, 88),
]


def do_foo(x, y):
    print('AAA', x, y)


def do_bar(s):
    print('BBB', s)


# AAA 11 22
# BBB hello
for tag, *args in records:
    if tag == 'AAA':
        do_foo(*args)
    elif tag == 'BBB':
        do_bar(*args)

line = 'bei:jing2020://http:1978:bei'
o, *p, q, s = line.split(':')

# bei
print(o)

# ['jing2020', '//http']
print(p)

# 1978
print(q)

# bei
print(s)


"""
分解一个字符串的过程
"""

s = 'Hello'
a, b, c, d, e = s

# H e l l o
print(a, b, c, d, e)

data = ['Acme', False, 99.99, (1, 2, 3), None, {'a': 11, 'b': 12}]
_, _, _, tuples, _, dicts = data

# (1, 2, 3) {'a': 11, 'b': 12}
print(tuples, dicts)
