from collections import defaultdict

a = {
    'a': [1, 2, 3],
    'b': [4, 5]
}

b = {
    'a': {1, 2, 3},
    'b': [4, 5]
}

c = defaultdict(list)
c['a'].append(1)
c['b'].append(2)
c['c'].append(3)

# defaultdict(<class 'list'>, {'a': [1], 'b': [2], 'c': [3]})
print(c)

d = defaultdict(set)
d['a'].add(1)
d['a'].add(2)
d['a'].add(3)

# defaultdict(<class 'set'>, {'a': {1, 2, 3}})
print(d)

e = {}
e.setdefault('a', []).append(1)
e.setdefault('a', []).append(2)
e.setdefault('a', []).append(3)
e.setdefault('b', []).append(4)
e.setdefault('b', []).append(5)

# {'a': [1, 2, 3], 'b': [4, 5]}
print(e)

f = {}
for k, v in f:
    if k not in f:
        f[k] = []
    f[k].append(v)
f = defaultdict(list)

# defaultdict(<class 'list'>, {})
print(f)

for k, v in f:
    f[k].append(v)

# defaultdict(<class 'list'>, {})
print(f)
