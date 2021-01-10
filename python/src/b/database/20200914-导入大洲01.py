"""
> pip3 install -U mysqlclient tqdm dotenv xlrd xlwt
> python3 ./20200914-导入大洲01.py ./20200914-大洲信息.xlsx ./20200914-城市距离.xlsx
"""

# coding=utf-8
import os
import sys
import json
import datetime
import MySQLdb
import traceback
import time
import re
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


def getDictFetchAll(cursor):
    """
    Return all rows from a cursor as a dict
    """
    columns = [col[0] for col in cursor.description]
    return [
        dict(zip(columns, row))
        for row in cursor.fetchall()
    ]


def getMaps():
    db = getMySQLConnect()
    cursor = db.cursor()
    cursor.execute('SELECT id,name FROM wx_walkup_map_483')
    maps = getDictFetchAll(cursor)
    cursor.close()
    db.close()
    return maps


def getCities(ids):
    db = getMySQLConnect()
    cursor = db.cursor()
    sql = 'SELECT id AS city_id,map_id FROM wx_walkup_city_483 WHERE id IN(%s)'
    cursor.execute(sql % ','.join(map(str, ids)))
    cities = getDictFetchAll(cursor)
    cursor.close()
    db.close()
    return cities


def getCountries():
    cursor = getMySQLConnect().cursor()
    cursor.execute('SELECT id,cn_formatted_name FROM wx_walkup_city_483')
    cities = getDictFetchAll(cursor)
    countries = []
    if len(cities):
        exists = set()
        for index, city in enumerate(cities):
            name = city['cn_formatted_name'].split('·')[0]
            if name not in exists:
                countries.append({'id': index + 1, 'name': name})
            exists.add(name)
    return countries


def getWorksheetItems(file, bar):
    items = []
    with open_workbook(file) as book:
        sheet = book.sheet_by_index(0)
        dt = datetime.datetime.now()
        at = dt.strftime('%Y-%m-%d %H:%M:%S')
        for row in range(1, sheet.nrows):
            bar.update(row * 10)
            time.sleep(0.1)
            ctype = sheet.cell_type(row, 0)
            values = sheet.row_values(row)
            name, cn_name, en_name, total_city, total_event, total_visa, description = values
            item = {}
            item['name'] = cn_name
            item['en_name'] = en_name
            item['total_energy'] = 0
            item['total_num'] = 0
            item['total_city'] = total_city
            item['total_event'] = total_event
            item['total_visa'] = total_visa
            item['sort_value'] = row
            item['status'] = 1
            item['description'] = description
            item['image'] = ''
            item['create_user'] = '刘洋'
            item['update_user'] = '刘洋'
            item['created_at'] = at
            item['updated_at'] = at
            items.append(tuple(item.values()))
    return items


def isNumber(value):
    try:
        float(value)
        return True
    except ValueError:
        pass

    try:
        import unicodedata
        unicodedata.numeric(value)
        return True
    except (TypeError, ValueError):
        pass
    return False


if __name__ == "__main__":
    print('🚀 Continent data import starting...')

    bar = tqdm(total=100)
    items = getWorksheetItems(sys.argv[1], bar)
    if not items:
        exit('🧨 Items is empty')

    # 插入大洲地图信息
    sql = ''
    sql += 'INSERT INTO wx_walkup_map_483 '
    sql += '(name,en_name,total_energy,total_num,total_city,total_event,total_visa,sort_value,status,description,image,create_user,update_user,created_at,updated_at) VALUES '
    sql += '(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)'

    try:
        db = getMySQLConnect()
        cursor = db.cursor()
        cursor.execute('SET NAMES utf8;')
        cursor.execute('SET character_set_connection=utf8;')
        cursor.execute('TRUNCATE TABLE wx_walkup_map_483')
        cursor.executemany(sql, items)
        bar.update(85)
        db.commit()
    except Exception as e:
        db.rollback()
        traceback.print_exc()
        exit(e)
    finally:
        cursor.close()
        db.close()

    maps = getMaps()
    mapdict = {}
    for item in maps:
        mapdict[item['id']] = item['name']

    items = []
    with open_workbook(sys.argv[2]) as book:
        sheet = book.sheet_by_index(0)
        for i in range(2, sheet.nrows):
            row = sheet.row_values(i)
            value = row[10]
            if value and isNumber(value):
                items.append({
                    'city_id': int(row[0]),
                    'total_energy': int(row[10])
                })
    cityids = [item['city_id'] for item in items]
    cities = getCities(cityids)
    citydict = {}
    for city in cities:
        citydict[city['city_id']] = city['map_id']

    values = {}
    for item in items:
        if item['city_id'] in citydict.keys():
            id = citydict[item['city_id']]
            values[id] = item
    column = ''
    for value in values:
        energy = values[value]['total_energy']
        column += 'WHEN {} THEN {} '.format(value, energy)
    sql = 'UPDATE wx_walkup_map_483 SET total_energy = CASE id %s END WHERE id IN(%s)' \
        % (column, ','.join(map(str, values.keys())))

    try:
        db = getMySQLConnect()
        cursor = db.cursor()
        cursor.execute('SET NAMES utf8;')
        cursor.execute('SET character_set_connection=utf8;')
        cursor.execute(sql)
        bar.update(95)
        db.commit()
    except Exception as e:
        db.rollback()
        traceback.print_exc()
        exit(e)
    finally:
        bar.update(100)
        cursor.close()
        db.close()
        bar.close()
    print('🎉 Successfully import continent data to MySQL!')
