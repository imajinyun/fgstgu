import csv
import sys
import os
import glob

path, writer_file, first_file, counter = sys.argv[1], sys.argv[2], True, 0
for reader_file in glob.glob(os.path.join(path, '*.csv')):
    filename = os.path.basename(reader_file)
    print(filename)
    with open(reader_file, 'r', newline='') as csv_reader:
        with open(writer_file, 'a', newline='') as csv_writer:
            reader = csv.reader(csv_reader)
            writer = csv.writer(csv_writer)
            if first_file:
                header = ['Number', 'Name', 'Age', 'School', 'Salary']
                writer.writerow(header)
                for row in reader:
                    writer.writerow(row)
                first_file = False
            else:
                header = next(reader, None)
                for row in reader:
                    if filename == '0207.csv':
                        if counter >= 3 and counter <= 10:
                            writer.writerow(row)
                        counter += 1
                    else:
                        writer.writerow(row)
print('Successfully spliced file data')
