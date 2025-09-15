#!/bin/bash

STATUS=$(playerctl status)
PLAY="playerctl play"
PAUSE="playerctl pause"

if [ "$STATUS" == "Playing" ]; then
  eval "$PAUSE"
elif [ "$STATUS" == "Paused" ]; then
  eval "$PLAY"
else
  notify-send "We are so fucked!"
fi
