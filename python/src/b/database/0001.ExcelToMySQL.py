"""
导入 Excel 到 MySQL 数据库。

CREATE TABLE `country_code` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `country_code` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '国别代码',
  `dialing_code` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '电话代码',
  `nation_code` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '国家代码',
  `cn_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '中文名称',
  `en_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '英文名称',
  `abbr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '简称',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '创建者',
  `reviser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '修改者',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='国家拨号代码表';

> pip3 install -U python-dotenv mysqlclient progressbar2 tqdm xlrd xlwt
> python3 ./mdm-country-code.xls
"""

import datetime
import json
import os
import sys
import time
import traceback
from datetime import date

import MySQLdb
from dotenv import load_dotenv
from tqdm import tqdm
from xlrd import open_workbook, xldate_as_tuple
from xlwt import Workbook


def get_mysql_connect():
    dirname = os.path.dirname(os.path.abspath('__file__'))
    dotenv_path = os.path.join(dirname, '.env')
    load_dotenv(dotenv_path)
    host = os.environ.get('MYSQL_DB_HOST')
    port = int(os.environ.get('MYSQL_DB_PORT'))
    dbname = os.environ.get('MYSQL_DB_NAME')
    dbuser = os.environ.get('MYSQL_DB_USER')
    dbpass = os.environ.get('MYSQL_DB_PASS')
    connect = MySQLdb.connect(host=host, port=port,
                              db=dbname, user=dbuser, passwd=dbpass)
    return connect


def migration() -> str:
    return """
        CREATE TABLE `country_code` (`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
            `country_code` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '国别代码',
            `dialing_code` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '电话代码',
            `nation_code` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '国家代码',
            `cn_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '中文名称',
            `en_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '英文名称',
            `abbr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '简称',
            `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '创建者',
            `reviser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '修改者',
            `created_at` datetime NULL DEFAULT NULL,
            `updated_at` datetime NULL DEFAULT NULL,
            PRIMARY KEY(`id`)
        ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '国家拨号代码表';
    """


def get_merged_cell_value(sheet, value, row, col):
    if value is None and value == '':
        for (rlow, rhigh, clow, chigh) in sheet.merged_cells:
            if rlow <= row < rhigh:
                if clow <= col < chigh:
                    value = sheet.cell_value(rlow, clow)
    return value


def get_worksheet_items(file):
    items = []
    with open_workbook(file) as book:
        sheet = book.sheet_by_index(0)
        dt = datetime.datetime.now()
        at = dt.strftime('%Y-%m-%d %H:%M:%S')
        for row in range(1, sheet.nrows):
            values = sheet.row_values(row)
            _, country_name, dailing_code, cn_name, en_name, abbr, \
                creator, reviser, created_at, updated_at = values
            item = {}
            item['country_name'] = country_name
            item['dailing_code'] = dailing_code
            item['nation_code'] = ''
            item['cn_name'] = cn_name
            item['en_name'] = en_name
            item['abbr'] = abbr
            item['creator'] = creator
            item['reviser'] = reviser
            item['created_at'] = at
            item['updated_at'] = at
            items.append(item)
    return items


if __name__ == "__main__":
    print('🚀 Country code data import starting...')

    items = get_worksheet_items(sys.argv[1])
    if not items:
        exit('🧨 Items is empty')

    values = []
    for key, item in enumerate(items):
        tuples = tuple(item.values())
        values.append(tuples)

    sql = ''
    sql += 'INSERT INTO country_code '
    sql += '(country_code,dailing_code,nation_code,cn_name,en_name,abbr,creator,reviser,created_at,updated_at) VALUES '
    sql += '(%s, %s, %s, %s, %s, %s, %s, %s)'

    db = get_mysql_connect()
    try:
        cursor = db.cursor()
        cursor.execute('SET NAMES utf8mb4;')
        cursor.execute('SET character_set_connection=utf8mb4;')
        cursor.execute('DROP TABLE IF EXISTS country_code;')
        cursor.execute(migration())
        # cursor.execute('TRUNCATE TABLE country_code;')
        cursor.executemany(sql, values)
        db.commit()
    except Exception as e:
        db.rollback()
        traceback.print_exc()
    finally:
        cursor.close()
        db.close()
    print('🎉 Successfully import country code data to MySQL!')
