#!/bin/sh

docker run \
  --port 8080:8080 \
  --tag trader \
  --volume /log/trader:/log:rw \
  --restart=unless-stopped \
  --detach
