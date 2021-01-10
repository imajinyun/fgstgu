"""
> pip3 install -U mysqlclient progressbar2 xlrd xlwt
> python3 ./20200914-导入国家02.py ./20200914签证.xlsx
"""

# coding=utf-8
import os
import sys
import json
import datetime
import MySQLdb
import traceback
import time
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
    cursor.execute('SELECT id,en_name,cn_formatted_name FROM wx_walkup_city_483')
    cities = getDictFetchAll(cursor)
    countries = []
    if len(cities):
        exists = set()
        for city in cities:
            name = city['cn_formatted_name'].split('.')[0]
            if city['en_name'] == 'Greater  Khingan':
                countries.append({'name': '中华河山'})
                break
            if name not in exists:
                countries.append({'name': name})
            exists.add(name)
    cursor.close()
    return countries

if __name__ == "__main__":
    print('🚀 Country data import starting...')

    countries = getCountries()
    dt = datetime.datetime.now()
    at = dt.strftime('%Y-%m-%d %H:%M:%S')

    values = []
    for index, country in enumerate(countries):
        item = {}
        item['id'] = index + 1
        item['name'] = country['name']
        item['status'] = 1
        item['description'] = country['name']
        item['created_at'] = at
        item['updated_at'] = at
        values.append(tuple(item.values()))

    try:
        db = getMySQLConnect()
        cursor = db.cursor()
        cursor.execute('SET NAMES utf8;')
        cursor.execute('SET character_set_connection=utf8;')
        cursor.execute('TRUNCATE TABLE wx_walkup_country_483')
        sql = 'INSERT INTO wx_walkup_country_483(id,name,status,description,created_at,updated_at) VALUES (%s, %s, %s, %s, %s, %s)'
        cursor.executemany(sql, values)
        db.commit()
    except Exception as e:
        db.rollback()
        traceback.print_exc()
        exit(e)
    finally:
        cursor.close()
        db.close()
    print('🎉 Successfully import country data to MySQL!')
