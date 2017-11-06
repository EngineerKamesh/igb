FROM golang
MAINTAINER Kamesh Balasubramanian kamesh@kamesh.com

# Declare required environment variables
ENV IGWEB_APP_ROOT=/go/src/github.com/EngineerKamesh/igb/igweb
ENV IGWEB_DB_CONNECTION_STRING="database:6379"
ENV IGWEB_MODE=production
ENV GOPATH=/go

# Get the required Go packages
RUN go get -u github.com/gopherjs/gopherjs
RUN go get -u honnef.co/go/js/dom
RUN go get -u -d -tags=js github.com/gopherjs/jsbuiltin
RUN go get -u honnef.co/go/js/xhr
RUN go get -u github.com/gopherjs/websocket
RUN go get -u github.com/tdewolff/minify/cmd/minify
RUN go get -u github.com/isomorphicgo/isokit 
RUN go get -u github.com/uxtoolkit/cog
RUN go get -u github.com/EngineerKamesh/igb

# Build and install the minifier
RUN go install github.com/tdewolff/minify

# Transpile and install the client-side application code
RUN cd $IGWEB_APP_ROOT/client; go get ./..; /go/bin/gopherjs build -m --verbose --tags clientonly -o $IGWEB_APP_ROOT/static/js/client.min.js

# Build and install the server-side application code
RUN go install github.com/EngineerKamesh/igb/igweb

# Generate the static assets
RUN /go/bin/igweb --generate-static-assets

# Minify IGWEB's CSS stylesheet
RUN /go/bin/minify --mime="text/css" $IGWEB_APP_ROOT/static/css/igweb.css > $IGWEB_APP_ROOT/static/css/igweb.min.css

# Specify the entrypoint
ENTRYPOINT /go/bin/igweb

# Expose port 8080 of the container
EXPOSE 8080
