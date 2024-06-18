import time  
from time import sleep 
import random 
import requests 
from lxml import etree   
from fake_useragent import UserAgent
from tqdm import tqdm
import mysql.connector
from datetime import datetime
from requests.adapters import HTTPAdapter
from urllib3.util.retry import Retry

url = 'https://www.myhuiban.com/conferences' 
headers = {  
    'User-Agent': UserAgent().random,  
    'Accept': '*/*',
    'Accept-Encoding': 'gzip, deflate, br, zstd',
    'Accept-Language': 'zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6',
    'Connection': 'keep-alive',
    'Cookie':'usprivacy=1Y--; __qca=P0-1597669840-1715239310443; sharedid=e7c8bc76-6b86-489c-9da8-4248032d6357; sharedid_cst=kSylLAssaw%3D%3D; _cc_id=961270ade12f96bd3823d08437cd22ab; cto_bidid=NC8nWF9wNUx3TDNzSWJLdEdqNXVXSUZaOElNMzZCNkxPdlA5VmZqMFZhc3NyNURWWFhCY29OUGVYR3dXUDM4OGJ6emNmOExmQyUyRlZpc21FOWFheTdTeHNyWTNTMEZFY3hpMWs2VnM1ZlRWYXdGbUw0JTNE; cto_dna_bundle=Hphqy193TU5PN2d4N2xNd2J1QmUlMkZlJTJCY2s1ek83NXYybGlMd1ZkZnozUHJLNzZaVWFwdER4MVQ0eXU2Wm04dzNKd3V0TUVKNE02MFVxaGd1d0xHVG5NYXIyYlElM0QlM0Q; cto_bundle=nPM8I19ZQXJSY1FqYUQyNXdlUWZMYjlwUWJ6Y0FMWFBzbkJlNEdsUmlDVzRXNGNxcTFWTHNqM3NXaTZmdWI1dDlaekZFMnlDRXBBQ3NsVFdYOVpWdzVzJTJGUiUyQkNwWDBGUngxanBEUjZwTE01cUFQcmJDb2R0NWdhYzk2RWM0RDFXNlg3JTJCZmgwcVhVVXdoNkdZSEROQWIwSHpEQVElM0QlM0Q; FCNEC=%5B%5B%22AKsRol-Z1xQSyE1O_lH6dPNJIGyrWRswpvtPi6GVcsoU0Vt3a7mHaAZktGYrWHYg-IMc0R2vgA5aesQTNoClZHa4NEnB-xWpcNXcFls4LeYqXB-28WbeZcsTNies0YuZ76rP56FulHDTzVLwafIAMoRMxr58I5aqvA%3D%3D%22%5D%5D; __gads=ID=eba12d6a8c731a57:T=1715239309:RT=1715392696:S=ALNI_MawbE7os-f-KgybKHzI94VOTJfqMw; __gpi=UID=00000e142e4325a9:T=1715239309:RT=1715392696:S=ALNI_Ma1IEagjw-cUBFHIwEpbAr6VIYLAQ; __eoi=ID=0b222e13a8e50b30:T=1715239309:RT=1715392696:S=AA-AfjbDIZt-RPSdA45nV1ZY6xXP; _gid=GA1.2.2031218189.1716815781; panoramaId_expiry=1716902203964; panoramaId=e2c74a4bce41c312bb7756d2f78da9fb927a912bad122d72bdd9345182762035; PHPSESSID=1b606a723a66e9c1039e7f0316833d66; 7a67863e0bfa801e8fae57623f8da4a8=3f0bb9a39706d5ed5ab370242eab7e15c17b28faa%3A4%3A%7Bi%3A0%3Bs%3A27%3A%2251265902073%40stu.ecnu.edu.cn%22%3Bi%3A1%3Bs%3A27%3A%2251265902073%40stu.ecnu.edu.cn%22%3Bi%3A2%3Bi%3A2592000%3Bi%3A3%3Ba%3A3%3A%7Bs%3A2%3A%22id%22%3Bs%3A5%3A%2265703%22%3Bs%3A4%3A%22name%22%3Bs%3A8%3A%22Yao+Ziqi%22%3Bs%3A4%3A%22type%22%3Bs%3A1%3A%22n%22%3B%7D%7D; _ga=GA1.1.625276632.1715172812; _ga_T0WW44V64X=GS1.1.1716815780.20.1.1716819260.28.0.0',
}
session = requests.Session()

def connect_db():
    return mysql.connector.connect(
        host="localhost",  # 数据库主机地址
        user="root",  # 数据库用户名
        password="spln13spln",  # 数据库密码
        database="paper_submission"  # 数据库名
    )
db = connect_db()
cursor = db.cursor()
cursor.execute("DROP TABLE IF EXISTS conferences_temp;")
cursor.execute("CREATE TABLE conferences_temp LIKE conferences;")
db.commit()

# 设置重试策略
retry_strategy = Retry(
    total=3,  # 总重试次数
    status_forcelist=[429, 500, 502, 503, 504],  # 对哪些状态码进行重试
    allowed_methods=["HEAD", "GET", "OPTIONS"]  # 对哪些方法进行重试
)
adapter = HTTPAdapter(max_retries=retry_strategy)
session = requests.Session()
session.mount("http://", adapter)
session.mount("https://", adapter)

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

def parse_datetime(datetime_str):
    if not datetime_str or datetime_str.strip() == '':
        return None  # 如果字符串为空或只包含空格，则返回 None
    try:
        # 在解析前去除字符串两端的空格
        cleaned_date_str = datetime_str.strip()
        return datetime.strptime(cleaned_date_str, '%Y-%m-%d')
    except ValueError:
        return None  # 如果日期格式不正确，则返回 None
    
def get_conference(url):
    try:
        response = session.get(url, headers=headers, timeout=60)  # 增加了超时设置
        if response.status_code == 200:
            parse = etree.HTML(response.text)  #解析网页  
            conferenceName = parse.xpath('//*[@id="yw0"]/div[2]/div/h5/text()')[0]
            parts = conferenceName.split(': ')
            abbreviation = parts[0].split(' ')[0] 
            fullname = parts[1] if len(parts) > 1 else ''
            link = safe_xpath_search(parse, '//*[@id="yw0"]/div[2]/div/a/text()')
            deadline = safe_xpath_search(parse, '//*[@id="yw0"]/div[2]/div/table/tr[1]/td[2]/div/text()')
            notificationDate = safe_xpath_search(parse, '//*[@id="yw0"]/div[2]/div/table/tr[2]/td[2]/div/text()')
            meetingDate = safe_xpath_search(parse, '//*[@id="yw0"]/div[2]/div/table/tr[3]/td[2]/div/text()')
            meetingVenue = safe_xpath_search(parse, '//*[@id="yw0"]/div[2]/div/table/tr[4]/td[2]/div/text()')
            sessions = safe_xpath_search(parse, '//*[@id="yw0"]/div[2]/div/table/tr[5]/td[2]/div/span/text()')
            deadline = parse_datetime(deadline)
            notificationDate = parse_datetime(notificationDate)
            meetingDate = parse_datetime(meetingDate)
            sessions = safe_convert_to_int(sessions)
            # ccf = safe_xpath_search(parse, '//*[@id="yw0"]/div[2]/div/div/span[1]/text()')
            ccf = parse.xpath('//*[@id="yw0"]/div[2]/div/div//text()')
            ccf = [text.strip() for text in ccf if text.strip()]
            if ccf[0] == 'CCF:':
                ccf = ccf[1]
            else:
                ccf = ''
            conferenceInfo = safe_xpath_search(parse, '//*[@id="yw1"]/div[2]/pre/text()')
            data = (ccf, abbreviation, fullname, link, deadline, notificationDate, meetingDate, meetingVenue, sessions, conferenceInfo)
            return data
        else:
            print(f"连接失败，状态码：{response.status_code}")
    except requests.exceptions.RequestException as e:
        print(f"请求出现异常：{e}")
        return ""
    
# 循环爬取会议
pages = 5000  #5000
base_url = 'https://www.myhuiban.com/conference/' 
for page_number in tqdm(range(1, pages + 1), desc="爬取会议"):
    url = base_url + str(page_number)
    data = get_conference(url)
    if not data:  
        print(f"第 {page_number} 页为空，跳过...")
        continue  
    cursor.execute("INSERT INTO conferences_temp (ccf_ranking, abbreviation, full_name, link, material_deadline, notification_date, meeting_date, meeting_venue, sessions, info) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", data)
    db.commit()
    sleep(random.uniform(0.05, 0.2))
    
try:
    cursor.execute("DROP TABLE IF EXISTS conferences;")
    cursor.execute("RENAME TABLE conferences_temp TO conferences;")
    db.commit()
    print("数据库更新成功！")
except mysql.connector.Error as err:
    print("执行数据库操作失败：", err)
finally:
    cursor.close()
    db.close()
