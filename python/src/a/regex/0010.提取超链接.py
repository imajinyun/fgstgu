# coding=utf-8

import re
import urllib.request


def solution() -> None:
    url = 'https://baidu.com/'
    with urllib.request.urlopen(url) as request:
        response = request.read().decode('utf-8')
        request.close()

        print("All href links:")
        for href in re.findall(r'<a\s[\s\S]+?</a>', response):
            print(href)


def solution2() -> None:
    url = 'https://baidu.com'
    with urllib.request.urlopen(url) as request:
        response = request.read().decode('utf-8')
        request.close()

        print("\nAll script codes:")
        for code in re.findall(r'<script[\s>][\s\S]+?</script>', response):
            print(code)


if __name__ == "__main__":
    solution()
    solution2()
