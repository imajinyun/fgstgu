# coding=utf-8

import re


def solution() -> None:
    regex1 = r'^[1-9]\d{13,16}[0-9xX]$'
    regex2 = r'^[1-9]\d{14}(\d{2}[0-9xX])?$'
    items = [
        # 不精确匹配
        re.search(regex1, '120101199810017032') is not None,  # True
        re.search(regex1, '120101199810017') is not None,  # True
        re.search(regex1, '12010119981001701X') is not None,  # True
        re.search(regex1, '12010120081001701x') is not None,  # True

        # 错误的匹配
        re.search(regex1, '329188200612289') is not None,  # True
        re.search(regex1, '32918820061228X') is not None,  # True

        # 精确的匹配
        re.search(regex2, '120101199810017032') is not None,  # True
        re.search(regex2, '120101199810018') is not None,  # True
        re.search(regex2, '12010119981001X') is not None,  # False
        re.search(regex2, '12010119981001502X') is not None,  # True
        re.search(regex2, '12010119981001502x') is not None,  # True
        re.search(regex2, '1201011998100152') is not None,  # False
        re.search(regex2, '12010119981001523') is not None,  # False

        # 其它
        re.search(r'^ab+$', 'ab') is not None,  # True
        re.search(r'^ab+$', 'abb') is not None,  # True
        re.search(r'^ab+$', 'abab') is not None,  # False
        re.search(r'^(ab)+$', 'ab') is not None,  # True
        re.search(r'^(ab)+$', 'abb') is not None,  # False
        re.search(r'^(ab)+$', 'abab') is not None,  # True
    ]
    for i in range(len(items)):
        print(items[i])


def solution2() -> None:
    regex = r'^<[^/]([^>]*[^/])?>$'
    items = [
        # Open tag 精确匹配
        re.search(regex, '<table>') is not None,  # True
        re.search(regex, '<i>') is not None,  # True
        re.search(regex, '</i>') is not None,  # False
        re.search(regex, '</table>') is not None,  # False
        re.search(regex, '<br />') is not None,  # False
    ]
    for i in range(len(items)):
        print(items[i])


def solution3() -> None:
    regex = r'^/[a-z]+(/[a-z]+(_[a-z]+)?\.php)?$'
    items = [
        # 路由匹配
        re.search(regex, '/foo') is not None,  # True
        re.search(regex, '/foo/bar.php') is not None,  # True
        re.search(regex, '/foo/bar_coo.php') is not None,  # True
        re.search(regex, '/foo/_') is not None,  # False
        re.search(regex, '/foo.php') is not None,  # False
    ]
    for i in range(len(items)):
        print(items[i])


if __name__ == "__main__":
    solution()
    solution2()
    solution3()
