#!/bin/bash

# Remove temporary images that `docker build` leaves behind.
docker images -f 'label=io.hockeyeth.temp=true' -q | xargs docker rmi
