"""

> pip3 install mysqlclient

CREATE TABLE IF NOT EXISTS `employees` (
  `id` int NOT NULL AUTO_INCREMENT,
  `number` int DEFAULT NULL,
  `name` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `age` tinyint DEFAULT NULL,
  `school` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `position` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `salary` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
"""
import csv
import sys
import MySQLdb

reader_file = sys.argv[1]
connect = MySQLdb.connect(host='127.0.0.1', port=3306, passwd='mysql',
                          db='test', user='root')
cursor = connect.cursor()
print(cursor)

reader = csv.reader(open(reader_file, 'r', newline=''))
header = next(reader)
for row in reader:
    data = []
    for col in range(len(header)):
        if col == 5:
            data.append(str(row[col]).lstrip('￥').strip())
        else:
            data.append(row[col])
    cursor.execute(
        'INSERT INTO employees VALUES (NULL, %s, %s, %s, %s, %s, %s)',
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
