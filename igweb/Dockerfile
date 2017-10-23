FROM golang
MAINTAINER Kamesh Balasubramanian kamesh@kamesh.com

# Declare required environment varialbes
ARG IGWEB_APP_ROOT=/go/src/github.com/EngineerKamesh/igb/igweb
ARG IGWEB_DB_CONNECTION_STRING="igweb-db:6379"
ARG IGWEB_MODE=production
ARG IGWEB_DISABLE_STATIC_ASSETS_BUNDLING=true
ARG GOPATH=/go

# Get the needed packages
RUN go get -u github.com/gopherjs/gopherjs
RUN go get -u honnef.co/go/js/dom
RUN go get -u -d -tags=js github.com/gopherjs/jsbuiltin
RUN go get -u honnef.co/go/js/xhr
RUN go get -u github.com/gopherjs/websocket
RUN go get -u github.com/isomorphicgo/isokit 
RUN go get -u github.com/uxtoolkit/cog
RUN go get -u github.com/EngineerKamesh/igb

# Transpile the client-side application code
RUN cd $IGWEB_APP_ROOT/client; go get ./..; gopherjs build -m -o $IGWEB_APP_ROOT/static/js/client.js

# Compile the server-side application code
RUN go install github.com/EngineerKamesh/igb/igweb

# Add the static assets
ADD ./static/css/cogimports.css $IGWEB_APP_ROOT/static/css/cogimports.css
ADD ./static/js/cogimports.js $IGWEB_APP_ROOT/static/js/cogimports.js
ADD ./static/templates/igweb.tmplbundle $IGWEB_APP_ROOT/static/templates/igweb.tmplbundle

# Specify the entrypoint
ENTRYPOINT /go/bin/igweb

# Expose port 8080 of the container
EXPOSE 8080
