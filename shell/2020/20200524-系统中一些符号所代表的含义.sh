#!/usr/bin/env bash

# Example: ./filename.sh a b c d

clear

printf "Number: %s\n" $#
printf "Filename: %s\n" $0
printf "Argument 1: %s\n" $1
printf "Argument 2: %s\n" $2
printf "Argument 3: %s\n" $3
printf "Argument 4: %s\n" $4
echo "Arguments:" $@
echo "Arguments: $*"
printf "Show process id: %s\n" $$
printf "Show last command execute status: %s\n" $?
