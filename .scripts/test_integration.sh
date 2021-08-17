#!/bin/bash

docker-compose down
docker container rm -f "$(docker container ls -q)"
docker-compose up -d --force-recreate
sleep 2
psql postgresql://connect:connect@127.0.0.1/connect < .scripts/database.sql
if [ "$1" = "" ]; then
    go test -v .
else
    go test -v -run $1
fi
docker-compose down

