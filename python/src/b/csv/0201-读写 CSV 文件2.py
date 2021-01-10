import sys
import pandas as pd

reader_file = sys.argv[1]
writer_file = sys.argv[2]

data = pd.read_csv(reader_file)
data.to_csv(writer_file, index=False)
