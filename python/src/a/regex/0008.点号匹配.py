# coding=utf-8

import re


def solution() -> None:
    quote_string = '"quote string" and another"'
    items = [
        # 点号的匹配
        re.search(r'^.$', 'a') is not None,  # True
        re.search(r'^.$', '0') is not None,  # True
        re.search(r'^.$', '*') is not None,  # True
        re.search(r'^.$', '`.') is not None,  # False
        re.search(r'^.$', '11') is not None,  # False
        re.search(r'^.$', '🍎') is not None,  # True

        # 换行符的匹配
        re.search(r'^.$', '\n') is not None,  # False
        re.search(r'(?s)^.$', '\n') is not None,  # True
        re.search(r'^[\s\S]$', '\n\n') is not None,  # True

        # 匹配优先量词（greedy quantifier，又称贪婪量词）。
        # 匹配优先量词，就是在拿不准是否要匹配的时候，优先尝试匹配，并且记下这个状态，以备将来『反悔』。
        # "quote string" and another"
        re.search(r'".*"', quote_string).group(0),

        # "quote string"
        re.search(r'"[^"]*"', quote_string).group(0),

        # "quote string"
        re.search(r'"[^"].*?"', quote_string).group(0),
    ]
    [print(items[i]) for i in range(len(items))]
    print()


def solution2() -> None:
    regex = r'((?!\\)"|[^"])*'
    quote_string1 = "quoted string with \"nested quote\""

    # quoted string with "nested quote"
    print(re.search(regex, quote_string1).group(0))

    quote_string2 = '"escaped blackslash\\" and "another"'

    # "escaped blackslash\" and "another"
    print(re.search(regex, quote_string2).group(0))

    quote_string3 = '"escaped backslash\\\"" and "another"'

    # "escaped backslash\"" and "another"
    print(re.search(regex, quote_string3).group(0))


if __name__ == "__main__":
    solution()
    solution2()
