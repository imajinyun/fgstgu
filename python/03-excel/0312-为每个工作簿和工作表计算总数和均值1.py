import os
import sys
import glob
from datetime import date
from xlrd import open_workbook, xldate_as_tuple
from xlwt import Workbook

reader_file, writer_file = sys.argv[1], sys.argv[2]
workbook = Workbook()
worksheet = workbook.add_sheet('SumAndAverage')
data = []
index = 3
headers = ['workbook', 'worksheet', 'worksheet_total',
           'worksheet_average', 'workbook_total', 'workbook_average']
data.append(headers)
for reader_file in glob.glob(os.path.join(reader_file, '*.xlsx')):
    with open_workbook(reader_file) as book:
        totals, numbers, items = [], [], []
        for sheet in book.sheets():
            total_sales, number_sales, sheets = 0, 0, []
            sheets.append(os.path.basename(reader_file))
            sheets.append(sheet.nrows)
            for row in range(1, sheet.nrows):
                try:
                    value = sheet.cell_value(row, index)
                    total_sales += float(str(value).strip('￥'))
                    number_sales += 1
                except:
                    total_sales += 0.
                    number_sales += 0.
            average_sales = '%.2f' % (total_sales / number_sales)
            sheets.append(total_sales)
            sheets.append(float(average_sales))
            totals.append(total_sales)
            numbers.append(float(number_sales))
            items.append(sheets)
        workbook_total = sum(totals)
        workbook_average = sum(totals) / sum(numbers)
        for item in items:
            item.append(workbook_total)
            item.append(workbook_average)
        data.extend(items)
for key, value in enumerate(data):
    for k, v in enumerate(value):
        worksheet.write(key, k, v)
workbook.save(writer_file)
