#!/bin/sh


old=""
while true
do
  new="$(curl http://m.karrug.com/g)"
  if [ "$new" != "$old" ]
  then
    notify-send "$new"
    old="$new"
  fi
  sleep 5
done
