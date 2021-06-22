#!/bin/bash
echo "Start deploy"
ARCH=amd64
HOST=site
RHOME=/home/zf413514
WDIR=site.ai-r.info/www
#DB=/opt/ptks/oapi/db
MDIR=$PWD/deploy
TARGET=$MDIR/$WDIR
APP=gomoku

# ### Build cgi
mkdir -p ./$ARCH
GOOS=linux GOARCH=$ARCH go build -o ./$ARCH/$APP ./

# ### Build wasm
mkdir -p ./web
GOARCH=wasm GOOS=js go build -o web/app.wasm

./$ARCH/$APP

mkdir -p $MDIR

### Mount deploy
docker-machine mount $HOST:$RHOME $MDIR

### Make target dir
mkdir -p $TARGET/$APP/web

# ### Deploy cgi
#mkdir -p $TARGET/cgi-bin
cp -v -u ./*.js $TARGET/$APP/
cp -v -u ./*.css $TARGET/$APP/
cp -v -u ./*.html $TARGET/$APP/
cp -v -u ./*.webmanifest $TARGET/$APP/
cp -v -u -r ./static $TARGET/$APP/
#chmod 750 $TARGET/hello.cgi
#cp -v -u sql_private_key.pem $TARGET/../
cp -v -u ./web/app.wasm $TARGET/$APP/web/

# ### Deploy perl
# cp -v -u test.pl $TARGET/
# chmod 750 $TARGET/test.pl

# ### Deploy perl
# cp -v -u test.py $TARGET/
# chmod 750 $TARGET/test.py

# ### Deploy perl
# cp -v -u index.php $TARGET/
# chmod 750 $TARGET/index.php

### Unmount
sleep 10s
docker-machine mount -u $HOST:$RHOME $MDIR

