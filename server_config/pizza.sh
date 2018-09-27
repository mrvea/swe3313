#!/bin/bash
# this is the startup script for pizzad


WORK_DIR=$(dirname "${BASH_SOURCE[0]}")
cd ${WORK_DIR}

source pizza.env

PIDFILE_DIR="/var/run/pizzad/${PIZZA_ENVIRONMENT}/"
PIDFILE="${PIDFILE_DIR}/pizzad.pid"
mkdir -p "${PIDFILE_DIR}"

LOG_DIR="/var/log/pizza/${PIZZA_ENVIRONMENT}/"
DEBUG_LOG="${LOG_DIR}debug.log"
INFO_LOG="${LOG_DIR}info.log"
export EYELOG_DEBUG_OUT="${DEBUG_LOG}"
export EYELOG_INFO_OUT="${INFO_LOG}"
mkdir -p $LOG_DIR

COMMAND="pizza start"

if [ -f "${PIDFILE}" ]; then
    echo "pizzad seems to be already running." 2>&1
    exit 0
fi

exec ./$COMMAND&
echo $! > $PIDFILE
echo "Started server with PID: $!" 2>&1
exit $?
