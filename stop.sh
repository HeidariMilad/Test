#!/bin/bash

PID=$(pgrep -f "easy-socks5")

if [ -z "$PID" ]; then
    echo "easy-socks5 is not running."
else
    echo "Stopping easy-socks5 (PID: $PID)..."
    kill $PID
    sleep 1
    if pgrep -f "easy-socks5" > /dev/null; then
        echo "Forcibly killing easy-socks5..."
        kill -9 $PID
    fi
    echo "Stopped."
fi
