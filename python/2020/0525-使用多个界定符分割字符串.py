import re

line = 'Tom Jack; Afred, Dell,Tim Foo:Bar'
items = re.split(r'[,;:\s]\s*', line)

# ['Tom', 'Jack', 'Afred', 'Dell', 'Tim', 'Foo', 'Bar']
print(items)

items = re.split(r'(?:,|;|:|\s)\s*', line)

# ['Tom', 'Jack', 'Afred', 'Dell', 'Tim', 'Foo', 'Bar']
print(items)

items = re.split(r'(,|;|:|\s)\s*', line)

# ['Tom', ' ', 'Jack', ';', 'Afred', ',', 'Dell', ',', 'Tim', ' ', 'Foo', ':', 'Bar']
print(items)

values = items[::2]
delimiters = items[1::2] + [' ']

# ['Tom', 'Jack', 'Afred', 'Dell', 'Tim', 'Foo', 'Bar'] [' ', ';', ',', ',', ' ', ':', '']
print(values, delimiters)

# Tom Jack;Afred,Dell,Tim Foo:Bar
print(''.join(k + v for k, v in zip(values, delimiters)))
