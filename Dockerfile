FROM golang:1.20.1-alpine3.17 as builder
COPY go.mod go.sum /go/src/ticket/
WORKDIR /go/src/ticket
RUN go mod download
COPY . /go/src/ticket
RUN GOOS=linux go build -a -o /go/src/ticket/build/ticket /go/src/ticket

FROM alpine
COPY --from=builder /go/src/ticket/build/ticket /usr/bin/ticket
COPY --from=builder /go/src/ticket/.env $HOME/.env
EXPOSE 3000 3000
ENTRYPOINT ["/usr/bin/ticket"]
