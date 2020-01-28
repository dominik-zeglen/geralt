FROM golang:1.12.7-stretch AS builder

ENV GOBIN /go/bin
ENV GO111MODULE on 
ENV CGO_ENABLED 0

RUN mkdir /app/
ADD . /app/
WORKDIR /app/

RUN go build main.go
RUN go build -o migrate migrations/main.go

FROM alpine
WORKDIR /app
RUN mkdir /app/app
COPY --from=builder /app/main /app/main
COPY --from=builder /app/migrate /app/migrate

CMD ["/app/main"]
