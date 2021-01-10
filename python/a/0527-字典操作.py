"""
创建字典元素
"""

dict1 = {'a': 12, 'b': False, 'c': 'ABC', '数学': 66.5}

# {'a': 12, 'b': False, 'c': 'ABC', '数学': 66.5} 12 False ABC 66.5
print(dict1, dict1['a'], dict1['b'], dict1['c'], dict1['数学'])


"""
添加字典元素
"""

dict2 = {'英语': 88, '数学': 72.5}
dict2['语文'] = 90

# {'英语': 88, '数学': 72.5, '语文': 90}
print(dict2)


"""
修改字典元素
"""

dict3 = {'a': 11, 'b': False, 'c': 'this is a test'}
dict3['c'] = 'Hello World'

# {'a': 11, 'b': False, 'c': 'Hello World'}
print(dict3)


"""
删除字典元素
"""

dict4 = {'语文': 98, '数学': 99, '英语': 88}
del dict4['英语']

# {'语文': 98, '数学': 99}
print(dict4)
