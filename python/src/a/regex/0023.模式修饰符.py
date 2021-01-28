# coding=utf-8

import re
import urllib.request


def solution() -> None:
    regex = r'(?i)the'
    items = [
        re.search(regex, 'The') is not None,  # True
        re.search(regex, 'the') is not None,  # True
        re.search(regex, 'THE') is not None,  # True
        re.compile(r'the', re.I).search("ThE") is not None,  # True
        re.search(r'the', 'ThE', re.I) is not None,  # True
    ]
    [print(items[i]) for i in range(len(items))]


def solution2() -> None:
    with urllib.request.urlopen('https://jd.com/') as request:
        response = request.read().decode('utf-8')
        request.close()

        print('\n')
        regex1 = r'(?i)<a\s+href\s*=\s*[\'"]?([^\'"\s]+)[\'"]?>([^<]+)</a>'
        regex2 = r'(?i)<title>([^>]+)</title>'
        regex3 = r'(?i)<img\s[^>]*?src=[\'"]?([^"\']+)[\'"]?[^>]*>'
        [print('link ->', href) for href in re.findall(regex1, response)]
        [print('title ->', title) for title in re.findall(regex2, response)]
        [print('img ->', img) for img in re.findall(regex3, response)]


def solution3() -> None:
    # 单行模式
    with urllib.request.urlopen('https://dumall.baidu.com/') as request:
        response = request.read().decode('utf-8')
        request.close()

        print('\n')
        regex = r'(?s)<script\s.*?</script>'
        [print('script ->', script) for script in re.findall(regex, response)]


def solution4() -> None:
    """
1 line
first
2 line
second
3 line
third
    """
    # 多行模式
    print('\n多行模式：')
    regex = r'(?m)^\d.*'
    multi = '1 line\nfirst\n2 line\nsecond\n3 line\nthird'
    print(re.findall(regex, solution4.__doc__))
    [print('text ->', text) for text in re.findall(regex, solution4.__doc__)]
    print('\n多行模式匹配后加 . 号：')
    print(re.sub(r'(?m)$', '.', multi))


def solution5():
    # 注释模式
    multi = '1 line\nfirst\n2 line\nsecond\n3 line\nthird'
    regex = r"""(?mx)  # enable multi-line and extended mode
    ^                  # start of the line
    \d                 # digit
    .*                 # rest of the line
    $                  # end of the line
    """
    print('\n注释模式：')
    print(re.findall(regex, multi))


if __name__ == '__main__':
    solution()
    solution2()
    solution3()
    solution4()
    solution5()
