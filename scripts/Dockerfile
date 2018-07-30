FROM python:3.6-slim-jessie


ENV GO_VERSION 1.10

RUN apt-get update && \
  apt-get install -y g++ pkg-config ca-certificates wget && \
  apt-get clean && \
  rm -fr /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN  wget -qO- https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz \
    | tar -C /usr/local -xz && \
    export PATH="/usr/local/go/bin:$PATH" && \
    go version; 

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH
