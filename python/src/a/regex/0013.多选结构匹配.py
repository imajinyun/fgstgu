# coding=utf-8

import re


def solution() -> None:
    # 匹配身份证号
    regex = r'^([1-9]\d{14}|[1-9]\d{16}[0-9xX])$'
    items = [
        re.search(regex, '120101199810017032') is not None,  # True
        re.search(regex, '12010119981001701X') is not None,  # True
        re.search(regex, '12010119981001701x') is not None,  # True
        re.search(regex, '120101199810092') is not None,  # True
        re.search(regex, '22010119981009x') is not None,  # False
        re.search(regex, '32010119981009X') is not None,  # False
    ]
    [print(items[i]) for i in range(len(items))]


def solution2() -> None:
    # 匹配 IP 地址
    regex = r'^((00)?[0-9]|0?[0-9]{2}|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.((00)?[0-9]|0?[0-9]{2}|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.((00)?[0-9]|0?[0-9]{2}|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.((00)?[0-9]|0?[0-9]{2}|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$'
    items = [
        re.search(regex, '0.0.0.0') is not None,  # True
        re.search(regex, '0.255.255.255') is not None,  # True
        re.search(regex, '192.168.1.1') is not None,  # True
        re.search(regex, '255.255.255.255') is not None,  # True
        re.search(regex, '127.0.0.1') is not None,  # True
        re.search(regex, '127.0.0.256') is not None,  # True
        re.search(regex, '003.010.011.005') is not None,  # True
        re.search(regex, '03.010.001.5') is not None,  # True
    ]
    [print(items[i]) for i in range(len(items))]


def solution3() -> None:
    # 精准匹配 HTML Tag
    regex = r"^<('[^']*'|\"[^\"]*\"|[^'\">])+>$"
    items = [
        re.search(regex, "<input name=txt value=\">\">") is not None,  # True
        re.search(regex, "<input name=txt value='>'>") is not None,  # True
        re.search(regex, '<a href="user-info" class="cls-user-href">')
        is not None,  # True
        re.search(regex, '<table style="color:#eee">') is not None,  # True
    ]
    [print(items[i]) for i in range(len(items))]


def solution4() -> None:
    regex1 = r'^ab|cd$'
    regex2 = r'^(ab|cd)$'
    regex3 = r'(^a|b|c)'
    items = [
        re.search(regex1, 'abc') is not None,  # True
        re.search(regex1, 'bcd') is not None,  # True
        re.search(regex1, 'ab') is not None,  # True
        re.search(regex1, 'cd') is not None,  # True
        re.search(regex1, 'fab') is not None,  # False
        re.search(regex1, 'cde') is not None,  # False
        re.search(regex1, 'abcd') is not None,  # True
        re.search(regex2, 'abc') is not None,  # False
        re.search(regex2, 'bcd') is not None,  # False
        re.search(regex2, 'ab') is not None,  # True
        re.search(regex2, 'cd') is not None,  # True
        re.search(regex2, 'abcd') is not None,  # False
        re.search(regex3, 'ab') is not None,  # True
        re.search(regex3, 'cd') is not None,  # True
        re.search(regex3, 'abc') is not None,  # True
        re.search(regex3, 'bcd') is not None,  # True
        re.search(regex3, 'abcd') is not None,  # True
        re.search(regex3, 'abcde') is not None,  # True
        re.search(regex3, 'bcdef') is not None,  # True
        re.search(regex3, 'cdefg') is not None,  # True
        re.search(regex3, 'defgh') is not None,  # False
    ]
    [print(items[i]) for i in range(len(items))]


if __name__ == "__main__":
    solution()
    solution2()
    solution3()
    solution4()
