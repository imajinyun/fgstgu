import sys
import pandas as pd

reader_file, writer_file = sys.argv[1], sys.argv[2]
data = pd.read_excel(reader_file, sheet_name=None, index_col=None)
columns = []
for sheet_name, data in data.items():
    columns.append(data.loc[:, ['Customer Name', 'Sale Amount']])
columns = pd.concat(columns, axis=0, ignore_index=True)
writer = pd.ExcelWriter(writer_file)
columns.to_excel(writer, sheet_name='AllSheets', index=False)
writer.save()
