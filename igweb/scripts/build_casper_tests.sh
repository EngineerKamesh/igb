#!/bin/bash
pushd .
cd $IGWEB_APP_ROOT
for gofile in $(find ./client/tests/go/*.go); do
	jsfile=${gofile//go/js}
	gopherjs build $gofile -o $jsfile | sed 's/^/    /'
done
popd .
