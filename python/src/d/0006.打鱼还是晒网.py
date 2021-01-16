# coding=utf-8
from typing import Dict


def solution(today: Dict[str, int], year: int) -> bool:
    """
    1. 问题描述：

    中国有句俗语叫『三天打鱼两天晒网』。某人从1990年1月1日起便开始『三天打鱼两天晒网』，
    问这个人在以后的某一天中是『打鱼』还是『晒网』。

    2. 问题分析：

    根据题意可以将解题过程分为以下三步：

    a. 计算从 1990 年 1 月 1 日开始至指定日期共有多少天；
    b. 由于『打鱼』和『晒网』的周期为 5 天，所以将计算出的天数用 5 去除；
    c. 根据余数判断他是在『打鱼』还是在『晒网』。若余数为 1，2，3，则他是在『打鱼』，否则是在『晒网』；
    """

    def is_leap_year(year: int) -> bool:
        if (year % 4 == 0 and year % 100 != 0) or year % 400 == 0:
            return True
        return False

    def total_days(today: Dict[str, int]) -> int:
        months = [0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
        total, year = 0, 1990
        while year < today['year']:
            total = (total + 366) if is_leap_year(year) else (total + 365)
            year += 1

        if is_leap_year(today['year']):
            months[2] += 1

        i = 0
        while i < today['month']:
            total += months[i]
            i += 1
        total += today['day']
        return total

    total = total_days(today)
    format = '%d年%d月%d日与%d年1月1日相差 %d 天'
    print(format % (today['year'], today['month'], today['day'], year, total))

    if 0 < (total % 5) < 4:
        return True
    return False


if __name__ == '__main__':
    while True:
        print("请输入指定日期（格式：2000 1 31）：")
        year, month, day = [int(i) for i in input().split()]
        if not (1990 <= year <= 9999):
            break
        if not (1 <= month <= 12):
            break
        if not (1 <= day <= 31):
            break
        today = {'year': year, 'month': month, 'day': day}
        if solution(today, 1990):
            print("今天去打鱼")
        else:
            print("今天去晒网")
