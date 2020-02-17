#!/bin/sh

set -e

if [ "$APP_ENV" = 'production' ]; then
  fupp-api
else
  go get github.com/pilu/fresh
  fresh
fi
