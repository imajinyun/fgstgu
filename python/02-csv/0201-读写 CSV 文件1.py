import sys

reader_file = sys.argv[1]
writer_file = sys.argv[2]

with open(reader_file, 'r', newline='') as reader:
    with open(writer_file, 'w', newline='') as writer:
        header = reader.readline().strip()
        headers = header.split(',')
        print(headers)
        writer.write(','.join(map(str, headers)) + '\n')
        for row in reader:
            row = row.strip()
            items = row.split(',')
            print(items)
            writer.write(','.join(map(str, items)) + '\n')
