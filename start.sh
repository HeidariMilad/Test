#!/bin/bash

# Stop existing instance if running
./stop.sh

echo "Starting easy-socks5 in background..."
nohup easy-socks5 > socks5.log 2>&1 &

sleep 1
./status.sh
