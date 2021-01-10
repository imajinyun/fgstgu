import sys
import pandas as pd

reader_file, writer_file = sys.argv[1], sys.argv[2]
data = pd.read_excel(reader_file, sheet_name='Sheet2018')
writer = pd.ExcelWriter(writer_file)
data.to_excel(writer, sheet_name='Sheet2018A', index=False)
writer.save()
