FROM golang:1.21.2

ENV GOTOOLCHAIN=go1.21.2+path
WORKDIR /go/src

COPY . .
RUN go mod download && go mod verify
RUN go build main.go

CMD ["./main"]
