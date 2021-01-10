import sys
import pandas as pd

reader_file, writer_file = sys.argv[1], sys.argv[2]
header = ['Number', 'Name', 'Age', 'School', 'Salary']
data = pd.read_csv(reader_file, header=None, names=header)
data.to_csv(writer_file, index=False)
