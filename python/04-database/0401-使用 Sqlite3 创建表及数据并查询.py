import sqlite3

connect = sqlite3.connect(':memory:')
query = """
CREATE TABLE students(
    number VARCHAR(20),
    name VARCHAR(20),
    gender CHAR(1),
    total_score FLOAT,
    created_at DATE
);
"""
connect.execute(query)
connect.commit()

data = [
    ('1001', 'Zhangming', 'M', '549.0', '2016-11-10'),
    ('1002', 'Qianxing', 'F', '619.5', '2016-09-01'),
    ('1003', 'Lizhenyu', 'M', '499.0', '2016-08-23'),
    ('1004', 'Wangxing', 'F', '586.5', '2016-09-03'),
]
statement = 'INSERT INTO students VALUES(?, ?, ?, ?, ?)'
connect.executemany(statement, data)
connect.commit()

cursor = connect.execute('SELECT * FROM students')
rows = cursor.fetchall()

counter = 0
for row in rows:
    print(row)
    counter += 1
print('Number of rows: %d' % (counter))
