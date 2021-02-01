#!/bin/bash

git init
git remote add origin git@github.com:zhixuan666/mc-heroku.git
git fetch
git pull origin master

echo "$(ls /minecraft/app)"

echo | frpc -f 4d5d9503a4164891:852776 &

java -Xms1024M -Xmx1024M -jar server.jar nogui

ehco "server is stop"
git add /minecraft/app
git commit -am "Last sync at ${time}"
git push origin master
ehco "all save"
