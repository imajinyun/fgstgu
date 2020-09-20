"""
> pip3 install -U mysqlclient progressbar2 xlrd xlwt
> python3 ./20200914-导入签证.py ./20200914签证.xlsx
"""

# coding=utf-8
import os
import sys
import json
import datetime
import MySQLdb
import time
import traceback
from tqdm import tqdm
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
    items = []
    with open_workbook(file) as book:
        sheet = book.sheet_by_index(0)
        dt = datetime.datetime.now()
        formattedAt = dt.strftime('%Y-%m-%d %H:%M:%S')
        for row in range(1, sheet.nrows):
            ctype = sheet.cell_type(row, 0)
            values = sheet.row_values(row)
            city_id, map_name, country_name, subject, opt1, opt2, answer, number = values
            city_id, number = int(city_id), int(number)
            # print(city_id, map_name, country_name, subject, opt1, opt2, answer, number)
            options = [{'name': opt1}, {'name': opt2}]
            item = {}
            item['map_id'] = map_name
            item['country_id'] = country_name
            item['city_id'] = city_id
            item['subject'] = subject
            item['options'] = json.dumps(options, ensure_ascii=False)
            item['status'] = 1
            item['answer'] = number
            item['created_at'] = formattedAt
            item['updated_at'] = formattedAt
            items.append(item)
    return items


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


def getCountries():
    cursor = getMySQLConnect().cursor()
    cursor.execute('SELECT id,name FROM wx_walkup_country_483')
    entities = getDictFetchAll(cursor)
    countries = []
    if len(entities):
        exists = set()
        for country in entities:
            name = country['name']
            if name not in exists:
                countries.append({'id': country['id'], 'name': name})
            exists.add(name)
    return countries


if __name__ == "__main__":
    print('🚀 Visa data import starting...')

    maps = getMaps()
    mapDict = {}
    for map in maps:
        mapDict[map['name']] = map['id']
    mapNames = list(mapDict)

    countries = getCountries()
    countryDict = {}
    for country in countries:
        countryDict[country['name']] = country['id']
    countryNames = list(countryDict)

    getMySQLConnect().cursor().execute('TRUNCATE TABLE wx_walkup_visa_483')
    items = getWorksheetItems(sys.argv[1])

    if not items:
        exit('🧨 Items is empty')

    values = []
    for key, item in enumerate(items):
        if item['map_id'] in mapNames:
            item['map_id'] = mapDict[item['map_id']]
        if item['country_id'] in countryNames:
            item['country_id'] = countryDict[item['country_id']]
        tuples = tuple(item.values())
        values.append(tuples)

    sql = ''
    sql += 'INSERT INTO wx_walkup_visa_483 '
    sql += '(map_id,country_id,city_id,subject,options,status,answer,created_at,updated_at) VALUES '
    sql += '(%s, %s, %s, %s, %s, %s, %s, %s, %s)'

    getMySQLConnect().cursor().execute('TRUNCATE TABLE wx_walkup_visa_483')
    try:
        db = getMySQLConnect()
        cursor = db.cursor()
        cursor.execute('SET NAMES utf8;')
        cursor.execute('SET character_set_connection=utf8;')
        cursor.executemany(sql, values)
        db.commit()
    except Exception as e:
        db.rollback()
        traceback.print_exc()
    finally:
        cursor.close()
        db.close()
    print('🎉 Successfully import visa data to MySQL!')
