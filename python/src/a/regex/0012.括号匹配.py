# coding=utf-8

import re


def solution() -> None:
    regex1 = r'^[1-9]\d{13,16}[0-9xX]$'
    regex2 = r'^[1-9]\d{14}(\d{2}[0-9xX])?$'
    regex3 = r'^ab+$'
    regex4 = r'^(ab)+$'
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
        re.search(regex3, 'ab') is not None,  # True
        re.search(regex3, 'abb') is not None,  # True
        re.search(regex3, 'abab') is not None,  # False
        re.search(regex4, 'ab') is not None,  # True
        re.search(regex4, 'abb') is not None,  # False
        re.search(regex4, 'abab') is not None,  # True
    ]
    [print(items[i]) for i in range(len(items))]


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
    [print(items[i]) for i in range(len(items))]


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
    [print(items[i]) for i in range(len(items))]


def solution4() -> None:
    regex = r'^[-\w.]{1,64}@([-a-zA-Z0-9]{1,63}\.)*[-a-zA-Z0-9]{1,63}$'
    items = [
        re.search(regex, 'abc@somehost') is not None,  # True
        re.search(regex, 'abc@somehost.com') is not None,  # True
        re.search(regex, 'abc@some-host.com') is not None,  # True
        re.search(regex, '123@some-host.info') is not None,  # True
        re.search(regex, 'abc123@sum.some-host.io') is not None,  # True
        re.search(regex, 'abc-123@sum.some-123host.info') is not None,  # True
        re.search(regex, 'abc-123@sum.some-123host.info') is not None,  # True
        re.search(regex, 'abc-123_@sum.some-123host.info') is not None,  # True
        re.search(regex, 'abc-123_@sum.1-123host.info') is not None,  # True
        re.search(regex, 'a-1_@a.1-2tX9.m.com') is not None,  # True
        re.search(regex, 'a-1_@a.1-2tX9.sing.cm.cn') is not None,  # True
    ]
    [print(items[i]) for i in range(len(items))]


if __name__ == "__main__":
    solution()
    solution2()
    solution3()
    solution4()
