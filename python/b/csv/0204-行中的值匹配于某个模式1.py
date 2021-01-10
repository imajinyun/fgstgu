import csv
import sys
import re

reader_file, writer_file = sys.argv[1], sys.argv[2]
pattern = re.compile(r'(?P<custom_group>^Zh.*)', re.I)

with open(reader_file, 'r', newline='') as csv_reader:
    with open(writer_file, 'w', newline='') as csv_writer:
        reader = csv.reader(csv_reader)
        writer = csv.writer(csv_writer)
        header = next(reader)
        writer.writerow(header)
        for row in reader:
            name = row[1]
            if pattern.search(name):
                writer.writerow(row)
