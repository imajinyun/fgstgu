from itertools import groupby
from operator import itemgetter

lists = [
    ('2012-05-21', 11),
    ('2012-05-21', 3),
    ('2012-05-22', 10),
    ('2012-05-22', 4),
    ('2012-05-22', 22),
    ('2012-05-23', 33),
]

"""
2012-05-21
  ('2012-05-21', 11)
  ('2012-05-21', 3)
2012-05-22
  ('2012-05-22', 10)
  ('2012-05-22', 4)
  ('2012-05-22', 22)
2012-05-23
  ('2012-05-23', 33)
"""
for key, items in groupby(lists, itemgetter(0)):
    print(key)
    for item in items:
        print(' ', item)

print('-' * 80)

rows = [
    {'code': '13123', 'date': '2021-09-10'},
    {'code': '23423', 'date': '2020-06-14'},
    {'code': '43623', 'date': '2018-01-23'},
    {'code': '73246', 'date': '2019-03-11'},
    {'code': '51123', 'date': '2020-06-14'},
    {'code': '33721', 'date': '2021-09-10'},
]

from operator import itemgetter
from itertools import groupby

rows.sort(key=itemgetter('date'))

"""
2018-01-23
  {'code': '43623', 'date': '2018-01-23'}
2019-03-11
  {'code': '73246', 'date': '2019-03-11'}
2020-06-14
  {'code': '23423', 'date': '2020-06-14'}
  {'code': '51123', 'date': '2020-06-14'}
2021-09-10
  {'code': '13123', 'date': '2021-09-10'}
  {'code': '33721', 'date': '2021-09-10'}
"""
for date, items in groupby(rows, key=itemgetter('date')):
    print(date)
    for item in items:
        print(' ', item)

print('-' * 80)

from collections import defaultdict
rows_by_date = defaultdict(list)
for row in rows:
    rows_by_date[row['date']].append(row)

"""
2018-01-23
2019-03-11
2020-06-14
2021-09-10
"""
for row in rows_by_date:
    print(row)

print('-' * 80)

"""
{'code': '23423', 'date': '2020-06-14'}
{'code': '51123', 'date': '2020-06-14'}
"""
for row in rows_by_date['2020-06-14']:
    print(row)
