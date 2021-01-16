# coding=utf-8
from typing import List


def solution() -> List[List[int]]:
    """
    1. 问题描述：

    中国古代数学家张丘建在他的《算经》中提出了一个著名的『百钱百鸡问题』：
    一只公鸡值五钱，一只母鸡值三钱，三只小鸡值一钱，
    现在要用百钱买百鸡，请问公鸡、母鸡、小鸡各多少只？

    2. 问题分析：

    用百钱如果只买公鸡，最多可以买 20 只，但题目要求买 100 只，由此可知，
    所买公鸡的数量肯定在 0～20 之间。同理，母鸡的数量在 0～33 之间。
    在此不妨把公鸡、母鸡和小鸡的数量分别设为 cock、hen、chicken，
    则 cock+hen+chicken=100，因此百钱买百鸡问题就转化成解不定方程组：

    {
        cock+hen+chicken=100
        5*cock+3*hen+chicken/3=100
    }
    """
    res = []
    for i in range(21):
        for j in range(34):
            k = 100 - j - i
            if 5 * i + 3 * j + k / 3.0 == 100:
                res.append([i, j, k])
    return res


if __name__ == '__main__':
    print(solution())
