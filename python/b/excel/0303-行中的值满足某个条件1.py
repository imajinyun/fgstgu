import sys
from datetime import date
from xlrd import open_workbook, xldate_as_tuple
from xlwt import Workbook

reader_file, writer_file = sys.argv[1], sys.argv[2]
workbook = Workbook()
worksheet = workbook.add_sheet('Sheet2018')
index, given = 3, 200000.0
with open_workbook(reader_file) as book:
    sheet = book.sheet_by_name('Sheet2018')
    data = []
    headers = sheet.row_values(0)
    data.append(headers)
    for row in range(1, sheet.nrows):
        items = []
        sales_amount = sheet.cell_value(row, index)
        if sales_amount > given:
            for col in range(sheet.ncols):
                value = sheet.cell_value(row, col)
                cell_type = sheet.cell_type(row, col)
                if sheet.cell_type(row, col) == 3:
                    value = xldate_as_tuple(value, book.datemode)
                    value = date(*value[0:3]).strftime('%Y-%m-%d')
                    items.append(value)
                else:
                    items.append(value)
        if items:
            data.append(items)
    for index, items in enumerate(data):
        for k, v in enumerate(items):
            worksheet.write(index, k, v)
workbook.save(writer_file)
