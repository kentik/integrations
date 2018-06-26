#!/bin/bash

ver=`kbt deb-version 2>/dev/null`
if test $? -ne 0; then
	ver=`git rev-parse HEAD`
	if test -n "$(git status --porcelain)"; then
		ver="dirty-$ver"
	fi
fi

date=`git show -s --format=%cD`
platform=`uname -srm`
distro=`lsb_release -d | cut -f2`
golang=`go version`

# only create the version.go files if they've changed, to avoid unnecessary builds
TEMP_FILE=$( mktemp )
OUTPUT_PATH=cmd/version.go

cp cmd/version.go.base $TEMP_FILE
sed -i.bak "s/XXX_GIT_HASH/$ver/g" $TEMP_FILE
sed -i.bak "s/XXX_DATE/$date/g" $TEMP_FILE
sed -i.bak "s|XXX_PLATFORM|$platform|g" $TEMP_FILE
sed -i.bak "s|XXX_DISTRO|$distro|g" $TEMP_FILE
sed -i.bak "s|XXX_GOLANG|$golang|g" $TEMP_FILE

cmp -s $TEMP_FILE $OUTPUT_PATH
RET=$?
if [ $RET -ne 0 ]; then
  echo "Version changed - updating ${OUTPUT_PATH}"
  cp $TEMP_FILE $OUTPUT_PATH
fi
rm $TEMP_FILE
