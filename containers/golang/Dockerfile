FROM millisecond/millisecond:latest

MAINTAINER millisecond@gmail.com

# docker build . -t millisecond-golang && docker run --rm -it millisecond-golang /bin/bash

RUN apt-get update && apt-get install -y --no-install-recommends \
		g++ \
		gcc \
		libc6-dev \
		make \
		pkg-config \
	&& rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.8
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 53ab94104ee3923e228a2cb2116e5e462ad3ebaeea06ff04463479d7f12d27ca

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
	&& echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz

ENV GOPATH /srv/workspace
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

COPY ./conf/supervisor-*.conf /etc/supervisor/conf/

ENV GO_PACKAGE github.com/millisecond/golang
ENV SRV /srv/workspace/src/${GO_PACKAGE}

RUN mkdir -p ${SRV}

COPY . ${SRV}/

WORKDIR ${SRV}/
