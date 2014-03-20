Site for hackersandpainters.sg
==

[![Build Status](https://drone.io/github.com/hackerspainters/hp/status.png)](https://drone.io/github.com/hackerspainters/hp/latest)

Getting started: Dev environment
==

You will also need to install mongodb.  This project uses the mgo library to persist data.

```bash
# Prepare directory and source code
mkdir -p $HOME/go/{src,bin,pkg}
cd $HOME/go/src
git clone git@github.com:hackerspainters/hp.git

# Set environment variables
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
export PROJ_PATH=$HOME/go/src/hp

# Get dependencies, build and run tests
cd $PROJ_PATH
go get ./...
go build ./...
go test ./...

# Run the project
go run main.go

# Then, browse to http://localhost:3000
```

Contributions and Deployment
==

Contributions are more than welcome and pull requests should be made to the `develop` branch.  Contributors should write unit tests.

Deployment is completely automated in a continuous deployment set-up with [drone.io](https://drone.io/github.com/hackerspainters/hp) on `master` branch.

Example deploy script used for such a deployment
==

```bash
#! /usr/bin/bash
export HOME=/home/web
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
echo $GOPATH
echo $PATH
mkdir -p $GOPATH/{src,bin,pkg}

export PROJ_PATH=$GOPATH/src/hp
rm -rf $PROJ_PATH
cp -rf $HOME/hackerspainters $PROJ_PATH

echo "Delete existing config.json"
rm $PROJ_PATH/config.json
echo "Use production server's config.json"
cp $HOME/config.json $PROJ_PATH/

cd $PROJ_PATH
pwd

go get ./...
go install ./...

echo "Installation completed. Restarting service..."
sudo systemctl restart hp

echo "Clean up deployment directory"
rm -rf $HOME/hackerspainters/*
```
