phones = {
    '小米': 1000.00,
    '华为': 4999.00,
    '三星': 3999.00,
    '谷歌': 2999.00,
    '苹果': 6999.00,
    '一加': 5999.00,
    '魅族': 3299.00
}

# (1000.0, '小米')
print(min(zip(phones.values(), phones.keys())))

# (6999.0, '苹果')
print(max(zip(phones.values(), phones.keys())))

# [(1000.0, '小米'), (2999.0, '谷歌'), (3299.0, '魅族'), (3999.0, '三星'), (4999.0, '华为'), (5999.0, '一加'), (6999.0, '苹果')]
print(sorted(zip(phones.values(), phones.keys())))

items = zip(phones.values(), phones.keys())

# (1000.0, '小米')
print(min(items))

# zip() 创建的迭代器只能被消费一次
# print(max(items))

# 1000.0 6999.0
print(min(phones.values()), max(phones.values()))

# 小米
print(min(phones, key=lambda k: phones[k]))

# 苹果
print(max(phones, key=lambda k: phones[k]))
