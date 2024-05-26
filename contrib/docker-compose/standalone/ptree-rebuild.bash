#!/bin/bash

set -eu

docker-compose \
	run --rm --entrypoint /bin/bash hockeyeth \
		-x -c '/hockeyeth/bin/hockeyeth-pbuild -config /hockeyeth/etc/hockeyeth.conf'
