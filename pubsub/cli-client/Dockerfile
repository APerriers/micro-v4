FROM golang:alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/cli-client
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make tidy build

FROM scratch
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/cli-client/cli-client /cli-client
ENTRYPOINT ["/cli-client"]
CMD []
