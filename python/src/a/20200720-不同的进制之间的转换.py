def bin2Any(s: str, base: int) -> int:
    """ 二进制转其它进制 """
    if base == 8:
        return oct(int(s, 2))
    elif base == 10:
        return int(s, 2)
    elif base == 16:
        return hex(int(s, 2))
    else:
        return -1


def oct2Any(s: str, base: int) -> int:
    """ 八进制转其它进制 """
    if base == 2:
        return bin(int(s, 8))
    elif base == 10:
        return int(s, 8)
    elif base == 16:
        return hex(int(s, 8))
    else:
        return -1


def dec2Any(s: str, base: int) -> int:
    """ 十进制转其它进制 """
    if base == 2:
        return bin(int(s, 10))
    elif base == 8:
        return oct(int(s, 10))
    elif base == 16:
        return hex(int(s, 10))
    else:
        return -1


def hex2Any(s: str, base: int) -> int:
    """ 十六进制转其它进制 """
    if base == 2:
        return bin(int(s, 16))
    elif base == 8:
        return oct(int(s, 16))
    elif base == 10:
        return int(s, 16)
    else:
        return -1


# 二进制转其它进制
a, b, c = '0b10000', '0b11000', '0b11100'
print(bin2Any(a, 8), bin2Any(b, 10), bin2Any(c, 16))

# 八进制转其它进制
a, b, c = '0o1', '0o7', '0o14'
print(oct2Any(a, 2), oct2Any(b, 10), oct2Any(c, 16))

# 十进制转其它进制
a, b, c = '1', '8', '24'
print(dec2Any(a, 2), dec2Any(b, 8), dec2Any(c, 16))

# 十六进制转其它进制
a, b, c = '0xFF', '0x38', '0xAF'
print(hex2Any(a, 2), hex2Any(b, 8), hex2Any(c, 10))
