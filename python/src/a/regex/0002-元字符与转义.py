# coding=utf-8

import re


def raw_string() -> None:
    """
    出现在不同位置，匹配含义不同。

    注意：只有开方括号 [ 需要转义，闭方括号 ] 不需要转义。
    """
    items = [
        r"^[0\-9]$" == "^[0\\-9]$",  # True
        re.search(r'^[0\-9]$', '3') is not None,  # False
        re.search(r'^[0\-9]$', '-') is not None,  # True
        re.search(r'^[0\-9]$', '0') is not None,  # True

        re.search(r'^[012]345$', '2345') is not None,  # True
        re.search(r'^[012]345$', '2]345') is not None,  # False
        re.search(r'^[012]345$', '5') is not None,  # False
        re.search(r'^[012]345$', ']') is not None,  # False
        re.search(r'^[012\]345]$', '2345') is not None,  # False
        re.search(r'^[012\]345]$', '2]345') is not None,  # False
        re.search(r'^[012\]345]$', '5') is not None,  # True
        re.search(r'^[012\]345]$', '3') is not None,  # True
        re.search(r'^[012\]345]$', ']') is not None,  # True

        re.search(r'^[012]345]$', '2345') is not None,  # False
        re.search(r'^[012]345]$', '3') is not None,  # False
        re.search(r'^[012]345]$', ']') is not None,  # False
        re.search(r'^[012\]345]$', '3') is not None,  # True
        re.search(r'^[012]$]', '[012]') is not None,  # False
        re.search(r'^\[012]$', '[012]') is not None,  # True
    ]
    for i in range(len(items)):
        print(items[i])


if __name__ == '__main__':
    raw_string()
