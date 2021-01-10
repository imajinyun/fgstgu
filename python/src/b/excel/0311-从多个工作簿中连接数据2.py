import sys
import os
import glob
import pandas as pd

path, writer_file = sys.argv[1], sys.argv[2]
workbooks = glob.glob(os.path.join(path, '*.xlsx'))
frames = []
for workbook in workbooks:
    worksheets = pd.read_excel(workbook, sheet_name=None, index_col=None)
    for sheet_name, data in worksheets.items():
        data['Purchase Date'] = data['Purchase Date'].apply(
            lambda x: x.strftime('%Y-%m-%d'))
        frames.append(data)
concat = pd.concat(frames, axis=0, ignore_index=True)
writer = pd.ExcelWriter(writer_file)
concat.to_excel(writer, sheet_name='AllWorkbookToOne', index=False)
writer.save()
