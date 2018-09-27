#!/bin/sh

rm /go/bin/pizza
cd /go/src/github.com/class/pizza
go install
/go/bin/pizza start
