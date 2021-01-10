import csv
import os
import sys
from datetime import date, datetime


def date_diff(a, b):
    try:
        diff = str(datetime.strptime(a, '%Y/%m/%d') -
                   datetime.strptime(b, '%Y/%m/%d')).split()[0]
    except:
        diff = 0
    if diff == '0:00:00':
        diff = 0
    return diff


reader_file, writer_file = sys.argv[1], sys.argv[2]
data, prev_name, prev_catetory, prev_date = {}, 'N/A', 'N/A', 'N/A'
first = True
today = date.today().strftime('%Y/%m/%d')
with open(reader_file, 'r', newline='') as csv_reader:
    reader = csv.reader(csv_reader)
    headers = next(csv_reader)
    counter = 0
    for row in reader:
        counter += 1
        curr_name, curr_category, curr_date = row[0], row[1], row[3]
        if curr_name not in data:
            data[curr_name] = {}
        if curr_category not in data[curr_name]:
            data[curr_name][curr_category] = 0
        if curr_name != prev_name:
            if first:
                first = False
            else:
                diff = date_diff(today, prev_date)
                if prev_catetory not in data[prev_name]:
                    data[prev_name][prev_catetory] = int(diff)
                else:
                    data[prev_name][prev_catetory] += int(diff)
        else:
            diff = date_diff(curr_date, prev_date)
            data[prev_name][prev_catetory] += int(diff)
        prev_name, prev_catetory, prev_date = curr_name, curr_category, curr_date
    if counter > 0:
        diff = date_diff(today, prev_date)
        if prev_catetory not in data[prev_name]:
            data[prev_name][prev_catetory] = int(diff)
        else:
            data[prev_name][prev_catetory] += int(diff)

headers = ['客户名称', '服务分类', '总时间（按天）']
with open(writer_file, 'w', newline='') as csv_writer:
    writer = csv.writer(csv_writer)
    writer.writerow(headers)
    print(headers[0], '\t',  headers[1], '\t', headers[2])
    for name, item in data.items():
        for category, date in data[name].items():
            items = []
            print(name, '\t', category, '\t', date)
            items.append(name)
            items.append(category)
            items.append(date)
            writer.writerow(items)
