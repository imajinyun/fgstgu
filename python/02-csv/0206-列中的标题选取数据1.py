import csv
import sys

reader_file, writer_file = sys.argv[1], sys.argv[2]
columns = ['Name', 'School']
items = []

with open(reader_file, 'r', newline='') as csv_reader:
    with open(writer_file, 'w', newline='') as csv_writer:
        reader = csv.reader(csv_reader)
        writer = csv.writer(csv_writer)
        header = next(reader, None)
        for index in range(len(header)):
            if header[index] in columns:
                items.append(index)
        writer.writerow(columns)
        for row in reader:
            outs = []
            for index in items:
                outs.append(row[index])
            writer.writerow(outs)
