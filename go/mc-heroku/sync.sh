#!/bin/bash

echo "Sync at now"

git add /minecraft/app/mc-heroku
git commit -am "Last sync at ${time}"
git push origin master

echo "Sync done"
