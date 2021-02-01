#!/bin/bash

echo "Sync at now"
git add /minecraft/app
git commit -am "Last sync at ${time}"
git pull
