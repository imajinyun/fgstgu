"""
> pip3 install -U mysqlclient progressbar xlrd xlwt
> python3 ./20200914-城市距离.py ./20200914城市距离.xlsx
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
        for row in range(2, sheet.nrows):
            size = sheet.ncols
            item = sheet.row_values(row, 0, size)
            if item:
                items.append({
                    'map_id': 0,
                    'city_id': int(item[0]),
                    'target_id': int(item[0]),
                    'position': 0,
                    'node_type': 3,
                    'unlock_coupon_num': 0,
                    'unlock_need_energy': int(item[4]),
                    'run_need_energy': int(item[6]),
                    'status': 1,
                    'create_user': '刘洋',
                    'update_user': '刘洋',
                    'created_at': at,
                    'updated_at': at
                })
    return items


def getDictFetchAll(cursor):
    """ Return all rows from a cursor as a dict """
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
    cursor.execute('SELECT id AS city_id,map_id FROM wx_walkup_city_483')
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
    print('🚀 Import node data starting...')

    maps = getMaps()
    mapdict = {}
    for map in maps:
        mapdict[map['id']] = map['name']

    cities = getCities()
    citydict = {}
    for city in cities:
        citydict[city['city_id']] = city['map_id']
    items = getWorksheetItems(sys.argv[1])

    if not items:
        exit('🧨 Items is empty')

    values = []
    i = j = k = l = m = n = p = q = 0
    for item in items:
        item['map_id'] = citydict[item['city_id']]
        if item['map_id'] == 1:
            i += 1
            item['position'] = i
        elif item['map_id'] == 2:
            j += 1
            item['position'] = j
        elif item['map_id'] == 3:
            k += 1
            item['position'] = k
        elif item['map_id'] == 4:
            l += 1
            item['position'] = l
        elif item['map_id'] == 5:
            m += 1
            item['position'] = m
        elif item['map_id'] == 6:
            n += 1
            item['position'] = n
        elif item['map_id'] == 7:
            p += 1
            item['position'] = p
        elif item['map_id'] == 8:
            q += 1
            item['position'] = q
        tuples = tuple(item.values())
        values.append(tuples)

    sql = ''
    sql += 'INSERT INTO wx_walkup_node_483 '
    sql += '(map_id,city_id,target_id,position,node_type,unlock_coupon_num,unlock_need_energy,run_need_energy,status,create_user,update_user,created_at,updated_at) VALUES '
    sql += '(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)'

    try:
        db = getMySQLConnect()
        cursor = db.cursor()
        cursor.executemany(sql, values)
        db.commit()
    except Exception as e:
        db.rollback()
        traceback.print_exc()
        exit(e)
    finally:
        cursor.close()
    print('🎉 Successfully import node data to MySQL!')
