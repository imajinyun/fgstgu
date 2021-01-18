# coding=utf-8

import re


def shorthands() -> None:
    """
    类似于这样 [\s\S]、[\w\W]、[\d\D] 的简记法可以表示任意字符。
    """
    items = [

        # \d 等价于 [0-9]
        re.search(r'^\d$', '8') is not None,  # True

        # \w 等价于 [0-9a-zA-Z_]
        re.search(r'^\w$', '_') is not None,  # True

        # \s 等价于 [\t\r\n\v\f]，第一个字符是空格
        re.search(r'^\s$', '') is not None,  # False
        re.search(r'^\s$', ' ') is not None,  # True
        re.search(r'^\s$', '\t') is not None,  # True

        # 用在普通字符组内部
        re.search(r'^[\da-zA-Z]$', '8') is not None,  # True
        re.search(r'^[\da-zA-Z]$', 'a') is not None,  # True
        re.search(r'^[\da-zA-Z]$', 'C') is not None,  # True

        # 用在排除型字符组内部
        re.search(r'^[^\w]$', '8') is not None,  # False
        re.search(r'^[^\w]$', 'A') is not None,  # False
        re.search(r'^[^\w]$', '_') is not None,  # False
        re.search(r'^[^\w]$', ',') is not None,  # True
        re.search(r'^[^\w]$', '^') is not None,  # True

        # \d 能匹配的字符，\D 一定不能匹配
        re.search(r'^\d$', '8') is not None,  # True
        re.search(r'^\d$', 'a') is not None,  # False
        re.search(r'^\D$', '9') is not None,  # False
        re.search(r'^\D$', 'c') is not None,  # True

        # \w 能匹配的字符，\W 一定不能匹配
        re.search(r'^\w$', 'c') is not None,  # True
        re.search(r'^\w$', '!') is not None,  # False
        re.search(r'^\W$', 'c') is not None,  # False
        re.search(r'^\W$', '!') is not None,  # True

        # \s 能匹配的字符，\S 一定不能匹配
        re.search(r'^\s$', '') is not None,  # False
        re.search(r'^\s$', ' ') is not None,  # True
        re.search(r'^\s$', '\t') is not None,  # True
        re.search(r'^\s$', '0') is not None,  # False
        re.search(r'^\S$', '') is not None,  # False
        re.search(r'^\S$', ' ') is not None,  # False
        re.search(r'^\S$', '\t') is not None,  # False
        re.search(r'^\S$', '0') is not None,  # True
    ]
    for i in range(len(items)):
        print(items[i])


if __name__ == '__main__':
    shorthands()
