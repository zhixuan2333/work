#-*- conding = utf-8 -*-
import requests
import re
import wget

# 原地址的m3m8  
# http://baidu.com-l-baidu.com/20190121/10957_e8fc01c1/1000k/hls/index.m3u8

# url: 这个m3m8的所有.ts集合
# urll: url中.ts集合的共同部分
urll = 'http://youku.cdn-163.com/20180507/6910_65bfcd86/1000k/hls/b3667a0141'


with open(r"great.mp4","wb") as mp4:
    # ts地址尾部的集合
    for a in range(8000,81509):
        url = urll + str(a) + '.ts'
        print(url)
        #请求.ts的地址
        resp = requests.get(url,stream=True) #通过流的方式来请求MP4
        for c in resp.iter_content(chunk_size = 1024*1024):#可以边下载边存到硬盘中
            if c:
                mp4.write(c)
    print("解析完成")