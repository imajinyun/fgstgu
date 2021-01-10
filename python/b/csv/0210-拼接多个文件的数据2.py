import os
import sys
import glob
import pandas as pd

path, writer_file = sys.argv[1], sys.argv[2]
files = glob.glob(os.path.join(path, '*.csv'))
items = []
first_file = True

for reader_file in files:
    filename = os.path.basename(reader_file)
    print(filename)
    if first_file:
        headers = ['Number', 'Name', 'Age', 'School', 'Position', 'Salary']
        data = pd.read_csv(reader_file, index_col=None, names=headers)
        # print(data)
        items.append(data)
        first_file = False
    elif filename == '0207.csv':
        data = pd.read_csv(reader_file, index_col=None, skiprows=3)
        data = data.drop(index=[8, 9, 10])
        data = data.reindex(data.index.drop(3))
        # print(data)
        items.append(data)
    else:
        data = pd.read_csv(reader_file, index_col=None)
        # print(data)
        items.append(data)

concat = pd.concat(items, axis=0, ignore_index=True, sort=False)
concat.to_csv(writer_file, index=False)
