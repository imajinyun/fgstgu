"""
> pip3 install -U mysqlclient progressbar2 xlrd xlwt
> python3 ./20200914-导入国家.py ./20200914签证.xlsx
"""

# coding=utf-8
import os
import sys
import json
import datetime
import MySQLdb
import time
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


if __name__ == "__main__":
    print('🚀 Country data import starting...')

    countries = getCountries()
    bar = progressbar.ProgressBar(max_value=progressbar.UnknownLength)
    dt = datetime.datetime.now()
    formattedAt = dt.strftime('%Y-%m-%d %H:%M:%S')

    getMySQLConnect().cursor().execute('TRUNCATE TABLE wx_walkup_country_483')
    for key, country in enumerate(countries):
        key += 1
        time.sleep(0.01)
        bar.update(key)

        item = {}
        item['id'] = country['id']
        item['name'] = country['name']
        item['description'] = country['name']
        item['status'] = 1
        item['created_at'] = formattedAt
        item['updated_at'] = formattedAt

        try:
            db = getMySQLConnect()
            cursor = db.cursor()
            cursor.execute('SET NAMES utf8;')
            cursor.execute('SET character_set_connection=utf8;')
            sql = 'INSERT INTO wx_walkup_country_483(id, name,status,description,created_at,updated_at) VALUES(%s, %s, %s, %s, %s, %s)'
            values = (item['id'], item['name'], item['status'], item['description'], item['created_at'], item['updated_at'])
            cursor.execute(sql, values)
            db.commit()
        except Exception as e:
            db.rollback()
            print(e)
    db.close()
    print('🎉 Successfully import country data to MySQL!')
