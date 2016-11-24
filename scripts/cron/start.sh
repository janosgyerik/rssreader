#!/bin/sh -e
# 
# File: start.sh
# Purpose: start rssreader for each configuration, if not already running
#

cd "$(dirname "$0")"
. ./common.sh

test -x rssreader || fatal "missing or not executable $PWD/rssreader"

names=$(names)
test "$names" || fatal "missing configurations in $PWD/conf/*.yml"

match_session() {
    ./list.sh | grep -F .$prefix-$1
}

for name in $names; do
    config=conf/$name.yml

    # stop running session, unless it matches "Attached" or "Detached"
    if ! match_session $name | awk '$0 !~ /tached/ { exit 1 }'; then
        ./stop.sh $name
    fi

    if ! match_session $name >/dev/null; then
        echo \* starting $name ...
        screen -d -m -S $prefix-$name ./single.sh $name
    fi
done
