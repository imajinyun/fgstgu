
def reverse1(s) -> str:
    return s[::-1]


def reverse2(s) -> str:
    items = list(s)
    items.reverse()
    return ''.join(items)


def reverse3(s) -> str:
    items = list(s)
    return ''.join(items[::-1])


def reverse4(s) -> str:
    from functools import reduce
    return reduce(lambda x, y: y + x, s)


def reverse5(s) -> str:
    if (len(s) < 1):
        return s
    return reverse5(s[1:]) + s[0]


def reverse6(s) -> str:
    items = list(s)
    result = ''
    while len(items) > 0:
        result += items.pop()
    return result


def reverse7(s) -> str:
    result = ''
    length = len(s) - 1
    for key, val in enumerate(s):
        result += s[length - key]
    return result


s = 'HelloWorld'
print(
    # dlroWolleH
    reverse1(s),

    # dlroWolleH
    reverse2(s),

    # dlroWolleH
    reverse3(s),

    # dlroWolleH
    reverse4(s),

    # dlroWolleH
    reverse5(s),

    # dlroWolleH
    reverse6(s),

    # dlroWolleH
    reverse7(s),
)
