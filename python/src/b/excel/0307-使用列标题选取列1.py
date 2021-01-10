import sys
from datetime import date
from xlrd import open_workbook, xldate_as_tuple
from xlwt import Workbook

reader_file, writer_file = sys.argv[1], sys.argv[2]
headers = ['Customer ID', 'Purchase Date']
workbook = Workbook()
worksheet = workbook.add_sheet('Sheet2018')
with open_workbook(reader_file) as book:
    sheet = book.sheet_by_name('Sheet2018')
    data = [headers]
    titles = sheet.row_values(0)
    filters = []
    for index in range(len(titles)):
        if titles[index] in headers:
            filters.append(index)
    for row in range(1, sheet.nrows):
        items = []
        for col in filters:
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
