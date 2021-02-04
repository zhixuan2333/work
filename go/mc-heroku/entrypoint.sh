#!/bin/bash

cd /minecraft


wget --header="Authorization: token $TOKEN" -O - \
   https://api.github.com/repos/$REPO/tarball/

wget --header="Authorization: token <OAUTH-TOKEN>" -O - \
    https://api.github.com/repos/<owner>/<repo>/tarball/<version> | \
    tar xz --strip-components=1 && \
    cp -r <dir1> <dir2> ... <dirn> <destination-dir>/

# git clone git@github.com:zhixuan666/mc-heroku.git --depth 1

echo "$(ls /minecraft/app/mc-heroku)"

echo | frpc -f $FRP_TOKEN &

java -Xms1024M -Xmx1024M -jar /minecraft/app/mc-heroku/server.jar nogui

echo "server is stop"

git add /minecraft/minecraft/app/mc-heroku
git commit -am "Last sync at ${time}"
git push origin master
echo "all save"
