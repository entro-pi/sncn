#!/bin/bash

mkdir ../pot

touch ../pot/broadcast
touch ../pot/tells
touch ../pot/who

echo Wallace >> ../pot/who
echo ::SENDER::WEASEL::SENDER::=@WEASEL@ hello ::=::SENDTO::WEASEL::SENDTO::::TIMESTAMP::Wednesday-21:17::TIMESTAMP:: >> ../pot/tells
echo ::SENDER::WEASEL::SENDER::=weasel says hello ::=::SENDTO::ALL::SENDTO::::TIMESTAMP::Wednesday-21:17::TIMESTAMP:: >> ../pot/broadcast
echo export LD_LIBRARY_PATH=/usr/local/lib64 >> ~/.bashrc
export LD_LIRBARY_PATH=/usr/local/lib64

echo -e "\033[38:2:0:200:0mIf no errors occurred, run the client with the --gui flag to continue. Enjoy!\033[0m"
