#!/bin/bash

cd /minecraft

curl https://$TOKEN:@api.github.com/repos/$REPO/releases/latest \
   | grep "tarball_url" \
   | cut -d '"' -f 4 \
   | curl -Lo /minecraft/mc-server.zip

# git clone git@github.com:zhixuan666/mc-heroku.git --depth 1

echo "$(ls /minecraft/app/mc-heroku)"

echo | frpc -f 4d5d9503a4164891:852776 &

java -Xms1024M -Xmx1024M -jar /minecraft/app/mc-heroku/server.jar nogui

echo "server is stop"

git add /minecraft/minecraft/app/mc-heroku
git commit -am "Last sync at ${time}"
git push origin master
echo "all save"
