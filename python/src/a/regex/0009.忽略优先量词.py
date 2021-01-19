# coding=utf-8

import re
import urllib.request


def solution() -> None:
    with urllib.request.urlopen("https://jd.com/") as url:
        response = url.read().decode('utf-8')
        url.close()

        print("\nJavaScript codes:")
        regex = r'<script type="text/javascript">[\s\S]*?</script>'
        for code in re.findall(regex, response):
            print(code)


def solution2() -> None:
    url = 'https://raw.githubusercontent.com/begeekmyfriend/tash/master/tash.c'
    with urllib.request.urlopen(url) as request:
        response = request.read().decode('utf-8')
        request.close()

        print("\nComments extract(e.g //):")
        for comment in re.findall(r'//.*', response):
            print(comment)


def solution3() -> None:
    url = 'https://raw.githubusercontent.com/begeekmyfriend/tash/master/tash.c'
    with urllib.request.urlopen(url) as request:
        response = request.read().decode('utf-8')
        request.close()

        print("\nComments extrace(e.g /**/):")
        for comment in re.findall(r'/\*[\s\S]*?\*/', response):
            print(comment)


if __name__ == "__main__":
    solution()
    solution2()
    solution3()
