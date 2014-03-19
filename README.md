Site for hackersandpainters.sg
==

[![Build Status](https://drone.io/github.com/hackerspainters/hp/status.png)](https://drone.io/github.com/hackerspainters/hp/latest)

Getting started: set up your development environment to contribute
==

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
