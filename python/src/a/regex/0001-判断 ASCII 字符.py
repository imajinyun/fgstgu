# coding=utf-8

import re


def match_ascii(ch: str) -> bool:
    return re.search("^[\x00-\x7F]$", ch) is not None


if __name__ == '__main__':
    chars = ['A', 'i', '<', '[', '*', '`', '@', '•', '0', '!']
    for i in range(len(chars)):
        print(chars[i] + ':', match_ascii(chars[i]))
