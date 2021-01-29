# coding=utf-8

import re


def solution() -> None:
    """
    必须出现的匹配。

    以下正则表达式并不能过滤 2021.01.31 这样的字符串：

    r'(\+?(\d+|\.\d+|\d+\.\d+)|\-?(\d+|\d+\.\d+))'
    """
    regex1 = r'^(?!<\.)(\+?(\d+|\.\d+|\d+\.\d+)|-?(\d+|\d+\.\d+))(?!\.)$'
    items1 = [
        re.search(regex1, '3.1415926') is not None,  # True
        re.search(regex1, '0.1415926') is not None,  # True
        re.search(regex1, '1000.00001415926') is not None,  # True
        re.search(regex1, '+0.1415926') is not None,  # True
        re.search(regex1, '-1000.1415926') is not None,  # True
        re.search(regex1, '10') is not None,  # True
        re.search(regex1, '.3') is not None,  # True
        re.search(regex1, '+.3') is not None,  # True
        re.search(regex1, '-0.3') is not None,  # True
        re.search(regex1, '2021.01.28') is not None,  # False
    ]
    regex2 = r'^(?=.+)(?:[1-9]\d*|0)?(?:\.\d+)?$'
    items2 = [
        re.search(regex2, '3.1415926') is not None,  # True
        re.search(regex2, '0.1415926') is not None,  # True
        re.search(regex2, '1000.00001415926') is not None,  # True
        re.search(regex2, '0.000001') is not None,  # True
        re.search(regex2, '.000001') is not None,  # True
        re.search(regex1, '10') is not None,  # True
        re.search(regex2, '.3') is not None,  # True
        re.search(regex2, '+0.1415926') is not None,  # False
        re.search(regex2, '-1000.1415926') is not None,  # False
        re.search(regex2, '+.3') is not None,  # False
        re.search(regex2, '-0.3') is not None,  # False
        re.search(regex1, '2021.01.28') is not None,  # False
    ]
    [print(items1[i]) for i in range(len(items1))]
    print('\n')
    [print(items2[i]) for i in range(len(items2))]
    print('\n')


if __name__ == '__main__':
    solution()
