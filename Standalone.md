## Installation Instructions (Standalone)

### Install Go and Setup Your Go Workspace

Make sure you're running the latest version of Go and you have set up your Go workspace properly. 

Refer to the following links for more information:

[Instructions to Install Go](https://golang.org/doc/install)

[Instructions to Setup your Go Workspace](https://golang.org/doc/code.html)


### Install GopherJS

`go get -u github.com/gopherjs/gopherjs`

Verify your installation by issuing the following command in a terminal window:
`gopherjs version`

You should see the GopherJS version.

### Install Some Really Helpful GopherJS Bindings

#### DOM Bindings
`go get honnef.co/go/js/dom`

#### JSBuiltin
`go get -u -d -tags=js github.com/gopherjs/jsbuiltin`

#### XHR
`go get -u honnef.co/go/js/xhr`

#### WebSocket
`go get -u github.com/gopherjs/websocket`


### Install IsoKit
`go get -u github.com/isomorphicgo/isokit`


### Install UXToolkit
`go get -u github.com/uxtoolkit/cog`


### Install the source code for the book
`go get -u github.com/EngineerKamesh/igb`

### Add the following entry to your .profile (or .bashrc)
`export IGWEB_APP_ROOT=${GOPATH}/src/github.com/EngineerKamesh/igb/igweb`

### Transpile the Client-Side Application
```
$ cd $IGWEB_APP_ROOT/client
$ go get ./..
$ gopherjs build
```

### Install and Run a Redis Instance Locally
```
$ cd ~/Downloads
$ wget http://download.redis.io/releases/redis-4.0.2.tar.gz
$ tar xzf redis-4.0.2.tar.gz
$ cd redis-4.0.2
$ make
$ sudo make install
$ redis-server
```

### Start the Server-Side Application
```
cd $IGWEB_APP_ROOT
go run igweb.go
```

### Load the Sample Data Set

Note: The following link will only work, once you have your local Redis instance and the IGWEB Server-Side Application running:

[Click this link to load the sample data set](http://localhost:8080/config/load-sample-data).

### Access the IGWEB Demo

You may access the IGWEB Demo at [http://localhost:8080/index](http://localhost:8080/index)
