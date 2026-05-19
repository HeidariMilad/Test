#!/bin/bash

PID=$(pgrep -f "easy-socks5")

if [ -z "$PID" ]; then
    echo "❌ easy-socks5 is NOT running."
else
    echo "✅ easy-socks5 is running with PID: $PID"
    echo "Listening on port: $(sudo netstat -tulpn | grep $PID | awk '{print $4}')"
fi

echo "--- Recent Logs (socks5.log) ---"
if [ -f socks5.log ]; then
    tail -n 20 socks5.log
else
    echo "No log file found yet."
fi
