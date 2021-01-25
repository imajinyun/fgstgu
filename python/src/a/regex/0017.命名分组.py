# coding=utf-8

import re


def solution() -> None:
    """
    捕获分组通常用数字编号来标识，但这样有几个问题：

    1. 数字编号不够直观，虽然规则是『从左向右按照开括号出现的顺序计数』，括号多了难免混淆；
    2. 引用时也不够方便；

    命名分组（named grouping），可以将它看作另一种捕获分组，但是标识是容易记忆和辨别的名字，而不是数字编号。

    在 Python 中，如果使用了命名分组，在表达式中反向引用时，必须使用 (?P=name) 的记法；
    而要进行正则表达式替换，则需要写作 \g<name>，其中的 name 是分组的名字。
    """
    regex = r'(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})'
    result = re.search(regex, "2021-01-25")

    # 2021 01 25
    print(result.group('year'), result.group('month'), result.group('day'))

    # 2021 01 25
    print(result.group(1), result.group(2), result.group(3))

    # True
    print(re.search(r'^(?P<char>[a-z])(?P=char)$', 'aa') is not None)

    # True
    print(re.sub('(?P<digit>\d)', r'\g<digit>0', '123') is not None)


if __name__ == "__main__":
    solution()
