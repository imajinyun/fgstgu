"""
> pip3 install -U mysqlclient tqdm dotenv xlrd xlwt
> python3 ./20200914-导入大洲.py ./20200914-大洲信息.xlsx
"""

# coding=utf-8
import os
import sys
import json
import datetime
import MySQLdb
import traceback
import time
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
    cursor = getMySQLConnect().cursor()
    cursor.execute('SELECT id,name FROM wx_walkup_map_483')
    maps = getDictFetchAll(cursor)
    return maps


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
        bar.update(90)
        cursor.executemany(sql, items)
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
