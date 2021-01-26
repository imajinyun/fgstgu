# coding=utf-8

import re


def solution() -> None:
    """
    ^ 既可以匹配整个字符串的起始位置，也可以匹配换行符之后的
    位置（设定多行模式最简单的办法是在正则表达式之前加上 (?m)）。
    """
    # 提取每行的第一个单词
    print(re.findall(r'(?m)^\w+', 'a\nb\nc\nd'))

    # 匹配整段文本的第一个单词
    print(re.findall(r'(?m)\A\w+', 'a\nb\nc\nd'))

    # 匹配整段文本的最后一个单词
    print(re.findall(r'\w+$', 'hello world'))
    print(re.findall(r'(?m)\w+$', 'hello python'))


def solution2() -> None:
    items = [
        # 仅使用 ^ 和 $ 验证 6 位数字字符串存在的问题
        re.search(r'^\d{6}$', '123456') is not None,  # True
        re.search(r'^\d{6}\n$', '123456\n') is not None,  # True

        # 借助 \A 和 \Z 完成更准确的数据验证
        re.search(r'\A\d{6}\Z', '123456') is not None,  # True
        re.search(r'\A\d{6}\Z', '123456\n') is not None,  # False
    ]
    [print(items[i]) for i in range(len(items))]


def solution3() -> None:
    text = 'a\nb\nc'
    print(re.sub(r'(?m)$', '</p>', re.sub(r'(?m)^', '<p>', text)))


def solution4() -> None:
    """
    去除行首的空白字符：

    [    a]
    [  b    ]
    [c  ]

    ->

    [a]
    [b    ]
    [c  ]

    去除行尾的空白字符：

    ->

    [a]
    [b]
    [c]
    """

    # 去除行首的空白字符
    text = '    a\n  b\t\nc  '
    text = re.sub(r'(?m)^\s+', '', text)
    print(text)

    # 去除行尾的空白字符
    text = re.sub(r'(?m)\s+$', '', text)
    print(text)


if __name__ == '__main__':
    solution()
    solution2()
    solution3()
    solution4()
