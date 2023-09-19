FROM golang:1.21

WORKDIR /COA

COPY . /COA

RUN go mod download
RUN go build server/main.go

CMD ["./main"]