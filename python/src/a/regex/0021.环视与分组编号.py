# coding=utf-8

import re


def solution() -> None:

    items = [
        # 单纯的环视结构并不影响引用分组
        re.search(r'(?!=ab)(cd)', 'abcd').group(0),  # cd
        re.search(r'(?!=ab)(cd)', 'abcd').group(1),  # cd

        # 环视结构中出现了捕获型括号，会影响分组（环视结构中出现了捕获型括号）
        re.search(r'^(?=(ab|cd))', 'abcd').group(0),  # ''
        re.search(r'^(?=(ab|cd))', 'abcd').group(1),  # ab

        # 环视结构指定使用非捕获括号 -> IndexError: no such group
        # re.search(r'^(?=(?:ab|cd))', 'abcd').group(1),

        re.search(r'(\d+)\w+\1', '123a12').group(0),  # 123a12
    ]
    [print(items[i]) for i in range(len(items))]


if __name__ == "__main__":
    solution()
