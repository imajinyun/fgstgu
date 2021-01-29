# coding=utf-8

import re


def solution() -> None:
    regex = r'^[-\w.]{1,64}@([-a-zA-Z0-9]{1,63}\.)*[-a-zA-Z0-9]{1,63}$'
    items = [
        re.search(regex, 'abc@somehost') is not None,  # True
        re.search(regex, 'abc@somehost.com') is not None,  # True
        re.search(regex, 'abc@some-host.com') is not None,  # True
        re.search(regex, '123@some-host.info') is not None,  # True
        re.search(regex, 'abc123@sum.some-host.io') is not None,  # True
        re.search(regex, 'abc-123@sum.some-123host.info') is not None,  # True
        re.search(regex, 'abc-123_@sum.some-123host.info') is not None,  # True
        re.search(regex, 'abc-123_@sum.1-123host.info') is not None,  # True
        re.search(regex, 'a-1_@a.1-2tX9.m.com') is not None,  # True
        re.search(regex, 'a-1_@a.1-2tX9.m.com.') is not None,  # False
        re.search(regex, 'a-1_@a.1-2tX9.sing.cm.cn') is not None,  # True
        re.search(regex, 'a--1__a..1--2@test.com.cn') is not None,  # True
        re.search(regex, '--1__a..1--2@test.com.cn') is not None,  # True
        re.search(regex, '..1__a..1--2@test.com.cn') is not None,  # True
        re.search(regex, '1_a.1.2..@test.com.cn') is not None,  # True
        re.search(regex, '1_a.1.2---@test.com.cn') is not None,  # True
        re.search(regex, 'a' * 65 + '@test.com.cn') is not None,  # False
    ]
    [print(items[i]) for i in range(len(items))]


def solution2() -> None:
    """
    把『第一个字符不能是点号或横线』也用环视表达，整个表达式就是：

    r'\A(?![-.])([-\w.](?!\.\.)){1,64}@([-a-zA-Z0-9]{1,63}\.)*[-a-zA-Z0-9]{1,63}\Z'
    """
    regex = r'\A\w([-\w.](?!\.\.)){0,63}@([-a-zA-Z0-9]{1,63}\.)*[-a-zA-Z0-9]{1,63}\Z'
    items = [
        re.search(regex, 'abc@somehost') is not None,  # True
        re.search(regex, 'abc@somehost.com') is not None,  # True
        re.search(regex, 'abc@some-host.com') is not None,  # True
        re.search(regex, '123@some-host.info') is not None,  # True
        re.search(regex, 'abc123@sum.some-host.io') is not None,  # True
        re.search(regex, 'abc-123@sum.some-123host.info') is not None,  # True
        re.search(regex, 'abc-123_@sum.some-123host.info') is not None,  # True
        re.search(regex, 'abc-123_@sum.1-123host.info') is not None,  # True
        re.search(regex, 'a-1_@a.1-2tX9.m.com') is not None,  # True
        re.search(regex, '1_a.1.2__@test.com.cn') is not None,  # True
        re.search(regex, '1_a.1.2--@test.com.cn') is not None,  # True
        re.search(regex, 'a-1_@a.1-2tX9.sing.cm.cn') is not None,  # True
        re.search(regex, 'a--1_a.1--2@test.com.cn') is not None,  # True
        re.search(regex, '1_a.1.2.-_@test.com.cn') is not None,  # True
        re.search(regex, '_1_a.1.2.-_@test.com.cn') is not None,  # True
        re.search(regex, '__1_a.1.2.-_@test.com.cn') is not None,  # True
        re.search(regex, '__1_a.1.2.-__@test.com.cn') is not None,  # True
        re.search(regex, '____func__tion____@test.com.cn') is not None,  # True
        re.search(regex, '_-.f-._t_.-@test.com.cn') is not None,  # True
        re.search(regex, '_f-._t_.--@test.com.cn') is not None,  # True
        re.search(regex, 'a-1_@a.1-2tX9.m.com.') is not None,  # False
        re.search(regex, 'a--1__a..1--2@test.com.cn') is not None,  # False
        re.search(regex, '--1__a..1--2@test.com.cn') is not None,  # False
        re.search(regex, '..1__a..1--2@test.com.cn') is not None,  # False
        re.search(regex, '1_a.1.2..@test.com.cn') is not None,  # False
        re.search(regex, '.1_a.1.2-@test.com.cn') is not None,  # False
        re.search(regex, '..1_a.1.2@test.com.cn') is not None,  # False
        re.search(regex, '--1_a.1.2@test.com.cn') is not None,  # False
        re.search(regex, 'a' * 65 + '@test.com.cn') is not None,  # False
    ]
    [print(items[i]) for i in range(len(items))]


if __name__ == "__main__":
    solution()
    print('\n')
    solution2()
