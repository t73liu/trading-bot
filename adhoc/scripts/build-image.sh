#!/bin/sh

docker build \
  --disable-content-trust false \
  --file ${TRADING_BOT_REPO}/Dockerfile \
  --tag trader .
