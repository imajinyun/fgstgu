# coding=utf-8

import re
import urllib.request


def solution() -> None:
    items = [
        # 通过 MatchObject 获得匹配的文本
        re.search(r'\d{6}', 'abc123456def').group(0),  # 123456
        re.search(r'\w{3}', '123abc456def').group(0),  # 123

        # 使用 re.findall() 提取数据
        re.findall(r'\d{6}',
                   "zipcode1 :100100, zipcode2:200200, zipcode3: 300300"
                   ),  # ['100100', '200200', '300300']
    ]
    [print(items[i]) for i in range(len(items))]


def solution2() -> None:
    with urllib.request.urlopen("https://tmall.com/") as url:
        response = url.read().decode('utf-8')
        url.close()

        print("\nOpen tags:")
        print(re.findall(r'<[^/>][^>]*[^/>]>', response)[0:6])

        print("\nClose tags:")
        print(re.findall(r'</[^>]+>', response)[0:6])

        print("\nSelf-closing tags:")
        print(re.findall(r'<[^>/]+/>', response)[0:6])


if __name__ == "__main__":
    solution()
    solution2()
