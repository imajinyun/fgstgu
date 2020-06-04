squares1 = []
for x in range(10):
    squares1.append(x ** 2)

# [0, 1, 4, 9, 16, 25, 36, 49, 64, 81]
print(squares1)


squares2 = [x ** 2 for x in range(10)]

# [0, 1, 4, 9, 16, 25, 36, 49, 64, 81]
print(squares2)


numbers1 = []
for x in range(50):
    if (x % 3 == 0):
        numbers1.append(x)

# [0, 3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48]
print(numbers1)


numbers2 = [i for i in range(50) if i % 3 is 0]

# [0, 3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48]
print(numbers2)


numbers = [i * i for i in range(50) if i % 3 is 0]

# [0, 9, 36, 81, 144, 225, 324, 441, 576, 729, 900, 1089, 1296, 1521, 1764, 2025, 2304]
print(numbers)


values = ['1', '2', '-3', '-', '4', 'N/A', '5']


def is_int(val):
    try:
        int(val)
        return True
    except ValueError:
        return False


vals = list(filter(is_int, values))

# ['1', '2', '-3', '4', '5']
print(vals)
