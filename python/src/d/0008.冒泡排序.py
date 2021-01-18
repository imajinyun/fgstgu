# coding=utf-8

from typing import List


def solution(item: List[int]) -> List[int]:
    """
    1. 问题描述：

    对 N 个整数（数据由键盘输入）进行升序排列。

    2. 问题分析：

    a. 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
    b. 对每一对相邻元素做同样的工作，从开始第一对到结尾的最后一对。在这一点，最后的元素应该会是最大的数。
    c. 针对所有的元素重复以上的步骤，除了最后一个。
    d. 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
    """
    n = len(item)
    for i in range(n):
        for j in range(n - i - 1):
            if item[j] > item[j + 1]:
                t = item[j]
                item[j] = item[j + 1]
                item[j + 1] = t
    return item


if __name__ == '__main__':
    print('请输入列表元素（列表末尾不要输入空格）：')
    x = input().split(' ')
    x = [int(x[v]) for _, v in enumerate(range(len(x)))]
    print('你输入的列表元素为：', x)
    print('冒泡排序后的列表为：', solution(x))
