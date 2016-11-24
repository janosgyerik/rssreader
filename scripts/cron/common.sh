#!/bin/sh -e
#
# File: common.sh
# Purpose: common functions
#

prefix=rssreader

info() {
    echo '[info]' $@
}

result() {
    echo '[result]' $@
}

warn() {
    echo '[WARN]' $@
}

fatal() {
    echo '[fatal]' $@
    exit 1
}

names() {
    for config in conf/*.yml; do
        test -f "$config" || continue
        name=${config#conf/}
        name=${name%.yml}
        echo $name
    done
}    
