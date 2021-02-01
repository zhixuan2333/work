#!/bin/bash

echo "Sync at now"

echo | git add /minecraft/app

echo | git commit -am "Last sync at ${time}"
