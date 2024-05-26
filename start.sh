#!/bin/bash

# bash start.sh --name --instance --port -github --database_host docker.for.mac.host.internal --database_name kodit_manager --database_username postgres --database_password --auth_host --auth_realm --auth_client --auth_client_secret --auth_public_key 

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

echo "# 1 - Build image for $name"
docker build \
-t the-mug-codes/$name \
--build-arg APP_NAME=$name \
--build-arg GITHUB_TOKEN=$github \
--no-cache .

echo "# 1 - Check and stop $name instance $instance"
docker stop  $name-$instance || true && docker rm $name-$instance || true

echo "# 2 - Start $name instance $instance"
docker run -d --name $name-$instance \
-p $port:80 \
-e INSTANCE=$instance \
-e DATABASE_HOST=$database_host \
-e DATABASE_NAME=$database_name \
-e DATABASE_USERNAME=$database_username \
-e DATABASE_PASSWORD=$database_password \
-e AUTH_HOST=$auth_host \
-e AUTH_REALM=$auth_realm \
-e AUTH_CLIENT=$auth_client \
-e AUTH_CLIENT_SECRET=$auth_client_secret \
-e AUTH_PUBLIC_KEY=$auth_public_key \
--restart on-failure:3 \
the-mug-codes/$name:latest

echo "# Finished all steps"