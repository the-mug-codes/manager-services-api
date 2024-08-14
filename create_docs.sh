#!/bin/bash

echo "###############################################################"
echo "#  _____ _          __  __              ___         _         #"
echo "# |_   _| |_  ___  |  \/  |_  _ __ _   / __|___  __| |___ ___ #"
echo "#   | | | ' \/ -_) | |\/| | || / _/`` | | (__/ _ \/ _``  / -_|_-< #"
echo "#   |_| |_||_\___| |_|  |_|\_,_\__, |  \___\___/\__,_\___/__/ #"
echo "#                              |___/                          #"
echo "###############################################################"

while [ $# -gt 0 ]; do
    if [[  $1 == "--"* ]]; then
        v="${1/--/}"
        declare "$v"="$2"
        shift
    fi
    shift
done

echo "# 1 - Prepring to generate docs for $name"
export GOPATH=$HOME/go && export PATH=$PATH:$GOPATH/bin

echo "# 2 - Generating docs files for $name"
swag init --parseDependency --parseInternal --generatedTime --quiet

echo "# Finished all steps"