#!/bin/bash

echo "Build Front End Assets..."
cd ../front
npm run build
cd dist

echo "Compile binary..."
cd ../../api
go build

echo "Build OSX app bundle..."
cd ../build
./setup.sh Bonfire bonfire512.png

echo "Copy binaries..."
cp ../api/api "Bonfire.app/Contents/MacOS/Bonfire"
cp ../api/binaries/darwin/ffmpeg "Bonfire.app/Contents/MacOS/ffmpeg"
