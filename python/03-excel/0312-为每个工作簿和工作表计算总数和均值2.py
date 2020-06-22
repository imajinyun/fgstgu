import sys
import glob
import os
import pandas as pd

path, writer_file = sys.argv[1], sys.argv[2]
data = []
workbooks = glob.glob(os.path.join(path, '*.xlsx'))
for workbook in workbooks:
    worksheets = pd.read_excel(workbook, sheet_name=None, index_col=None)
    workbook_totals = []
    workbook_numbers = []
    worksheet_frames = []
    books = None
    sheets = None
    for sheet_name, item in worksheets.items():
        column = [
            float(str(value).strip('￥')) for value in item.loc[:, 'Sale Amount']
        ]
        total_sales = pd.DataFrame(column).sum()
        number_sales = len(item.loc[:, 'Sale Amount'])
        workbook_totals.append(total_sales)
        workbook_numbers.append(number_sales)
        items = {
            'workbook': os.path.basename(workbook),
            'worksheet': sheet_name,
            'worksheet_total': total_sales,
            'worksheet_average': total_sales / number_sales
        }
        worksheet_frames.append(pd.DataFrame(items, columns=[
            'workbook', 'worksheet', 'worksheet_total', 'worksheet_average',
        ]))
    sheets = pd.concat(worksheet_frames, axis=0, ignore_index=True)
    workbook_total = pd.DataFrame(workbook_totals).sum()
    workbook_number = pd.DataFrame(workbook_numbers).sum()
    stats = pd.DataFrame({
        'workbook': os.path.basename(workbook),
        'workbook_total': workbook_total,
        'workbook_average': workbook_total / workbook_number
    }, columns=['workbook', 'workbook_total', 'workbook_average'])
    books = pd.merge(sheets, stats, on='workbook', how='left')
    data.append(books)
concat = pd.concat(data, axis=0, ignore_index=True)
writer = pd.ExcelWriter(writer_file)
concat.to_excel(writer, sheet_name='SumAndAverageSheet', index=False)
writer.save()
