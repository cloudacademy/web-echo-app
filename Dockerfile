FROM golang:1.19.0-bullseye as builder
RUN useradd -u 1500 cloudacademy

WORKDIR /go/src/webapp/
COPY main.go ./
COPY go.mod ./

RUN CGO_ENABLED=0 GOOS=linux go build -o webapp .

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/src/webapp/webapp /go/bin/

USER cloudacademy

ENV HOSTPORT=0.0.0.0:8080
EXPOSE 8080

CMD ["/go/bin/webapp"]