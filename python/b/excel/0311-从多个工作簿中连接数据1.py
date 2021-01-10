import sys
import os
import glob
from datetime import date
from xlrd import open_workbook, xldate_as_tuple
from xlwt import Workbook

path, writer_file, data, first = sys.argv[1], sys.argv[2], [], True
workbook = Workbook()
worksheet = workbook.add_sheet('AllSheets')
for reader_file in glob.glob(os.path.join(path, '*.xlsx')):
    print(os.path.basename(reader_file))
    with open_workbook(reader_file) as book:
        for sheet in book.sheets():
            if first:
                headers = sheet.row_values(0)
                data.append(headers)
                first = False
            for row in range(1, sheet.nrows):
                rows = []
                for col in range(sheet.ncols):
                    value = sheet.cell_value(row, col)
                    if sheet.cell_type(row, col) == 3:
                        value = xldate_as_tuple(value, book.datemode)
                        value = date(*value[0:3]).strftime('%Y-%m-%d')
                        rows.append(value)
                    else:
                        rows.append(value)
                if rows:
                    data.append(rows)
for key, value in enumerate(data):
    for k, v in enumerate(value):
        worksheet.write(key, k, v)
workbook.save(writer_file)
