#!/usr/bin/env bash

clear

for fgbg in 38 48; do
  for color in {0..256}; do
    printf "\033[%s;5;%sm [%03s] \t" $fgbg $color $color
    echo -en "\033[0m"
    if [ $((($color + 1) % 10)) == 0 ]; then
      echo
    fi
  done
  echo
  echo
done
