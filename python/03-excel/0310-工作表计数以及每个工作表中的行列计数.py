import sys
import os
import glob
from xlrd import open_workbook

path, counter = sys.argv[1], 0
for reader_file in glob.glob(os.path.join(path, '*.xlsx')):
    workbook = open_workbook(reader_file)
    print('Workbook: %s' % os.path.basename(reader_file))
    print('Number of worksheets: %d' % workbook.nsheets)
    for worksheet in workbook.sheets():
        print('Worksheet name: {0}\tRows: {1}\tColumns: {2}'.format(
            worksheet.name, worksheet.nrows, worksheet.ncols))
    counter += 1
    print('-' * 80)
print('\nNumber of Excel workbooks: %d' % counter)
