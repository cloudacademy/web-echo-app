FROM golang:1.19.0-bullseye as builder

WORKDIR /go/src/webapp/
COPY main.go ./
COPY go.mod ./

RUN CGO_ENABLED=0 GOOS=linux go build -o webapp .

FROM scratch

COPY --from=builder /go/src/webapp/webapp /go/bin/

CMD ["/go/bin/webapp"]