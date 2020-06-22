nums = [1, 2, 3, 4, 5]
a = sum(i * i for i in nums)

# 55 16 16
print(a, sum((1, 2, 3), 10), sum([1, 2, 3], 10))

import os
import sys

sys.path.append('..')

files = os.listdir('.idea')

# This is a file with an xml suffix
if any(file.endswith('.xml') for file in files):
    print('This is a file with an xml suffix')
else:
    print('This is not a file with an xml suffix')

tups = ('RMB', 20, 128.88)

# RMB 20 128.88
print(' '.join(str(x) for x in tups))

dicts = [
    {'name': 'A', 'price': 56.05},
    {'name': 'B', 'price': 18.00},
    {'name': 'C', 'price': 80.35},
    {'name': 'D', 'price': 12.50}
]
result = min(x['price'] for x in dicts)

# 12.5
print(result)
