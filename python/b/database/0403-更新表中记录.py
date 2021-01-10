import csv
import sys
import sqlite3

reader_file = sys.argv[1]
connect = sqlite3.connect(':memory:')
query = """
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
connect.execute(query)
connect.commit()

data = [
    ('1001', 'Lisi', 23, 'A', 'Web Engineer', ''),
    ('1002', 'Wangwu', 31, 'B', 'Java Engineer', ''),
    ('1003', 'Mayun', 29, 'C', 'Python Engineer', ''),
    ('1004', 'Chenlong', 27, 'D', 'JavaScript Engineer', '')
]
for tup in data:
    print(tup)
statement = "INSERT INTO employees VALUES(NULL, ?, ?, ?, ?, ?, ?)"
connect.executemany(statement, data)
connect.commit()
print('-' * 80)

reader = csv.reader(open(reader_file, 'r'), delimiter=',')
header = next(reader, None)
print(header)
for row in reader:
    data = []
    for col in range(len(header)):
        if col in [1, 3, 5]:
            data.append(row[col])
    data.reverse()
    print(data)
    sql = 'UPDATE employees SET salary=?, school=? WHERE name=?;'
    connect.execute(sql, data)
connect.commit()
print('-' * 80)

cursor = connect.execute('SELECT * FROM employees')
rows = cursor.fetchall()
for row in rows:
    output = []
    for col in range(len(row)):
        output.append(str(row[col]))
    print(output)
