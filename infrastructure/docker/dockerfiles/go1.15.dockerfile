FROM alpine:3.12

RUN apk --no-cache add ca-certificates wget
RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub
RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.28-r0/glibc-2.28-r0.apk
RUN apk add glibc-2.28-r0.apk

ARG BUILD_DATE
ARG VCS_REF

ENV GOLANG_VERSION=1.15.6
ENV PROTOC_VERSION=3.6.1

RUN apk add --no-cache --update ca-certificates wget git curl unzip openssh && \
	apk upgrade --update --no-cache

RUN set -ex && apk add --no-cache bash gcc musl-dev openssl openssl-dev alpine-sdk go
RUN export \
		# set GOROOT_BOOTSTRAP such that we can actually build Go
		GOROOT_BOOTSTRAP="$(go env GOROOT)" \
		# ... and set "cross-building" related vars to the installed system's values so that we create a build targeting the proper arch
		# (for example, if our build host is GOARCH=amd64, but our build env/image is GOARCH=386, our build needs GOARCH=386)
		GOOS="$(go env GOOS)" \
		GOARCH="$(go env GOARCH)" \
		GO386="$(go env GO386)" \
		GOARM="$(go env GOARM)" \
		GOHOSTOS="$(go env GOHOSTOS)" \
		GOHOSTARCH="$(go env GOHOSTARCH)" \
	;

RUN wget -q "https://golang.org/dl/go${GOLANG_VERSION}.src.tar.gz" -O golang.tar.gz && \
	tar -C /usr/local -xzf golang.tar.gz && \
	rm golang.tar.gz && \
	cd /usr/local/go/src && \
	for p in /go-alpine-patches/*.patch; do \
		[ -f "$p" ] || continue; \
		patch -p2 -i "$p"; \
	done; \
	\
	./make.bash && \
	\
	rm -rf /go-alpine-patches

RUN wget -q "https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip" -O protoc.zip && \
	cd /usr && \
	unzip ../protoc.zip && \
	rm -rf /protoc.zip && \
	rm -rf /usr/readme.txt && \
	apk del --purge unzip

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && \
	chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

RUN git config --global http.https://gopkg.in.followRedirects true
RUN go get -u -v github.com/golang/dep/cmd/dep


RUN go get -u google.golang.org/protobuf/cmd/protoc-gen-go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go

RUN go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

RUN go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
