import csv
import sys

reader_file, writer_file = sys.argv[1], sys.argv[2]
filters = ['1002', '1004']

with open(reader_file, 'r', newline='') as csv_reader:
    with open(writer_file, 'w', newline='') as csv_writer:
        reader = csv.reader(csv_reader)
        writer = csv.writer(csv_writer)
        header = next(reader)
        writer.writerow(header)
        for row in reader:
            number = row[0]
            if number in filters:
                writer.writerow(row)
