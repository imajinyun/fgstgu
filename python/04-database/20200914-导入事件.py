"""
> pip3 install -U mysqlclient progressbar2 xlrd xlwt
> python3 ./20200914-导入事件.py ./20200914事件.xlsx
"""

# coding=utf-8
import os
import sys
import json
import datetime
import MySQLdb
import time
import traceback
import progressbar
from dotenv import load_dotenv
from datetime import date
from xlrd import open_workbook, xldate_as_tuple
from xlwt import Workbook


def getMySQLConnect():
    dotenv_path = os.path.join('../' + os.path.dirname(__file__), '.env')
    load_dotenv(dotenv_path)
    host = os.environ.get('MYSQL_DB_HOST')
    port = int(os.environ.get('MYSQL_DB_PORT'))
    dbname = os.environ.get('MYSQL_DB_NAME')
    dbuser = os.environ.get('MYSQL_DB_USER')
    dbpass = os.environ.get('MYSQL_DB_PASS')
    connect = MySQLdb.connect(host=host, port=port,
                              db=dbname, user=dbuser, passwd=dbpass)
    return connect


def getMergedCellValue(sheet, value, row, col):
    if value is None and value == '':
        for (rlow, rhigh, clow, chigh) in sheet.merged_cells:
            if rlow <= row < rhigh:
                if clow <= col < chigh:
                    value = sheet.cell_value(rlow, clow)
    return value


def getWorksheetItems(file):
    result = []
    dt = datetime.datetime.now()
    at = dt.strftime('%Y-%m-%d %H:%M:%S')
    with open_workbook(file) as book:
        sheet = book.sheet_by_index(0)
        items = []
        for row in range(1, sheet.nrows):
            record = {}
            item = []
            for col in range(sheet.ncols):
                value = sheet.row_values(row)[col]
                ctype = sheet.cell(row, col).ctype
                cvalue = sheet.cell_value(row, col)
                if ctype == 2 and cvalue % 1 == 0.0:
                    value = int(value)
                value = getMergedCellValue(sheet, value, row, col)
                item.append(value)
            items.append(item)
            if row % 3 == 0:
                options = []
                for key, val in enumerate(items):
                    options.append({
                        'name': val[5],
                        'coins': val[7],
                        'desc': val[6]
                    })
                record['id'] = items[key][1]
                record['subject'] = items[0][3]
                record['options'] = json.dumps(options, ensure_ascii=False)
                record['event_type'] = 2
                record['status'] = 1
                record['map_id'] = 0
                record['city_id'] = items[0][0]
                record['country_id'] = items[0][2].split('-')[0]
                record['answer'] = 0
                record['description'] = items[0][4]
                record['created_at'] = at
                record['updated_at'] = at
                result.append(record)
                items = []
                options = []
    return result


def getDictFetchAll(cursor):
    """Return all rows from a cursor as a dict"""
    columns = [col[0] for col in cursor.description]
    return [
        dict(zip(columns, row))
        for row in cursor.fetchall()
    ]


def getMaps():
    cursor = getMySQLConnect().cursor()
    cursor.execute('SELECT id,name FROM wx_walkup_map_483')
    maps = getDictFetchAll(cursor)
    return maps


def getCities():
    cursor = getMySQLConnect().cursor()
    cursor.execute('SELECT id,en_name FROM wx_walkup_city_483')
    cities = getDictFetchAll(cursor)
    return cities


def getCountries():
    cursor = getMySQLConnect().cursor()
    cursor.execute('SELECT id,name FROM wx_walkup_country_483')
    entities = getDictFetchAll(cursor)
    countries = []
    if len(entities):
        for entity in entities:
            countries.append({'id': entity['id'], 'name': entity['name']})
    return countries


def getChinaProvinces():
    return [
        '北京', '天津', '河北', '山西', '内蒙古', '辽宁',
        '吉林', '黑龙江', '上海', '江苏', '浙江', '安徽', '福建',
        '江西', '山东', '河南', '湖北', '湖南', '广东', '广西',
        '海南', '重庆', '四川', '贵州', '云南', '西藏', '陕西',
        '甘肃', '青海', '宁夏', '新疆', '台湾', '澳门', '香港',
    ]


if __name__ == "__main__":
    print('🚀 Event data import starting...')

    countries = getCountries()
    countrydict = {}
    for country in countries:
        countrydict[country['name']] = country['id']
    countrynames = list(countrydict)
    items = getWorksheetItems(sys.argv[1])

    if not items:
        exit('🧨 Items is empty')

    values = []
    for item in items:
        if item['country_id'] in countrynames:
            item['country_id'] = int(countrydict[item['country_id']])
        elif item['country_id'] in getChinaProvinces():
            item['country_id'] = 1
        tuples = tuple(item.values())
        values.append(tuples)

    sql = ''
    sql += 'INSERT INTO wx_walkup_question_483 '
    sql += '(`id`, `subject`, `options`, `event_type`, `status`, `map_id`, `country_id`, `city_id`, `answer`, `description`, `created_at`, `updated_at`) VALUES '
    sql += '(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)'

    getMySQLConnect().cursor().execute('TRUNCATE TABLE wx_walkup_question_483')
    try:
        db = getMySQLConnect()
        cursor = db.cursor()
        cursor.executemany(sql, values)
        db.commit()
    except Exception as e:
        db.rollback()
        traceback.print_exc()
    finally:
        cursor.close()
    print('🎉 Successfully import event data to MySQL!')
