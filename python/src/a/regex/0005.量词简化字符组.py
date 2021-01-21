# coding=utf-8

import re


def solution() -> None:
    items = [
        re.search(r'^\d{6}$', '100860') is not None,  # True
        re.search(r'^\d{6}$', '100A00') is not None,  # False
    ]
    [print(items[i]) for i in range(len(items))]


def solution2() -> None:
    regex = r'^\d{4,6}$'
    items = [
        # 元素至少出现 m 次，最多出现 n 次
        re.search(regex, '123') is not None,  # False
        re.search(regex, '1234') is not None,  # True
        re.search(regex, '123456') is not None,  # True
        re.search(regex, '1234567') is not None,  # False
    ]
    [print(items[i]) for i in range(len(items))]


def solution3() -> None:
    regex = r'^\d{4,}$'
    items = [
        # 元素最少出现 m 次，出现次数无上限（隐式的上限是 65536）
        re.search(regex, '123') is not None,  # False
        re.search(regex, '1234') is not None,  # True
        re.search(regex, '123456789') is not None,  # True
    ]
    [print(items[i]) for i in range(len(items))]


def solution4() -> None:
    regex = r'^\d{0,6}$'
    items = [
        # 元素可以不出现，也可以出现，最多出现 n 次（在某些语言中可以写为 {,n}）
        re.search(regex, '12345') is not None,  # True
        re.search(regex, '123456') is not None,  # True
        re.search(regex, '1234567') is not None,  # False
    ]
    [print(items[i]) for i in range(len(items))]


if __name__ == "__main__":
    solution()
    solution2()
    solution3()
    solution4()
