#!/bin/sh -e
# 
# File: stop.sh
# Purpose: stop rssreader processes
#

cd "$(dirname "$0")"
. ./common.sh

test $# = 0 && set -- $(names)

for name; do
    info stopping $name ...
    screen -S $prefix-$name -wipe >/dev/null || :
    screen -S $prefix-$name -p 0 -X stuff 
done
