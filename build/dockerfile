FROM golang:alpine

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum main.go ./
RUN go mod download && go mod verify

RUN go build -v -o bin .

ENTRYPOINT ["/usr/src/app/bin"]
CMD ["app"]
