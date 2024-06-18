import mysql.connector
import time  
from time import sleep 
import random 
import requests 
from lxml import etree   
from fake_useragent import UserAgent
from tqdm import tqdm
import re
from datetime import datetime

def connect_db():
    return mysql.connector.connect(
        host="localhost",  # 数据库主机地址
        user="root",  # 数据库用户名
        password="password",  # 数据库密码
        database="paper_submission"  # 数据库名
    )

db = connect_db()
cursor = db.cursor()

cursor.execute("DROP TABLE IF EXISTS special_issues_temp;")
cursor.execute("DROP TABLE IF EXISTS journals_temp;")
cursor.execute("CREATE TABLE journals_temp LIKE journals;")
cursor.execute("CREATE TABLE special_issues_temp LIKE special_issues;")
db.commit()
 
headers = {  
    'User-Agent': UserAgent().random,  
    'Accept': '*/*',
    'Accept-Encoding': 'gzip, deflate, br, zstd',
    'Accept-Language': 'zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6',
    'Connection': 'keep-alive',
    'Cookie':'usprivacy=1Y--; __qca=P0-1597669840-1715239310443; sharedid=e7c8bc76-6b86-489c-9da8-4248032d6357; sharedid_cst=kSylLAssaw%3D%3D; _cc_id=961270ade12f96bd3823d08437cd22ab; cto_bidid=NC8nWF9wNUx3TDNzSWJLdEdqNXVXSUZaOElNMzZCNkxPdlA5VmZqMFZhc3NyNURWWFhCY29OUGVYR3dXUDM4OGJ6emNmOExmQyUyRlZpc21FOWFheTdTeHNyWTNTMEZFY3hpMWs2VnM1ZlRWYXdGbUw0JTNE; cto_dna_bundle=Hphqy193TU5PN2d4N2xNd2J1QmUlMkZlJTJCY2s1ek83NXYybGlMd1ZkZnozUHJLNzZaVWFwdER4MVQ0eXU2Wm04dzNKd3V0TUVKNE02MFVxaGd1d0xHVG5NYXIyYlElM0QlM0Q; cto_bundle=nPM8I19ZQXJSY1FqYUQyNXdlUWZMYjlwUWJ6Y0FMWFBzbkJlNEdsUmlDVzRXNGNxcTFWTHNqM3NXaTZmdWI1dDlaekZFMnlDRXBBQ3NsVFdYOVpWdzVzJTJGUiUyQkNwWDBGUngxanBEUjZwTE01cUFQcmJDb2R0NWdhYzk2RWM0RDFXNlg3JTJCZmgwcVhVVXdoNkdZSEROQWIwSHpEQVElM0QlM0Q; FCNEC=%5B%5B%22AKsRol-Z1xQSyE1O_lH6dPNJIGyrWRswpvtPi6GVcsoU0Vt3a7mHaAZktGYrWHYg-IMc0R2vgA5aesQTNoClZHa4NEnB-xWpcNXcFls4LeYqXB-28WbeZcsTNies0YuZ76rP56FulHDTzVLwafIAMoRMxr58I5aqvA%3D%3D%22%5D%5D; __gads=ID=eba12d6a8c731a57:T=1715239309:RT=1715392696:S=ALNI_MawbE7os-f-KgybKHzI94VOTJfqMw; __gpi=UID=00000e142e4325a9:T=1715239309:RT=1715392696:S=ALNI_Ma1IEagjw-cUBFHIwEpbAr6VIYLAQ; __eoi=ID=0b222e13a8e50b30:T=1715239309:RT=1715392696:S=AA-AfjbDIZt-RPSdA45nV1ZY6xXP; _gid=GA1.2.2031218189.1716815781; panoramaId_expiry=1716902203964; panoramaId=e2c74a4bce41c312bb7756d2f78da9fb927a912bad122d72bdd9345182762035; PHPSESSID=1b606a723a66e9c1039e7f0316833d66; 7a67863e0bfa801e8fae57623f8da4a8=3f0bb9a39706d5ed5ab370242eab7e15c17b28faa%3A4%3A%7Bi%3A0%3Bs%3A27%3A%2251265902073%40stu.ecnu.edu.cn%22%3Bi%3A1%3Bs%3A27%3A%2251265902073%40stu.ecnu.edu.cn%22%3Bi%3A2%3Bi%3A2592000%3Bi%3A3%3Ba%3A3%3A%7Bs%3A2%3A%22id%22%3Bs%3A5%3A%2265703%22%3Bs%3A4%3A%22name%22%3Bs%3A8%3A%22Yao+Ziqi%22%3Bs%3A4%3A%22type%22%3Bs%3A1%3A%22n%22%3B%7D%7D; _ga=GA1.1.625276632.1715172812; _ga_T0WW44V64X=GS1.1.1716815780.20.1.1716819260.28.0.0',
}   
session = requests.Session()

def safe_xpath_search(parse, xpath):
    results = parse.xpath(xpath)
    if results:
        return results[0]
    else:
        return ""

def safe_convert_to_int(value):
    try:
        return int(value)
    except ValueError:
        return None

def safe_convert_to_float(value):
    try:
        return float(value)
    except (ValueError, TypeError):
        return None

def parse_datetime(datetime_str):
    if not datetime_str or datetime_str.strip() == '':
        return None  # 如果字符串为空或只包含空格，则返回 None
    try:
        # 在解析前去除字符串两端的空格
        cleaned_date_str = datetime_str.strip()
        return datetime.strptime(cleaned_date_str, '%Y-%m-%d')
    except ValueError:
        return None  # 如果日期格式不正确，则返回 None

def get_journal(url):
    global journal_id
    response = session.get(url, headers=headers)
    if response.status_code != 200:
        print(response.status_code,"连接失败")
        return
    parse = etree.HTML(response.text)  #解析网页 
    all_tr = parse.xpath('//*[@id="yw6"]/table/tbody/tr')
    for tr in all_tr:
        journal_id += 1
        ccf = safe_xpath_search(tr,'./td[1]/span/text()')
        abbreviation = safe_xpath_search(tr,'./td[2]/text()')
        if abbreviation == '\xa0':
            abbreviation = ''
        fullname = safe_xpath_search(tr,'./td[3]/a/text()')
        impactFactor = safe_xpath_search(tr,'./td[4]/text()') 
        publisher = safe_xpath_search(tr,'./td[5]/text()')
        issn = safe_xpath_search(tr,'./td[6]/text()')
        # 进入期刊界面爬取剩余信息
        journalUrl = 'https://www.myhuiban.com'+tr.xpath('./td[3]/a/@href')[0]
        res = session.get(journalUrl, headers=headers)
        ps = etree.HTML(res.text)
        link = safe_xpath_search(ps,'//*[@id="yw0"]/div[2]/div/a/text()')
        description = safe_xpath_search(ps,'//*[@id="yw1"]/div[2]/pre/text()')
        deadline = ''
        specialIssues = []
        issueList = ps.xpath('//*[@id="yw2"]/div[2]/pre')
        if issueList:
            lastIssue = issueList[-1].xpath('./text()')[0]
            match = re.search(r'截稿日期:\s*(\d{4}-\d{2}-\d{2})', lastIssue)
            deadline = match.group(1)
            # 将special issues写入specialissues表
            for issue in issueList:
                issueData = (journal_id, issue.xpath('./text()')[0])
                specialIssues.append(issueData)
                # cursor.execute("INSERT INTO specialissues (journal_id, IssuesContent) VALUES (%s, %s)", issueData)

        impactFactor = safe_convert_to_float(impactFactor)
        deadline = parse_datetime(deadline)
        data = (ccf, abbreviation, fullname, impactFactor, publisher, issn, link, deadline, description)
        cursor.execute("INSERT INTO journals_temp (ccf_ranking, abbreviation, full_name, impact_factor, publisher, issn, link, deadline, description) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s)", data)
        for issueData in specialIssues:
            cursor.execute("INSERT INTO special_issues_temp (journal_id, issue_content) VALUES (%s, %s)", issueData)
        db.commit()
        sleep(random.uniform(0.1, 0.5))

base_url = 'https://www.myhuiban.com/journals?ajax=yw6&Journal_page='

pages = 56 #56
journal_id = 0
for page_number in tqdm(range(1, pages + 1), desc="爬取期刊"):
    url = f'{base_url}{page_number}'
    get_journal(url)
    sleep(random.uniform(0.5, 2))

try:
    cursor.execute("DROP TABLE IF EXISTS special_issues;")
    cursor.execute("DROP TABLE IF EXISTS journals;")
    cursor.execute("RENAME TABLE journals_temp TO journals;")
    cursor.execute("RENAME TABLE special_issues_temp TO special_issues;")
    db.commit()
    print("数据库更新成功！")
except mysql.connector.Error as err:
    print("执行数据库操作失败：", err)
finally:
    cursor.close()
    db.close()


