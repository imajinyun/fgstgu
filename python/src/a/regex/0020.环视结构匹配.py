# coding=utf-8

import re


def solution() -> None:
    """
    环视一共分为四种：

    | 名称 | 记法 | 方向 | 结构内表达式匹配成功的返回值 |
    | --- | --- | --- | --- |
    | 肯定顺序环视（positive-lookahead） | (?=...) | 向右 | True |
    | 否定顺序环视（negative-lookahead） | (?!...) | 向右 | False |
    | 肯定逆序环视（positive-lookbehind）| (?<=...) | 向左 | True |
    | 否定逆序环视（negative-lookbehind）| (?<!...) | 向左 | False |

    四种环视匹配的位置：
                |1|2|3|4|5|
    (?=\d{3})   ⇡ ⇡ ⇡
    (?!\d{3})         ⇡ ⇡ ⇡
    (?<=\d{3})        ⇡ ⇡ ⇡
    (?<!\d{3})  ⇡ ⇡ ⇡

    对于字符串 12345，以 \d{3} 为表达式的四种环视能匹配的位置分别是：
    右侧必须出现三个数字字符，右侧不能出现三个数字字符，左侧必须出现三个数字字符，左侧不能出现三个数字字符。
    """
    regex = r"\A<(?!/)('[^']*'|\"[^\"]*\"|[^'\">])+(?<!/)>\Z"
    input1 = '<input name="text" value=">">'
    input2 = '<input name="text" value=\'>\'>'
    input3 = '<input name="text" value="">'
    items = [
        re.search(regex, input1) is not None,  # True
        re.search(regex, input2) is not None,  # True
        re.search(regex, input3) is not None,  # True
        re.search(regex, '<br/>') is not None,  # False
        re.search(regex, '<i>') is not None,  # True
        re.search(regex, '<img src="url"/>') is not None,  # False
    ]
    print(items)


def solution2() -> None:
    # 格式化数字字符串

    # -> 123,456
    print(re.sub(r'(?<=\d)(?=(\d{3})+(?!\d))', r',', '123456'))

    # -> 1,234,567,890
    print(re.sub(r'(?<=\d)(?=(\d{3})+(?!\d))', r',', '1234567890'))

    # 去掉文本中不必要的空白字符
    text = '中  英  文，this is a test， 有 多 余的空白 符 '

    # -> 中英文，this is a test，有多余的空白符
    print(re.sub(r'(?<![a-zA-Z])\s+(?![a-zA-Z])', '', text))


def solution3() -> None:
    text = '  中  英  文，this is a test， 有 多 余的空白 符  '

    # -> [中英文，this is a test，有多余的空白符]
    print('[' + re.sub(r'(?<![a-zA-Z])\s+(?![a-zA-Z])', '', text) + ']')

    # -> [ 中英文，this is a test，有多余的空白符 ]
    print('[' + re.sub(r'(?<=[^a-zA-Z])\s+(?=[^a-zA-Z])', '', text) + ']')


def solution4() -> None:
    # 准确匹配主机名
    regex = r'^(?=[-a-zA-Z0-9.]{0,255}(?![-a-zA-Z0-9.]))((?!-)[-a-zA-Z0-9]{1,63}\.)*((?!-)[-a-zA-Z0-9]){1,63}$'
    items = [
        re.search(regex, 'localhost') is not None,  # True
        re.search(regex, 'example.com') is not None,  # True
        re.search(regex, 'www.baidu.com') is not None,  # True
        re.search(regex, '-tmall.com') is not None,  # False
        re.search(regex, ('a' * 64) + '.com') is not None,  # False
        re.search(regex, ('k' * 256) + '.org') is not None,  # False
    ]
    [print(items[i]) for i in range(len(items))]


if __name__ == '__main__':
    solution()
    solution2()
    solution3()
    solution4()
