FROM golang:1.13.8-alpine3.11 AS builder
ENV GO111MODULE=on
WORKDIR /sender
COPY . .
RUN go build \
    -o /bin/sender \
    ./main.go

FROM alpine:3.11
COPY --from=builder /bin/sender /bin/sender
CMD ["/bin/sender"]
