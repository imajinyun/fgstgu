import sys
from datetime import date
from xlrd import open_workbook, xldate_as_tuple
from xlwt import Workbook

reader_file, writer_file = sys.argv[1], sys.argv[2]
workbook = Workbook()
worksheet = workbook.add_sheet('Sheet2018')
columns = [1, 4]

with open_workbook(reader_file) as book:
    sheet = book.sheet_by_name('Sheet2018')
    data = []
    for row in range(sheet.nrows):
        items = []
        for col in columns:
            value = sheet.cell_value(row, col)
            if sheet.cell_type(row, col) == 3:
                value = xldate_as_tuple(value, book.datemode)
                value = date(*value[0:3]).strftime('%Y-%m-%d')
                items.append(value)
            else:
                items.append(value)
        data.append(items)
    for key, value in enumerate(data):
        for k, v in enumerate(value):
            worksheet.write(key, k, v)
workbook.save(writer_file)
