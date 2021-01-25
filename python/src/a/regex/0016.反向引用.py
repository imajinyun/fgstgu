# coding=utf-8

import re


def solution() -> None:
    regex = r'^([a-z])\1$'
    items = [
        re.search(regex, 'aa') is not None,  # True
        re.search(regex, 'dd') is not None,  # True
        re.search(regex, 'ac') is not None,  # False
        re.search(regex, 'boot') is not None,  # False
    ]
    [print(items[i]) for i in range(len(items))]


def solution2() -> None:
    regex = r'^<([^>]+)>[\s\S]*?</\1>$'
    items = [
        re.search(regex, '<strong>Test</strong>') is not None,  # True
        re.search(regex, '<h1>Title</h1>') is not None,  # True
        re.search(regex, '<h1>Text</body>') is not None,  # False
    ]
    [print(items[i]) for i in range(len(items))]


def solution3() -> None:
    regex = r'<([a-zA-Z0-9]+)(\s[^>]+)?>[\s\S]*?</\1>'
    text = '<span class="class-text">Text</span>'
    items = [
        re.search(regex, '<body>Test</body>') is not None,  # True
        re.search(regex, '<b>Test</b>') is not None,  # True
        re.search(regex, text) is not None,  # True
    ]
    [print(items[i]) for i in range(len(items))]


def solution4() -> None:
    segment = r'(0{0,2}[0-9]|0?[0-9]{2}|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'
    regex = r'^(' + segment + r'\.){3}' + segment + r'$'
    items = [
        re.search(regex, '0.0.0.0') is not None,  # True
        re.search(regex, '127.0.0.1') is not None,  # True
        re.search(regex, '192.168.1.1') is not None,  # True
        re.search(regex, '001.08.01.010') is not None,  # True
        re.search(regex, '001.08.01.010') is not None,  # True
        re.search(regex, '008.008.008.008') is not None,  # True
        re.search(regex, '070.070.070.070') is not None,  # True
        re.search(regex, '255.255.255.0') is not None,  # True
        re.search(regex, '255.255.255.255') is not None,  # True
        re.search(regex, '255.000.255.255') is not None,  # True
        re.search(regex, '255.000.255.256') is not None,  # False
        re.search(regex, '288.000.255.256') is not None,  # False
        re.search(regex, '356.000.255.256') is not None,  # False
        re.search(regex, '999.000.255.258') is not None,  # False
    ]
    [print(items[i]) for i in range(len(items))]


def solution5() -> None:
    # 使用 g<n> 消除二义性
    items = [
        re.sub(r'(\d)', r'\g<1>0', '123'),  # 102030
        re.sub(r'(\d)', r'\g<1>0', '102030'),  # 100020003000
    ]
    [print(items[i]) for i in range(len(items))]


if __name__ == '__main__':
    solution()
    solution2()
    solution3()
    solution4()
    solution5()
