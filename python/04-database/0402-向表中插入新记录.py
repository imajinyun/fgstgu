import csv
import sys
import sqlite3

reader_file = sys.argv[1]
connect = sqlite3.connect('employees.db')
cursor = connect.cursor()
create_table = """
CREATE TABLE IF NOT EXISTS employees(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    number CHAR(4),
    name VARCHAR(20),
    age TINYINT,
    school VARCHAR(50),
    position VARCHAR(50),
    salary FLOAT
);
"""
cursor.execute(create_table)
connect.commit()

reader = csv.reader(open(reader_file, 'r'), delimiter=',')
header = next(reader, None)
for row in reader:
    data = []
    for col in range(len(header)):
        data.append(row[col])
    print(data)
    cursor.execute(
        'INSERT INTO employees VALUES(NULL, ?, ?, ?, ?, ?, ?)',
        data
    )
connect.commit()
print('')

output = cursor.execute('SELECT * FROM employees')
rows = output.fetchall()
for row in rows:
    output = []
    for col in range(len(row)):
        output.append(str(row[col]))
    print(output)
