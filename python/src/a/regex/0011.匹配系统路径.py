# coding=utf-8

import re


def solution() -> None:
    a, b = '/usr/local/bin/python3', 'C:\\Program Files\\Python 3.8.4\\Python.exe'
    items = [
        # 匹配 Linux/Unix 路径
        re.search(r'^.*/', a).group(0),  # /usr/local/bin/
        re.search(r'[^/]*$', a).group(0),  # python3

        # 匹配 Windows 路径
        re.search(r'^.*\\', b).group(0),  # C:\Program Files\Python 3.8.4\
        re.search(r'[^\\]*$', b).group(0),  # Python.exe
    ]
    [print(items[i]) for i in range(len(items))]


if __name__ == "__main__":
    solution()
