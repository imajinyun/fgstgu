# coding=utf-8

import re


def solution() -> None:
    items = [
        re.sub(r'[a-z]', ' ', "1a2b3c4d5f"),
        re.sub(r'[0-9]', ' ', "1a2b3c4d5f"),
    ]
    [print(items[i]) for i in range(len(items))]


def solution2() -> None:
    regex = r'(\d{4})-(\d{2})-(\d{2})'

    items = [
        # 12/31/2020
        re.sub(regex, r'\2/\3/\1', '2020-12-31'),

        # 2020年12月31日
        re.sub(r'(\d{4})-(\d{2})-(\d{2})', r'\1年\2月\3日', '2020-12-31'),

        # ASCII 编码为 0 的字符无法显示
        re.sub(r'(\d{4})-(\d{2})-(\d{2})', '\\0', '2020-12-31'),

        # ASCII 编码为 0 的字符无法显示
        re.sub(r'(\d{4})-(\d{2})-(\d{2})', r'\0', '2020-12-31'),

        # [2020-12-31]
        re.sub(r'((\d{4})-(\d{2})-(\d{2}))', '[\\1]', '2020-12-31'),

        # [2020-12-31]
        re.sub(r'((\d{4})-(\d{2})-(\d{2}))', r'[\1]', '2020-12-31'),
    ]
    [print(items[i]) for i in range(len(items))]


if __name__ == "__main__":
    solution()
    solution2()
