#!/bin/bash

# bash start.sh --name --instance --port -github --database_host docker.for.mac.host.internal --database_name kodit_manager --database_username postgres --database_password --auth_host --auth_realm --auth_client --auth_client_secret --auth_public_key 

echo "###############################################################"
echo "#  _____ _          __  __              ___         _         #"
echo "# |_   _| |_  ___  |  \/  |_  _ __ _   / __|___  __| |___ ___ #"
echo "#   | | | ' \/ -_) | |\/| | || / _/`` | | (__/ _ \/ _``  / -_|_-< #"
echo "#   |_| |_||_\___| |_|  |_|\_,_\__, |  \___\___/\__,_\___/__/ #"
echo "#                              |___/                          #"
echo "###############################################################"

echo "# 1 - Prepring to generate docs"
export GOPATH=$HOME/go && export PATH=$PATH:$GOPATH/bin

echo "# 2 - Generating docs files"
swag init --parseDependency --parseInternal --generatedTime 

echo "# Finished all steps"