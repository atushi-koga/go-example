FROM golang:1.21.0

ENV GOTOOLCHAIN=local
WORKDIR /go/src

COPY . .
RUN go mod download && go mod verify
RUN go build main.go

CMD ["./main"]
