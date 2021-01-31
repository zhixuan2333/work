#!/bin/bash

echo | frpc -f 4d5d9503a4164891:852776 &

git clone git@github.com:zhixuan666/mc-heroku.git .

java -Xms1024M -Xmx1024M -jar server.jar nogui

git add .

git commit -am "Last sync at ${time}"