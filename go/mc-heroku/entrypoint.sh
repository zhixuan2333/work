#!/bin/bash

cd /minecraft/app

git clone git@github.com:zhixuan666/mc-heroku.git

echo "$(ls /minecraft/app/mc-heroku)"

echo | frpc -f 4d5d9503a4164891:852776 &

java -Xms1024M -Xmx1024M -jar /minecraft/app/mc-heroku/server.jar nogui

echo "server is stop"

git add /minecraft/minecraft/app/mc-heroku
git commit -am "Last sync at ${time}"
git push origin master
echo "all save"
