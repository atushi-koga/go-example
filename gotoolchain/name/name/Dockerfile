FROM golang:1.21.0

ENV GOTOOLCHAIN=go1.22.0
WORKDIR /go/src

COPY .. .
RUN go mod download && go mod verify
RUN go build main.go

CMD ["./main"]
