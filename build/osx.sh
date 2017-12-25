#!/bin/bash

#Bonfire OS

# Mac OSX .app builder
cd ../front
npm run build
cd dist
#go-bindata-assetfs ./...
#cp bindata_assetfs.go ../../api

cd ../../api
go build *go

cd ../build
./setup.sh Bonfire bonfire512.png

