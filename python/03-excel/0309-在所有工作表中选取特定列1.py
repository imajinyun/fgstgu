import sys
from datetime import date
from xlrd import open_workbook, xldate_as_tuple
from xlwt import Workbook

reader_file, writer_file = sys.argv[1], sys.argv[2]
columns, first, cols = ['Customer Name', 'Sale Amount'], True, []
workbook = Workbook()
worksheet = workbook.add_sheet('AllSheets')
with open_workbook(reader_file) as book:
    data = [columns]
    for sheet in book.sheets():
        if first:
            headers = sheet.row_values(0)
            for column in range(len(headers)):
                if headers[column] in columns:
                    cols.append(column)
            first = False
        for row in range(1, sheet.nrows):
            items = []
            for col in cols:
                value = sheet.cell_value(row, col)
                if sheet.cell_type(row, col) == 3:
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
