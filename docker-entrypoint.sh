#!/bin/sh

set -e

/app/goxygen $1 $2

cd /app/generated/ 

chown -R $UID:$GID /app/generated/$2