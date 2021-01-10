"""
$ pip3 install requests
"""
import requests
import time
import random
from openpyxl import Workbook


def getWebResult(url, page, name):
    origin = 'https://www.lagou.com'
    referer = origin + \
        '/jobs/list_{}/p-city_2?&cl=false&fromSearch=true&labelWords=&suginput='.format(
            name)
    headers = {
        'Origin': origin,
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.36',
        'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8',
        'Accept': 'application/json, text/javascript, */*; q=0.01',
        'Accept-Encoding': 'gzip, deflate, br',
        'Accept-Language': 'zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7,ca;q=0.6,cy;q=0.5',
        'Referer': referer,
    }
    first = False if page > 1 else True
    data = {'first': first, 'pn': page, 'kd': name}
    session = requests.Session()
    session.get(referer, headers=headers, timeout=3)
    cookies = session.cookies
    json = requests.post(url, data, headers=headers,
                         cookies=cookies, timeout=5).json()
    items = json['content']['positionResult']['result']
    result = []
    for i in items:
        info = []
        info.append(i.get('companyShortName', '无'))
        info.append(i.get('companyFullName', '无'))
        info.append(i.get('industryField', '无'))
        info.append(i.get('companySize', '无'))
        info.append(i.get('salary', '无'))
        info.append(i.get('city', '无'))
        info.append(i.get('education', '无'))
        result.append(info)
    return result


def main():
    name = 'Python'
    wb = Workbook()
    sheet = wb.active
    sheet.title = name
    city = '北京'
    url = 'https://www.lagou.com/jobs/positionAjax.json?city={}&needAddtionalResult=false'.format(
        city)
    for page in range(1, 2):
        info = getWebResult(url, page, name)
        time.sleep(random.randint(5, 10))
        [sheet.append(row) for row in info]
    path = '/tmp/{}职位信息.xlsx'.format(name)
    wb.save(path)


if __name__ == "__main__":
    main()
