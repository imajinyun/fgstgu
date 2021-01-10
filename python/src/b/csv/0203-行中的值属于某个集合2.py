import sys
import pandas as pd

reader_file, writer_file = sys.argv[1], sys.argv[2]
data = pd.read_csv(reader_file)
filters = ['1002', '1003']
condition = data.loc[data['Number'].isin(filters), :]
condition.to_csv(writer_file, index=False)
