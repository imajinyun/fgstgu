a = {
    'x': 1,
    'y': 2,
    'm': 3,
}

b = {
    'x': 3,
    'y': 2,
    'n': 1,
}

"""
多运行几次后会发现两种结果交替出现
"""
# {'x', 'y'} 或 {'y', 'x'}
print(a.keys() & b.keys())

# {'m'}
print(a.keys() - b.keys())

# {('y', 2)}
print(a.items() & b.items())

c = {key: a[key] for key in a.keys() - {'m', 'n'}}

# {'x': 1, 'y': 2} 或 {'y': 2, 'x': 1}
print(c)
