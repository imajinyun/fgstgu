import csv
import glob
import sys
import os

path = sys.argv[1]
file_counter = 0
for file in glob.glob(os.path.join(path, '*.csv')):
    row_counter = 1
    with open(file, 'r', newline='') as csv_reader:
        reader = csv.reader(csv_reader)
        header = next(reader, None)
        for row in reader:
            row_counter += 1
    print('File: {0!s}\t\nRows: {1:d}\t\nCols: {2:d}\n'.format(
        os.path.basename(file), row_counter, len(header)))
    file_counter += 1
print('Number of files: {0:d}'.format(file_counter))
