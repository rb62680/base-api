# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/pushpaldev/base-api

# Change workdir
WORKDIR /go/src/github.com/pushpaldev/base-api

# Install the dependencies
RUN go get -t -v ./...

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/pushpaldev/base-api

ENV GIN_MODE release
ENV BASEAPI_ENV prod

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/base-api

# Document that the service listens on port 4000.
EXPOSE 4000

