dicts = {
    'a': 42,
    'b': 28,
    'c': 99,
    'd': 72,
    'e': 23,
    'f': 61,
}

a = {key: val for key, val in dicts.items() if val >= 60}
b = dict((key, val) for key, val in dicts.items() if val >= 60)

# {'c': 99, 'd': 72, 'f': 61} {'c': 99, 'd': 72, 'f': 61}
print(a, b)

searchs = {'a', 'c', 'e'}
c = {key: val for key, val in dicts.items() if key in searchs}
d = {key: dicts[key] for key in dicts.keys() & searchs}

# {'a': 42, 'c': 99, 'e': 23} {'a': 42, 'c': 99, 'e': 23}
print(c, d)
