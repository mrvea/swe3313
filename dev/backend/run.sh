#!/bin/sh

PIDFILE=/var/run/backend

go build -o /go/bin/backend

if [ $? -eq 0 ]; then
    if [ -f $PIDFILE ]; then
        PID=$(cat $PIDFILE)
        echo "Stopping running process: $PID"
        kill -9 $PID
        rm $PIDFILE
    fi

    /go/bin/backend start&

    if [ $? -eq 0 ]; then
        echo $! > $PIDFILE
    fi
fi

