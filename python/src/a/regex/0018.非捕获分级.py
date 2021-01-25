# coding=utf-8

import re


def solution() -> None:
    """
    非捕获分组不需要保存匹配的文本，整个表达式的效率也因此提高，但是看起来不如捕获分组美观。
    """
    print(re.search(r'(\d{4})-(\d{2})-(\d{2})', '2021-01-25').group(2))
    print(re.search(r'(?:\d{4})-(\d{2})-(\d{2})', '2021-01-25').group(1))


if __name__ == '__main__':
    solution()
