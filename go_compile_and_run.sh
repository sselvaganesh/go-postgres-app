#!/bin/bash

GoModStatus=`go env GO111MODULE`
echo "info: go module flag is \"$GoModStatus\""

exitIfNotSuccessful(){
	if [[ "$?" -ne "0" ]]; then
		echo "exiting."
		exit
	fi
}

dir=`pwd`
AppName=${dir##*/}
echo "compiling go lang app \"$AppName\""
echo "---------- test script execution -----------------"
go clean -testcache && go test ./...
exitIfNotSuccessful
echo "--------------------------------------------------"

#Build the package
echo "building the package..."
go build
exitIfNotSuccessful

OldBinary="$GOPATH/bin/$AppName"

if [[ -x "$OldBinary" ]]; then
	echo "removing the old executable $OldBinary"
	rm "$OldBinary"
else
	echo "$OldBinary executable not found."
fi

#Install the package
go install
exitIfNotSuccessful

echo "new binary \"$AppName\" has been installed. ($GOBIN/$AppName)"
echo "---------- app execution ---------------------------"
$GOBIN/$AppName
echo "----------------------------------------------------"
