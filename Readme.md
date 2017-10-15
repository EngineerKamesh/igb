# Isomorphic Go by Kamesh Balasubramanian

This is the official source code repository for the Isomorphic Go book published by Packt. The book teaches you how to develop isomorphic web applications using the Go programming language.

[![IsomorphicGo](https://raw.githubusercontent.com/EngineerKamesh/igb/master/assets/isomorphicgo_cover_thumb.png)](https://www.packtpub.com/web-development/isomorphic-go)

ISBN: 9781788394185

## The solutions you will learn how to build in this book:

### Live Chat (with a bot)
<img border="1" width="360" src="https://raw.githubusercontent.com/EngineerKamesh/igb/master/assets/videos/livechat.gif">

#### Topics Covered
* Real-Time Web Applications
* Communicating over a WebSocket connection

### Carousel
<img border="1" width="360" src="https://raw.githubusercontent.com/EngineerKamesh/igb/master/assets/videos/carousel.gif">

#### Topics Covered
* Cogs: Reusable Components
* Building a Hybrid Cog (Go + JavaScript)

### Products 
<img border="1" width="360" src="https://raw.githubusercontent.com/EngineerKamesh/igb/master/assets/videos/products.gif">

#### Topics Covered
* End to End Application Routing
* Classic Web Server Architecture for Initial Page Load
* Single Page Application Architecture for Subsequent Interactions
* Implementing Rest API Endpoints

### Shopping Cart 
<img border="1" width="360" src="https://raw.githubusercontent.com/EngineerKamesh/igb/master/assets/videos/shoppingcart.gif">	

#### Topics Covered
* Isomorphic Handoff
* Implementing Rest API Endpoints

### About
<img border="1" width="360" src="https://raw.githubusercontent.com/EngineerKamesh/igb/master/assets/videos/about.gif">

#### Topics Covered
* Isomorphic Template Rendering
* Sharing Template Data and Functions Across Environments
* Building a Pure Cog using functionality from an existing Go package
* Implementing Rest API Endpoints

### Contact Form
<img border="1" width="360" src="https://raw.githubusercontent.com/EngineerKamesh/igb/master/assets/videos/contact.gif">

#### Topics Covered
* Isomorphic Forms
* Reusing Validation Logic
* Datepicker Hybrid Cog (Go + JavaScript)
* Implementing Rest API Endpoints

### Live Clock
<img border="1" width="360" height="180" src="https://raw.githubusercontent.com/EngineerKamesh/igb/master/assets/videos/liveclock.gif">

#### Topics Covered
* Cogs: Reusable Components
* The Virtual Dom
* Building a Pure Cog (Go only)



## Installation Instructions

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
\

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



