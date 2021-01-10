from collections import ChainMap

a = {'x': 1, 'z': 3}
b = {'y': 2, 'z': 4}
c = ChainMap(a, b)

# 1 2 3
print(c['x'], c['y'], c['z'])

# 3
print(len(c))

# ['y', 'z', 'x'] [2, 3, 1]
print(list(c.keys()), list(c.values()))

c['z'] = 10
c['w'] = 40
del c['x']

# {'z': 10, 'w': 40} {'y': 2, 'z': 4} ChainMap({'z': 10, 'w': 40}, {'y': 2, 'z': 4})
print(a, b, c)
