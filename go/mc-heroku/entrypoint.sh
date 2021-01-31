#!/bin/bash

echo | frpc -f 4d5d9503a4164891:852776 &

echo | /minecraft/mc-heroku &

sleep 15

java -Xms1024M -Xmx1024M -jar server.jar nogui