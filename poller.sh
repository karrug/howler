#!/bin/sh


while true
do
  notify-send "$(curl http://m.karrug.com/g)"
  sleep 5
done
