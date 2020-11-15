FROM golang:1.12.7-stretch AS builder

ENV GOBIN /go/bin
ENV GO111MODULE on 
ENV CGO_ENABLED 0

RUN mkdir /app/
ADD . /app/
WORKDIR /app/

RUN go build -o bin/main main.go
RUN go build -o bin/migrate migrations/*.go
RUN go build -o bin/client client/main.go

FROM alpine
WORKDIR /app
RUN mkdir /app/app
COPY --from=builder /app/bin/main /app/main
COPY --from=builder /app/bin/migrate /app/migrate
COPY --from=builder /app/bin/client /app/client

CMD ["/app/main"]
