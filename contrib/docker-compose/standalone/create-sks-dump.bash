#!/bin/bash

set -eu

PGP_EXPORT=$(awk -F= '/^PGP_EXPORT=/ { print $2 }' .env | tail -1)

# bring deployment back up regardless of failures
cleanup() {
    docker-compose up -d
}
trap cleanup EXIT

docker-compose stop hockeyeth
docker-compose rm -f hockeyeth
docker-compose run --rm \
    --volume "${PGP_EXPORT:-/var/cache/hockeyeth}:/hockeypuck/export" \
    --entrypoint /bin/bash \
    hockeyeth -xe -c \
		'mkdir -p /hockeyeth/export/dump; find /hockeyeth/export/dump -type f -exec rm {} +; /hockeyeth/bin/hockeyeth-dump -config /hockeyeth/etc/hockeyeth.conf -path /hockeyeth/export/dump'
