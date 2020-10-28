#-*- conding = utf-8 -*-
import requests
import re
import wget
import os
from time import time
from multiprocessing.dummy import Pool as ThreadPool

items = []
url   = "http://youku.cdn-163.com/20180507/6910_65bfcd86/"

def index_m3u8():
    print('下载文件列表')
    open('index.m3u8', 'wb').write(requests.get(url + 'index.m3u8').content)
    sortdata()

def jump():
    f = open('index.m3u8', 'r')
    line = f.readline().replace('\n', '')
    if line[ -5 : ] == ".m3u8" :
        global url
        jump_url = url + line
        url = url + line[ : -10 ]
        open('index.m3u8', 'wb').write(requests.get(jump_url).content)
        sortdata()
        print('jump')
        return False
    else :
        print('no jump')
        return True

def sortdata():
    print('读取文件')
    code = open('index.m3u8', 'r').read()
    open('index.m3u8', 'w').write(re.sub(r'(?m)^ *#.*\n?', '', code))

def urls():
    f = open('index.m3u8','r')
    line = f.readline().replace('\n', '')
    while line:
        urll = url + line
        path = './temp/' + line
        item = (path, urll)
        items.append(item)
        line = f.readline().replace('\n', '')

def download(items):
    path, url = items
    wget.download(url, path)

def mix():
    os.system("copy /B .\\temp\*.ts .\\great.mp4")
    os.system("del /F /S /Q .\\temp\\*.ts")

def main():
    index_m3u8()
    jump()
    urls()
    ThreadPool(9).map(download, items)
    mix()

main()
