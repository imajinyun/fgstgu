import csv
import sys

reader_file = sys.argv[1]
writer_file = sys.argv[2]

with open(reader_file, 'r', newline='') as csv_reader:
    with open(writer_file, 'w', newline='') as csv_writer:
        reader = csv.reader(csv_reader)
        writer = csv.writer(csv_writer)
        header = next(reader)
        print(header)
        writer.writerow(header)
        for row in reader:
            number = str(row[0]).strip()
            salary = str(row[5]).strip('￥').replace(',', '')
            print(number, salary)
            if number == '1006' or float(salary) > 50000.0:
                writer.writerow(row)
