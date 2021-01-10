"""
任找一段 MySQL 错误日志使用脚本分析
"""
import sys


reader_file, writer_file = sys.argv[1], sys.argv[2]
messages, errors = {}, []
first = True
with open(reader_file, 'r', newline='') as text_reader:
    for row in text_reader:
        if '[ERROR]' in row:
            item = row.split(' ', 4)
            day = str(item[0]).strip().split(':')[0]
            error = item[4].strip('\n').strip()
            if error not in errors:
                errors.append(error)
            if day not in messages:
                messages[day] = {}
            if error not in messages[day]:
                messages[day][error] = 1
            else:
                messages[day][error] += 1
writer = open(writer_file, 'w', newline='')
headers = ['日期']
headers.extend(errors)
headers = ','.join(map(str, headers)) + '\n'
print(headers)

writer.write(headers)
for day, info in messages.items():
    items = []
    items.append(day)
    for index in range(len(errors)):
        if errors[index] in info.keys():
            items.append(info[errors[index]])
        else:
            items.append(0)
    output = ','.join(map(str, items)) + '\n'
    print(output)
    writer.write(output)
writer.close()
