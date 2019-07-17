#!/bin/bash
set -e

FOLDER=/app/gqlgen
FILE=/app/gqlgen/gqlgen.yml

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

echo 'Running'
reflex -c /reflex/reflex.conf

exec "$@"