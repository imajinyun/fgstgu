# coding=utf-8

import re


def negated_character_class() -> None:
    items = [
        # 第一个匹配一个除 0~9 之外的字符，第二个匹配一个 0~9 之间的字符
        re.search(r'^[^0-9][0-9]$', 'A8') is not None,  # True
        re.search(r'^[^0-9][0-9]$', 'x6') is not None,  # True
        re.search(r'^[^0-9][0-9]$', '^6') is not None,  # True
        re.search(r'^[^0-9][0-9]$', '8') is not None,  # False
        re.search(r'^[^0-9][0-9]$', 'A9') is not None,  # True

        # 匹配一个 -、0、9 之外的字符
        re.search(r'^[^-09]$', '-') is not None,  # False
        re.search(r'^[^-09]$', '8') is not None,  # True

        # 匹配一个 0~9 之外的字符
        re.search(r'^[^0-9]$', '-') is not None,  # True
        re.search(r'^[^0-9]$', '8') is not None,  # False

        # 匹配一个 0、1、2 之外的字符
        re.search(r'^[^012]$', '^') is not None,  # True

        # 匹配 0、^、1、2 四个字符中的一个
        re.search(r'^[0^12]$', '^') is not None,  # True
    ]
    for i in range(len(items)):
        print(items[i])


if __name__ == '__main__':
    negated_character_class()
