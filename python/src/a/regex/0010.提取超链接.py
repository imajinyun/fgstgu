# coding=utf-8

import re
import urllib.request


def solution() -> None:
    url = 'https://baidu.com/'
    with urllib.request.urlopen(url) as request:
        response = request.read().decode('utf-8')
        request.close()

        print("All href links:")
        # 粗略版本：r'<a\s[\s\S]+?</a>'
        regex = r"<a(?=\s)[^>]*(?<=\s)href\s*=\s*[\"']?([^\"'\s]+)[\"']?[^>]*>([\s\S]+?)</a>"
        for href in re.findall(regex, response):
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
