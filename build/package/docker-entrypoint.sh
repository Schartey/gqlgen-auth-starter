#!/usr/bin/env bash

pid=0

int_proc() {
  if [ $pid -ne 0 ]; then
    kill -SIGINT "$pid"
    wait "$pid"
  fi
  exit 143;
}

# docker sends sigterm on ctrl-c
# setup handler
# on callback, kill the last background process, which is `tail -f /dev/null` and execute the specified handler
trap 'kill ${!}; int_proc' SIGINT
trap 'kill ${!}; int_proc' SIGTERM

# setup application
FOLDER=/app/graphql
FILE=/app/graphql/gqlgen.yml

if [ ! -d "$FOLDER" ]; then
    echo 'Creating gqlgen folder'
    mkdir $FOLDER
fi

if [ ! -f "$FILE" ]; then
    cd /app/gqlgen/
    echo 'Initializing gqlgen'
    go run github.com/99designs/gqlgen init
    echo 'Finished Initilaizing'
fi

cd /app

#run application
echo 'Running'
exec reflex -c /reflex/reflex.conf &
pid="$!"

# wait forever
while true
do
  tail -f /dev/null & wait ${!}
done