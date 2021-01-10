import csv
import glob
import os
import sys

path, writer_file = sys.argv[1], sys.argv[2]
headers = ['FileName', 'TotalSalaries', 'AverageSalaries']

csv_writer = open(writer_file, 'a', newline='')
writer = csv.writer(csv_writer)
writer.writerow(headers)


def to_float(source):
    return float(str(source).strip().lstrip('￥'))


for reader_file in glob.glob(os.path.join(path, '*.csv')):
    with open(reader_file, 'r', newline='') as csv_reader:
        filename = os.path.basename(reader_file)
        print(filename)
        reader = csv.reader(csv_reader)

        if filename in ['0201.csv', '0207.csv']:
            header = next(reader)
        items = []
        items.append(os.path.basename(reader_file))
        total_salaries, average_salaries = 0.0, 0.0
        counter = 0
        for row in reader:
            if filename == '0207.csv':
                if counter > 2 and counter < 11:
                    total_salaries += to_float(row[5])
                    print(to_float(row[5]))
                    average_salaries += 1
                    pass
            else:
                total_salaries += to_float(row[5])
                average_salaries += 1
            counter += 1
        average_salaries = '{0:.2f}'.format(total_salaries / average_salaries)
        items.append(total_salaries)
        items.append(average_salaries)
        writer.writerow(items)
csv_writer.close()
