#!/bin/sh
ulimit -n 1024
./bindata.sh
goemon -c goemon.yml | jsx -w --harmony src/jsx/ assets/js/
