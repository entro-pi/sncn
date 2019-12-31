#!/bin/bash

go build
mv client snowcrash.network

sudo cp -R ../pot /usr/bin/pot
sudo cp snowcrash.network /usr/bin/snowcrash.network

echo -e "\033[38:2:0:200:0mIf no errors occurred, run the client with 'snowcrash.network' Enjoy!\033[0m"
