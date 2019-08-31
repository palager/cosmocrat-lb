FROM golang:alpine as builder

WORKDIR /go/src/tcpproxy
RUN apk --no-cache add git
RUN go get github.com/palager/cosmocrat-lb/cmd/tcpproxy

FROM alpine:latest
COPY --from=builder /go/bin/tcpproxy /bin

CMD ["/bin/tcpproxy"]

