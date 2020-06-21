import csv
import sys
import MySQLdb

writer_file = sys.argv[1]
connect = MySQLdb.connect(host='127.0.0.1', port=3306,
                          db='test', user='root', passwd='mysql')
cursor = connect.cursor()
writer = csv.writer(open(writer_file, 'w', newline=''), delimiter=',')
headers = ['Number', 'Name', 'Age', 'School', 'Position', 'Salary']
writer.writerow(headers)
sql = 'SELECT number,name,age,school,position,salary FROM employees WHERE salary>30000'
cursor.execute(sql)
rows = cursor.fetchall()
[writer.writerow(row) for row in rows]
