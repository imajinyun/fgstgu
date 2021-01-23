# coding=utf-8

import re
import urllib.request


def solution() -> None:
    regex = r'(\d{4})-(\d{2})-(\d{2})'
    items = [
        re.search(regex, '2021-01-22').group(),  # 2021-01-22
        re.search(regex, '2021-01-22').group(0),  # 2021-01-22
        re.search(regex, '2021-01-22').group(1),  # 2021
        re.search(regex, '2021-01-22').group(2),  # 01
        re.search(regex, '2021-01-22').group(3),  # 22
    ]
    [print(items[i]) for i in range(len(items))]


def solution2() -> None:
    """
      (((\d{4})-(\d{2}))-(\d{2}))
    0 |**_____*_*______*_*______*|
    1 |**_____*_*______*_*______|
    2  |*_____*_*______| *      *
    3   |_____| *      * *      *
    4           |______| *      *
    5                    |______|
    """
    regex = r'(((\d{4})-(\d{2}))-(\d{2}))'
    items = [
        re.search(regex, '2021-01-22').group(),  # 2021-01-22
        re.search(regex, '2021-01-22').group(0),  # 2021-01-22
        re.search(regex, '2021-01-22').group(1),  # 2021-01-22
        re.search(regex, '2021-01-22').group(2),  # 2021-01
        re.search(regex, '2021-01-22').group(3),  # 2021
        re.search(regex, '2021-01-22').group(4),  # 01
        re.search(regex, '2021-01-22').group(5),  # 22
    ]
    [print(items[i]) for i in range(len(items))]


def solution3() -> None:
    with urllib.request.urlopen("https://jd.com/") as request:
        response = request.read().decode('utf-8')
        request.close()

        regex = r"<a\s+href\s*=\s*[\"']?([^\"'\s]+)[\"']?>([^<]+)</a>"
        for item in re.findall(regex, response):
            print(item[1], '->', item[0])


def solution4() -> None:
    with urllib.request.urlopen("https://baidu.com") as request:
        response = request.read().decode('utf-8')
        request.close()

        regex = r"<img\s+[^>]*?src=['\"]?([^\"'\s]+)['\"]?[^>]*>"
        [print(item) for item in re.findall(regex, response)]


if __name__ == '__main__':
    solution()
    solution2()
    solution3()
    solution4()
