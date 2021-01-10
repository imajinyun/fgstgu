import os
import sys
import glob
import pandas as pd

path, writer_file = sys.argv[1], sys.argv[2]
files = glob.glob(os.path.join(path, '*.csv'))
items = []

for reader_file in files:
    filename = os.path.basename(reader_file)
    print(filename)
    if filename == '0207.csv':
        data = pd.read_csv(reader_file, index_col=None, skiprows=3)
        data = data.drop(index=[8, 9, 10])
        # print(data)
    elif filename == '0208.csv':
        headers = ['Number', 'Name', 'Age', 'School', 'Position', 'Salary']
        data = pd.read_csv(reader_file, index_col=None, names=headers)
        # print(data)
    else:
        data = pd.read_csv(reader_file, index_col=None)
        # print(data)
    total_salaries = pd.DataFrame([
        float(str(value).strip('￥')) for value in data.loc[:, 'Salary']
    ]).sum()
    average_salaries = pd.DataFrame([
        float(str(value).strip('￥')) for value in data.loc[:, 'Salary']
    ]).mean()
    data = {
        'File Name': filename,
        'Total Salaries': total_salaries,
        'Average Salaries': average_salaries,
    }
    headers = ['File Name', 'Total Salaries', 'Average Salaries']
    items.append(pd.DataFrame(data, columns=headers))
concat = pd.concat(items, axis=0, ignore_index=True)
concat.to_csv(writer_file, index=False)
