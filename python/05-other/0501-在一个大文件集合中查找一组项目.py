import os
import sys
import glob
import csv
from datetime import date
import time
from xlrd import open_workbook, xldate_as_tuple
from xlwt import Workbook

reader_file, path, writer_file = sys.argv[1], sys.argv[2], sys.argv[3]
target = []
with open(reader_file, 'r', newline='') as csv_reader:
    reader = csv.reader(csv_reader)
    for row in reader:
        target.append(row[0])

writer = csv.writer(open(writer_file, 'a', newline=''))
file_counter, line_counter, item_numbers = 0, 0, 0
for reader_file in glob.glob(os.path.join(path, '*.*')):
    file_counter += 1
    if reader_file.split('.')[1] == 'csv':
        with open(reader_file, 'r', newline='') as csv_reader:
            reader = csv.reader(csv_reader)
            headers = next(reader)
            for row in reader:
                items = []
                for col in range(len(headers)):
                    if col == 3:
                        items.append(str(row[col]).lstrip('￥').strip())
                    elif col == 4:
                        arr = time.strptime(row[col], "%Y/%m/%d")
                        items.append(time.strftime('%Y-%m-%d', arr))
                    else:
                        items.append(row[col])
                items.append(os.path.basename(reader_file))
                if row[0] in target:
                    writer.writerow(items)
                    item_numbers += 1
                line_counter += 1
    elif reader_file.split('.')[1] in ['xls', 'xlsx']:
        workbook = open_workbook(reader_file)
        for worksheet in workbook.sheets():
            try:
                headers = worksheet.row_values(0)
            except IndexError:
                pass
            for row in range(1, worksheet.nrows):
                items = []
                for col in range(len(headers)):
                    value = worksheet.cell_value(row, col)
                    if col == 0:
                        items.append(str(value).split('.')[0])
                    elif col == 3:
                        items.append(str(value).lstrip('￥').strip())
                    elif col == 4:
                        value = xldate_as_tuple(value, workbook.datemode)
                        value = date(*value[0:3]).strftime('%Y-%m-%d')
                        items.append(value)
                    else:
                        items.append(str(value).strip())
                basename = os.path.basename(reader_file)
                items.append(basename + '@' + worksheet.name)
                tmp = str(worksheet.cell(row, 0).value).strip()
                if tmp.split('.')[0] in target:
                    writer.writerow(items)
                    item_numbers += 1
                line_counter += 1
print('Number of files: ', file_counter)
print('Number of lines: ', line_counter)
print('Number of item numbers: ', item_numbers)
