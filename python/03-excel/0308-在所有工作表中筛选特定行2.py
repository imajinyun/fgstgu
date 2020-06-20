import sys
import datetime
import pandas as pd

from xlrd import xldate_as_tuple

reader_file, writer_file = sys.argv[1], sys.argv[2]
data = pd.read_excel(reader_file, sheet_name=None, index_col=None)
rows, given = [], 250000
for sheet_name, data in data.items():
    data['Purchase Date'] = data['Purchase Date'].apply(
        lambda x: x.strftime('%Y-%m-%d'))
    rows.append(data[data['Sale Amount'].astype(float) > given])
rows = pd.concat(rows, axis=0, ignore_index=True)
writer = pd.ExcelWriter(writer_file)
rows.to_excel(writer, sheet_name='AllSheets', index=False)
writer.save()
