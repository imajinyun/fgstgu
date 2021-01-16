# coding=utf-8

from typing import List


def solution() -> List[int]:
    """
    1. 问题描述：

    一辆卡车违反交通规则，撞人后逃跑。现场有三人目击该事件，但都没有记住车号，只记下了车号的一些特征。
    甲说：牌照的前两位数字是相同的；
    乙说：牌照的后两位数字是相同的，但与前两位不同；
    丙是数学家，他说：4位的车号刚好是一个整数的平方。请根据以上线索求出车号。

    2. 问题分析：

    按照题目的要求造出一个前两位数相同、后两位数相同且相互间又不同的4位整数，
    然后判断该整数是否是另一个整数的平方。即求一个 4 位数 a1a2a3a4，满足如下条件：

    * a1 = a2 (1<=a1<=9, 0<=a2<=9)
    * a3 = a4 (a3>=0, a4<=9)
    * a1 != a3
    * 1000*a1 + 100 *a2 + 10 * a3 + a4 = x^2 (x∈Z)
    """
    res, flag = [], False
    for i in range(10):
        if flag:
            break
        for j in range(10):
            if flag:
                break
            if i != j:
                k = 1000 * i + 100 * i + 10 * j + j
                # 车牌号为四位数，需限定范围。
                for t in range(31, 100):
                    if t * t == k:
                        res.append(k)
                        flag = True
                        break
    return res


if __name__ == "__main__":
    print(solution())
