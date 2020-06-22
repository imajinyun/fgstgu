import time
import datetime

"""
字符串类型的日期转换为时间戳
"""
times = time.strptime('2020-05-25 23:59:59', '%Y-%m-%d %H:%M:%S')

# time.struct_time(tm_year=2020, tm_mon=5, tm_mday=25, tm_hour=23, tm_min=59, tm_sec=59, tm_wday=0, tm_yday=146, tm_isdst=-1)
print(times)

stamp = int(time.mktime(times))

# 1588694399
print(stamp)


"""
更改字符串类型日期的显示格式
"""
times = time.strptime('2020-05-25 23:59:59', '%Y-%m-%d %H:%M:%S')

# 2020/05/25 23:59:59
print(time.strftime('%Y/%m/%d %H:%M:%S', times))
print('-' * 80)


"""
时间戳转换为指定格式的日期
"""
stamp = 1590422399
times = time.localtime(stamp)

# 2020-05-25 23:59:59
print(time.strftime('%Y-%m-%d %H:%M:%S', times))

dates = datetime.datetime.fromtimestamp(stamp)

# 2020-05-25 23:59:59
print(dates.strftime('%Y-%m-%d %H:%M:%S'))

dates = datetime.datetime.utcfromtimestamp(stamp)

# 2020-05-25 15:59:59
print(dates.strftime('%Y-%m-%d %H:%M:%S'))
print('-' * 80)


"""
获取当前时间并且用指定格式显示
"""
now = int(time.time())
times = time.localtime(now)

# time.struct_time(tm_year=2020, tm_mon=6, tm_mday=22, tm_hour=21, tm_min=34, tm_sec=4, tm_wday=0, tm_yday=174, tm_isdst=0)
print(times)

# 2020-06-22 21:35:37
print(time.strftime('%Y-%m-%d %H:%M:%S', times))

now = datetime.datetime.now()

# 2020-06-22 21:36:23.021165
print(now)

# 2020-06-22 21:36:23
print(now.strftime('%Y-%m-%d %H:%M:%S'))
