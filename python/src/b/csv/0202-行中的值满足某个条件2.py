import sys
import pandas as pd

reader_file, writer_file = sys.argv[1], sys.argv[2]
data = pd.read_csv(reader_file)
data['Salary'] = data['Salary'].str.strip('￥').astype(float)
condition = data.loc[(data['Name'].str.contains('Z'))
                     & (data['Salary'] >= 35000.0), :]
condition.to_csv(writer_file, index=False)
