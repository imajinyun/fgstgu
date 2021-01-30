# coding=utf-8

import re
from enum import unique


def solution() -> None:
    """
    准确删除注释。
    """
    regex = r'(?m)^(((?<!/)/(?!/)|\"(\\.|[^\\\"])*\"|[^/\"])*)//.*$'
    items = [
        # -> String s = "some text";
        re.sub(regex, r'\1', "String s = \"some text\"; //this is a test\""),

        # -> String url = "http://t.com";
        re.sub(regex, r'\1', "String url = \"http://t.com\";//this is a test"),

        # -> String url = "http://t.com";
        re.sub(regex, r'\1', "String url = \"http://t.com\";"),

        # -> String s = "a b";
        re.sub(regex, r'\1', "String s = \"a b\"; // a with \"some text\""),
    ]
    [print(items[i]) for i in range(len(items))]


def solution2() -> None:
    """
    回调函数处理四舍五入。
    """
    def process(matcher) -> float:
        match = matcher.group(0)
        if re.search('\\.[0-4]', match) is not None:
            return '%.1f' % float(match)
        else:
            return '%.1f' % (int(float(match)) + 1)

    # -> 3.0 10.2 0.1 1.3 2.5
    text = "2.618 10.236 0.123456789 1.3456789 2.456789"
    print(re.sub(r'\d+\.\d+', process, text))


def solution3() -> None:
    def process(matcher) -> str:
        if matcher.group(0).startswith('&#x'):  # 十六进制
            return chr(int(matcher.group(1), 16))
        else:  # 十进制
            return chr(int(matcher.group(1), 10))

    text1 = '&#x6536;&#x53D1;'
    text2 = '&#25910;&#21457;'
    print(re.sub(r'&#x?([0-9a-fA-F]+);', process, text1))
    print(re.sub(r'&#x?([0-9a-fA-F]+);', process, text2))


if __name__ == "__main__":
    solution()
    solution2()
    solution3()
