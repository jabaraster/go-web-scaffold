#!/bin/sh
ulimit -n 1024
goemon -c goemon.yml | ./watch-jsx.sh
