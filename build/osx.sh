#!/bin/bash

#Bonfire OS

# Mac OSX .app builder
cd ../front
npm run build
cd dist
#go-bindata-assetfs ./...
#cp bindata_assetfs.go ../../api

cd ../../api
go build main.go routes.go assets_vfsdata.go

cd ../build
./setup.sh Bonfire bonfire512.png

