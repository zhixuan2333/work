#!/bin/bash

cd /minecraft

git clone git@github.com:zhixuan666/mc-heroku.git

echo "$(ls /minecraft/mc-heroku)"

echo | frpc -f 4d5d9503a4164891:852776 &

java -Xms1024M -Xmx1024M -jar server.jar nogui

echo "server is stop"

git add /minecraft/mc-heroku
git commit -am "Last sync at ${time}"
git push origin master
echo "all save"
