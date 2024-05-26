#!/bin/bash

set -eu

docker-compose run --rm --entrypoint \
	/bin/bash hockeyeth -x -c \
		'/hockeyeth/bin/hockeyeth-load -config /hockeyeth/etc/hockeyeth.conf /hockeyeth/import/dump/*.pgp'
