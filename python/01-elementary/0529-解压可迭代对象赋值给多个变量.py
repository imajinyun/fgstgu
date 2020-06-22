"""
解析列表中的中间数据
"""
items = [99.5, 88, 73, 82.5, 77, 90.5, 82]
first, *middle, last = items

# 99.5 [88, 73, 82.5, 77, 90.5] 82
print(first, middle, last)


"""
解析列表中某类特征数据
"""
rows = ['Tom', 'jack@example.com', '773-2323233', '882-1298983', '662-2898989']
name, email, *phone = rows

# Tom jack@example.com ['773-2323233', '882-1298983', '662-2898989']
print(name, email, phone)
