FROM golang:1.12.7-stretch AS builder

ENV GOBIN /go/bin
ENV GO111MODULE on 

RUN mkdir /app/
ADD . /app/
WORKDIR /app/

RUN go build main.go

FROM alpine
WORKDIR /app
RUN mkdir /app/app
COPY --from=builder /app/main /app

CMD ["/app/main"]
