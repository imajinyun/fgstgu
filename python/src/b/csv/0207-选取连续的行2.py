import sys
import pandas as pd

reader_file, writer_file = sys.argv[1], sys.argv[2]
data = pd.read_csv(reader_file, header=None)
data = data.drop([0, 1, 2, 12, 13, 14])
data.columns = data.iloc[0]
data = data.reindex(data.index.drop(3))
data.to_csv(writer_file, index=False)
