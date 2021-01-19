# coding=utf-8

from typing import List


def solution(x: List[int], m: int) -> int:
    """
    1. 问题描述：

    N 个有序整数数列已放在一维数组中，利用二分查找法查找整数 m 在数组中的位置。
    若找到，则输出其下标值；反之，则输出 -1。

    2. 问题分析：

    二分查找法（也叫折半查找）其本质是分治算法的一种，所谓分治算法指的是分而治之，
    即将较大规模的问题分解成几个较小规模的问题，这些子问题互相独立且与原问题相同，
    通过对较小规模问题的求解达到对整个问题的求解。
    当我们将问题分解成两个较小问题求解时的分治方法称为二分法。
    需要注意的是，二分查找法只适用于有序序列。

    二分查找法的基本算法是：每次查找前先确定数组中待查的范围。
    假设指针 low 和 high(low<high) 分别指示待查范围的下界和上界，
    指针 mid 指示待查范围的中间位置，即 mid=(low+high)/2，把 m 与中间位置（mid）上元素的值进行比较。
    如果 m 的值大于中间位置上元素的值，则下一次的查找范围放在中间位置之后的元素中；
    反之，下一次的查找范围放在中间位置之前的元素中。直到 low>high，查找结束。
    """
    low, high, index = 0, len(x) - 1, -1
    while low <= high:
        mid = low + (high - low) // 2
        if m < x[mid]:
            high = mid - 1
        else:
            if m > x[mid]:
                low = mid + 1
            else:
                index = mid
                break
    return -1 if index < 0 else index


if __name__ == '__main__':
    x = [
        0, 1, 2, 5, 11, 17, 20, 23, 29, 33, 33, 89, 98, 99, 99, 100, 100, 107,
        108, 1024, 2048, 9892
    ]
    m = int(input("Enter your find value: "))
    i = solution(x, m)
    print("Index:", i)
    print("Value:", "None" if i < 0 else x[i])
