import sys
import pandas as pd

reader_file, writer_file = sys.argv[1], sys.argv[2]
data = pd.read_excel(reader_file, 'Sheet2018', index_col=None)
condition = data.loc[:, ['Customer ID', 'Purchase Date']]
writer = pd.ExcelWriter(writer_file)
condition.to_excel(writer, sheet_name='Sheet2018', index=False)
writer.save()
