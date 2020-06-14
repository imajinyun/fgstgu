from operator import itemgetter

a = [1, 2, 3, 4, 5]
b = itemgetter(1)

# 2
print(b(a))

c = itemgetter(1, 0)

# (2, 1)
print(c(a))


rows = [
    {'group': 'A', 'name': 'Zhan', 'uid': 1001},
    {'group': 'B', 'name': 'Zhou', 'uid': 1002},
    {'group': 'C', 'name': 'Weng', 'uid': 1004},
    {'group': 'D', 'name': 'Ling', 'uid': 1003},
]

rows_by_group = sorted(rows, key=itemgetter('group'))
rows_by_uid = sorted(rows, key=itemgetter('uid'))

"""
[
    {'group': 'A', 'name': 'Zhan', 'uid': 1001},
    {'group': 'B', 'name': 'Zhou', 'uid': 1002},
    {'group': 'C', 'name': 'Weng', 'uid': 1004},
    {'group': 'D', 'name': 'Ling', 'uid': 1003}
]
"""
print(rows_by_group)

"""
[
    {'group': 'A', 'name': 'Zhan', 'uid': 1001},
    {'group': 'B', 'name': 'Zhou', 'uid': 1002},
    {'group': 'D', 'name': 'Ling', 'uid': 1003},
    {'group': 'C', 'name': 'Weng', 'uid': 1004}
]
"""
print(rows_by_uid)

rows_by_group_name = sorted(rows, key=itemgetter('name', 'group'))

"""
[
    {'group': 'D', 'name': 'Ling', 'uid': 1003},
    {'group': 'C', 'name': 'Weng', 'uid': 1004},
    {'group': 'A', 'name': 'Zhan', 'uid': 1001},
    {'group': 'B', 'name': 'Zhou', 'uid': 1002}
]
"""
print(rows_by_group_name)

rows_by_group = sorted(rows, key=lambda k: k['group'])
rows_by_group_name = sorted(rows, key=lambda k: (k['group'], k['name']))

"""
[

    {'group': 'A', 'name': 'Zhan', 'uid': 1001},
    {'group': 'B', 'name': 'Zhou', 'uid': 1002},
    {'group': 'C', 'name': 'Weng', 'uid': 1004},
    {'group': 'D', 'name': 'Ling', 'uid': 1003}
]
"""
print(rows_by_group)

"""
[

    {'group': 'A', 'name': 'Zhan', 'uid': 1001},
    {'group': 'B', 'name': 'Zhou', 'uid': 1002},
    {'group': 'C', 'name': 'Weng', 'uid': 1004},
    {'group': 'D', 'name': 'Ling', 'uid': 1003}
]
"""
print(rows_by_group_name)

# {'group': 'A', 'name': 'Zhan', 'uid': 1001}
print(min(rows, key=itemgetter('uid')))

# {'group': 'C', 'name': 'Weng', 'uid': 1004}
print(max(rows, key=itemgetter('uid')))
