# coding=utf-8

import re


def solution() -> None:
    """
    出现在不同位置，匹配含义不同。

    注意：只有开方括号 [ 需要转义，闭方括号 ] 不需要转义。
    """
    regex1 = r'^[0\-9]$'
    regex2 = r'^[012]345$'
    regex3 = r'^[012]345]$'
    items = [
        r"^[0\-9]$" == "^[0\\-9]$",  # True
        re.search(regex1, '3') is not None,  # False
        re.search(regex1, '-') is not None,  # True
        re.search(regex1, '0') is not None,  # True
        re.search(regex2, '2345') is not None,  # True
        re.search(regex2, '2]345') is not None,  # False
        re.search(regex2, '5') is not None,  # False
        re.search(regex2, ']') is not None,  # False
        re.search(regex3, '2345') is not None,  # False
        re.search(regex3, '3') is not None,  # False
        re.search(regex3, ']') is not None,  # False
        re.search(r'^[012\]345]$', '3') is not None,  # True
        re.search(r'^[012]$]', '[012]') is not None,  # False
        re.search(r'^\[012]$', '[012]') is not None,  # True
    ]
    [print(items[i]) for i in range(len(items))]


def solution2() -> None:
    regex = r'^[012\]345]$'
    items = [
        re.search(regex, '2345') is not None,  # False
        re.search(regex, '2]345') is not None,  # False
        re.search(regex, '5') is not None,  # True
        re.search(regex, '3') is not None,  # True
        re.search(regex, ']') is not None,  # True
    ]
    [print(items[i]) for i in range(len(items))]


if __name__ == '__main__':
    solution()
    solution2()
