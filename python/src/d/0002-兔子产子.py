# coding=utf-8

def solution(n: int) -> int:
    """
    1．问题描述：

    有一对兔子，从出生后的第3个月起每个月都生一对兔子。
    小兔子长到第3个月后每个月又生一对兔子，假设所有的兔子都不死，问30个月内每个月的兔子总对数为多少？

    2. 问题分析：

    月数  小兔子对数  中兔子对数  老兔子对数  兔子总对数
    1     1         0          0         1
    2     0         1          0         1
    3     1         0          1         2
    4     1         1          1         3
    5     2         1          2         5
    6     3         2          3         8
    7     5         3          5         13

    说明：不满 1 个月的兔子为小兔子，满 1 个月不满 2 个月的为中兔子，满 3 个月以上的为老兔子。
    """
    if n == 1 or n == 2:
        return 1
    else:
        return solution(n - 1) + solution(n - 2)


def solution2(n: int) -> int:
    # 每次计算两个，循环变量循环到原来的一半。
    res, tmp, i = 0, 1, 1
    while i <= (n // 2):
        res = res + tmp
        tmp = res + tmp
        i += 1
    return res


if __name__ == "__main__":
    print(solution(30))
    print(solution2(30))
