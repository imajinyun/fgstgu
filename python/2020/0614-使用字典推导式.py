dicts = {
    'a': 11,
    'b': 55,
    'A': 77,
    'B': 33,
    'C': 44,
}
result = {
    k.lower(): dicts.get(k.lower(), 0) + dicts.get(k.upper(), 0)
    for k in dicts.keys() if k.lower() in ['a', 'b']
}

# {'a': 88, 'b': 88}
print(result)


dicts = {
    'a': 33,
    'b': 77,
}
result = {v: k for k, v in dicts.items()}

# {33: 'a', 77: 'b'}
print(result)


dicts = {
    'JavaScript': 49.9,
    'Python': 69.9,
    'Java': 59.9,
    'Go': 45.9,
    'PHP': 79.9,
}
a = {key: value for key, value in dicts.items() if value > 50}

# {'Python': 69.9, 'Java': 59.9, 'PHP': 79.9}
print(a)

filters = {'Python', 'Java', 'JavaScript'}
b = {key: value for key, value in dicts.items() if key in filters}

# {'JavaScript': 49.9, 'Python': 69.9, 'Java': 59.9}
print(b)

# 这种方式运行慢
c = dict((key, value) for key, value in dicts.items() if value > 50)

# {'Python': 69.9, 'Java': 59.9, 'PHP': 79.9}
print(c)

# 这种方式运行慢
d = {key: dicts[key] for key in dicts.keys() if key in filters}

# {'JavaScript': 49.9, 'Python': 69.9, 'Java': 59.9}
print(d)
