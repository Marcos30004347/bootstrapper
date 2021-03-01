# Build phase
FROM golang:1.15 as build

WORKDIR /go/src/github.com/marcos30004347/kubeapi
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

# Run phase
FROM alpine:latest

RUN apk --no-cache add ca-certificates

# Copy the executable
COPY --from=build /go/src/github.com/marcos30004347/kubeapi/kubeapi /

ENTRYPOINT ["/kubeapi"]
