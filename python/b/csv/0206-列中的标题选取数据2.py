import sys
import pandas as pd

reader_file, writer_file = sys.argv[1], sys.argv[2]
data = pd.read_csv(reader_file)
condition = data.loc[:, ['Name', 'School']]
condition.to_csv(writer_file, index=False)
