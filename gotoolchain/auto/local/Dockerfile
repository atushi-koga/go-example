FROM golang:1.23.0

ENV GOTOOLCHAIN=auto
WORKDIR /go/src

COPY . .
RUN go mod download && go mod verify
RUN go build main.go

CMD ["./main"]
