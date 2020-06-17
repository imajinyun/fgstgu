"""
> pip3 install xlrd xlwt
"""
import sys
from datetime import date
from xlrd import open_workbook, xldate_as_tuple
from xlwt import Workbook

reader_file, writer_file = sys.argv[1], sys.argv[2]
workbook = Workbook()
worksheet = workbook.add_sheet('Sheet2021')
with open_workbook(reader_file) as book:
    sheet = book.sheet_by_name('Sheet2018')
    for row in range(sheet.nrows):
        headers = []
        for col in range(sheet.ncols):
            if sheet.cell_type(row, col) == 3:
                value = xldate_as_tuple(
                    sheet.cell_value(row, col),
                    book.datemode
                )
                value = date(*value[0:3]).strftime('%Y-%m-%d')
                headers.append(value)
                worksheet.write(row, col, value)
            else:
                value = sheet.cell_value(row, col)
                headers.append(value)
                worksheet.write(row, col, value)
workbook.save(writer_file)
