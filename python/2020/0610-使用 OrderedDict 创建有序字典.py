import collections

dicts = collections.OrderedDict()
dicts['k1'] = 'v1'
dicts['k2'] = 'v2'
dicts['k3'] = 'v3'

# OrderedDict([('k1', 'v1'), ('k2', 'v2'), ('k3', 'v3')])
print(dicts)

# ('k3', 'v3') OrderedDict([('k1', 'v1'), ('k2', 'v2')])
print(dicts.popitem(), dicts)

# ('k2', 'v2') OrderedDict([('k1', 'v1')])
print(dicts.popitem(), dicts)

dicts.clear()

# OrderedDict()
print(dicts)
