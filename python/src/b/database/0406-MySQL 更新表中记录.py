import csv
import sys
import MySQLdb

writer_file = sys.argv[1]
connect = MySQLdb.connect(host='127.0.0.1', port=3306,
                          db='test', user='root', passwd='mysql')
cursor = connect.cursor()
reader = csv.reader(open(writer_file, 'r', newline=''), delimiter=',')
headers = next(reader, None)
for row in reader:
    data = []
    for col in range(len(headers)):
        if col in [1, 3, 5]:
            if col == 5:
                data.append(str(row[col]).lstrip('￥').strip())
            else:
                data.append(str(row[col]).strip())
    data.reverse()
    print(data)
    cursor.execute(
        'UPDATE employees SET salary=%s, school=%s WHERE name=%s',
        data
    )
connect.commit()
cursor.execute('SELECT * FROM employees')
rows = cursor.fetchall()
for row in rows:
    output = []
    for col in range(len(row)):
        output.append(str(row[col]))
    print(output)
