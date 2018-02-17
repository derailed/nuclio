#!/bin/bash
while true; do \
  http 192.168.64.38:31418 --print=b --pretty=colors;
  sleep 2 ; \
  clear ; \
done