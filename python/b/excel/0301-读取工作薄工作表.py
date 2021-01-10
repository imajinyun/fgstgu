"""
> python3 0301-读取工作薄工作表.py saleas.xlsx
"""

import sys
from xlrd import open_workbook

reader_file = sys.argv[1]
workbook = open_workbook(reader_file)
print('Number of worksheets: ', workbook.nsheets)

for sheet in workbook.sheets():
    print('Worksheet name: {0}\tRows: {1}\tCols: {2}'.format(
        sheet.name, sheet.nrows, sheet.ncols))
