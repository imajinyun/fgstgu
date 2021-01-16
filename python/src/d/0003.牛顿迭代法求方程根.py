# coding=utf-8
from typing import List


def solution(a: List[int], b: List[int], c: List[int], d: List[int]) -> List[float]:
    """
    1. 问题描述：

    编写用牛顿迭代法求方程根的函数。方程为 ax^3+bx^2+cx+d=0，系数 a、b、c、d 由主函数输入，
    求 x 在 1 附近的一个实根。求出根后，由主函数输出。
    牛顿迭代法的公式：x=x0-(f(x0)/f`(x0))，设迭代到 |x-x0|≤10^-5 时结束。
    求解一元三次方程 ax^3+bx^2+cx+d=0 的解。

    2. 问题分析：

    牛顿迭代法是取 x0 之后，在这个基础上找到比 x0 更接近的方程根，一步一步迭代，
    从而找到更接近方程根的近似根。设 r 是 f(x)=0 的根，选取 x0 作为 r 的初始近似值，
    过点 (x0,f(x0)) 做曲线 y=f(x) 的切线 L，L 的方程为 y=f(x0)+f'(x0)(x-x0)，
    则 L 与 x 轴交点的横坐标 x1=x0-f(x0)/f'(x0)，称 x1 为 r 的一次近似值；
    过点 (x1,f(x2)) 做曲线 y=f(x) 的切线，并求出该切线与 x 轴交点的横坐标 x2=x1-f(x1)/f'(x1)，
    称 x2 为 r 的二次近似值；重复以上过程，得到 r 的近似值序列，其中，xn+1=xn-f(xn)/f'(xn) 称为 r 的 n+1 次近似值。
    上述过程即为牛顿迭代法的求解过程。

    3. 算法设计：

    a. 在 1 附近找任一实数作为 x0 的初值，我们取 1.5，即 x0=1.5；
    b. 用初值 x0 代入方程中计算此时的 f(x0) 及 f'(x0)；程序中用变量 f 描述方程的值，用 fd 描述方程求导之后的值；
    c. 计算增量 y=f/fd。
    d. 计算下一个 x，x=x0-h；
    e. 用新产生的x替换原来的 x0，为下一次迭代做好准备；
    f. 若 |x-x0|＞=1e-5，则转到步骤 c 继续执行，否则转到步骤 g；
    g. 所求 x 就是方程 ax3+bx2+cx+d=0 的根，将其输出；
    """

    def f(x: List[float]) -> List[float]:
        return a * x * x * x + b * x * x + c * x + d

    def fd(x: List[float]) -> List[float]:
        return 3 * a * x * x + 2 * b * x + c

    x, e = 1.5, 1e-5
    x0 = x
    y = f(x0) / fd(x0)
    x = x0 - y
    while abs(x - x0) >= e:
        x0 = x
        y = f(x0) / fd(x0)
        x = x0 - y
    return x


if __name__ == '__main__':
    print("请输入方程的系数：")
    a, b, c, d = map(float, input().split())
    print("输入的方程系数为：", a, b, c, d)
    x = solution(a, b, c, d)
    print("所求方程的根为 x = %.6f" % x)
