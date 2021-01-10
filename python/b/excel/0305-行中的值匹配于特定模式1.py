import sys
import re
from datetime import date
from xlrd import open_workbook, xldate_as_tuple
from xlwt import Workbook

reader_file, writer_file = sys.argv[1], sys.argv[2]
pattern = re.compile(r'(?P<custom_pattern>^200-.*)')
workbook = Workbook()
worksheet = workbook.add_sheet('Sheet2018')
index = 2
with open_workbook(reader_file) as book:
    sheet = book.sheet_by_name('Sheet2018')
    data = []
    headers = sheet.row_values(0)
    data.append(headers)
    for row in range(1, sheet.nrows):
        items = []
        if pattern.search(sheet.cell_value(row, index)):
            for col in range(sheet.ncols):
                value = sheet.cell_value(row, col)
                if sheet.cell_type(row, col) == 2:
                    value = xldate_as_tuple(value, book.datemode)
                    value = date(*value[0:3]).strftime('%Y-%m-%d')
                    items.append(value)
                else:
                    items.append(value)
        if items:
            data.append(items)
    for key, value in enumerate(data):
        for k, v in enumerate(value):
            worksheet.write(key, k, v)
workbook.save(writer_file)
