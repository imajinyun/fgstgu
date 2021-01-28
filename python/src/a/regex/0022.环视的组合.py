# coding=utf-8

import re


def solution() -> None:
    # 并列多个环视
    regex = r'^(?=\d+)(?!999)'
    items = [
        re.search(regex, '1234') is not None,  # True
        re.search(regex, 'example') is not None,  # False
        re.search(regex, '999example') is not None,  # False
        re.search(regex, 'example999') is not None,  # False
        re.search(regex, '999123456') is not None,  # False
        re.search(regex, '123456999') is not None,  # True
    ]
    [print(items[i]) for i in range(len(items))]


def solution2() -> None:
    # 环视作为多选分支
    regex = r'^((?!\d)|(?=\d\D))'
    items = [
        re.search(regex, 'ab') is not None,  # True
        re.search(regex, 'a4') is not None,  # True
        re.search(regex, 'aaa888') is not None,  # True
        re.search(regex, '4') is not None,  # False
        re.search(regex, '44') is not None,  # False
        re.search(regex, '444444aaaaaa') is not None,  # False
    ]
    [print(items[i]) for i in range(len(items))]


def solution3() -> None:
    items = [
        # 反向引用时不会保留断言的判断（反向引用时，之前捕获分组中的断言都会被忽略）
        re.search(r'(\bcat\b).*?\1', 'cat cate') is not None,  # True
        re.search(r'(\bcat\b)\s+\b\1\b', 'cat cate') is not None,  # False
        re.search(r'(\bcat\b)\s+\b\1\b', 'cat cat') is not None,  # True
    ]
    [print(items[i]) for i in range(len(items))]


if __name__ == '__main__':
    solution()
    solution2()
    solution3()
