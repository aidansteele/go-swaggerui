#!/bin/sh
set -eux

# TODO: fetch master upstream and change submodule ref

mkdir slim
cp \
  upstream/dist/swagger-ui-bundle.js \
  upstream/dist/swagger-ui-standalone-preset.js \
  upstream/dist/swagger-ui.css \
  upstream/dist/index.html \
  slim

packr -z
go test

git add *-packr.go
git commit -m 'Updated upstream'
