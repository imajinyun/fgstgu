import csv
import sys

reader_file, writer_file = sys.argv[1], sys.argv[2]
columns = [0, 3]

with open(reader_file, 'r', newline='') as csv_reader:
    with open(writer_file, 'w', newline='') as csv_writer:
        reader = csv.reader(csv_reader)
        writer = csv.writer(csv_writer)
        for row in reader:
            items = []
            for index in columns:
                items.append(row[index])
            writer.writerow(items)
