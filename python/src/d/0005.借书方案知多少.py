# coding=utf-8
from typing import List


def solution() -> List[int]:
    """
    1. 问题描述：

    小明有 5 本新书，要借给 A、B、C 三位小朋友，若每人每次只能借 1 本，则可以有多少种不同的借法？

    2. 问题分析：

    本题属于数学中常见的排列组合问题，即求从 5 个数中取 3 个不同数的排列组合的总数。
    我们可以将 5 本书进行 1～5 编号，A、B、C 三个人每次都可以从 5 本书中任选 1 本，
    即每人都有 5 种选择，由于 1 本书不可能同时借给一个以上的人，
    因此只要这三个人所选书的编号不同，则即为一次有效的借阅方法。
    """
    res = []
    for i in range(1, 6):
        for j in range(1, 6):
            k = 1
            while k < 6 and i != j:
                if i != k and j != k:
                    res.append([i, j, k])
                k += 1
    return res


if __name__ == "__main__":
    res = solution()
    print(len(res))
    print(res)
