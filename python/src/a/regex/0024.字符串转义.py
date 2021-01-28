# coding=utf-8

import re


def solution() -> None:
    # 转义
    items = [
        re.search('\\\\', '\\') is not None,  # True
        re.search('\\n', '\n') is not None,  # True
        re.search('\n', '\n') is not None,  # True
        re.search('\\t', '\t') is not None,  # True
        re.search('\t', '\t') is not None,  # True
        re.search('\ba\b', 'a') is not None,  # False
        re.search('\\ba\\b', 'a') is not None,  # True
        re.search('\d', '1') is not None,  # True
        re.search('\(', '(') is not None,  # True
        re.search(re.escape('ca*t'), 'ca*t') is not None,  # True
        re.search(r'ca*t', 'ca*t') is not None,  # False
        re.search(r'ca\*t', 'ca*t') is not None,  # True
        re.search(r'[*]', '*') is not None,  # True
        re.search(r'[?]', '?') is not None,  # True
        re.search(r'\[()', '(') is not None,  # False
    ]
    [print(items[i]) for i in range(len(items))]


def solution2() -> None:
    items = [
        re.search(r'[0\]9]', ']') is not None,  # True
        re.search(r'[0-9]', '-') is not None,  # False
        re.search(r'[0\-9]', '-') is not None,  # True
        re.search(r'[-09]', '-') is not None,  # True
        re.search(r'[-09]', '-') is not None,  # True
        re.search(r'[^ab]', '^') is not None,  # True
        re.search(r'[^ab]', 'a') is not None,  # False
        re.search(r'[\^ab]', '^') is not None,  # True
        re.search(r'[\^ab]', 'a') is not None,  # True
        re.search(r'[a^b]', '^') is not None,  # True
        re.search(r'[a^b]', 'a') is not None,  # True
    ]
    [print(items[i]) for i in range(len(items))]


def solution3() -> None:
    print('\n')
    items = [
        re.search(r'[中国]', '中国') is not None,  # True
        re.search(r'^.$', '中') is not None,  # True
        re.search(r'^.$', '中国') is not None,  # False
        re.findall(r'\bregex\b', '正则 regex'),  # ['regex']
        re.findall(r'(?u)\bregex\b', '正则 regex'),  # ['regex']
    ]
    [print(items[i]) for i in range(len(items))]
    print('\n')
    re.compile(r"[中国]", re.DEBUG)


if __name__ == '__main__':
    solution()
    solution2()
    solution3()
