import sys
import re
import pandas as pd

reader_file, writer_file = sys.argv[1], sys.argv[2]

data = pd.read_csv(reader_file)
condition = data.loc[data['Name'].str.startswith('Zh'), :]
condition.to_csv(writer_file, index=False)
