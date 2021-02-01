#!/bin/bash

echo | frpc -f 4d5d9503a4164891:852776 &

mkdir /minecraft/app

git init
git remote add origin git@github.com:zhixuan666/mc-heroku.git
git fetch
git pull origin main

echo "$(ls /minecraft)"
echo "$(ls /minecraft/app)"

java -Xms1024M -Xmx1024M -jar server.jar nogui

git add /minecraft/app
git commit -am "Last sync at ${time}"
git push origin main
