FROM golang:1.14-alpine

WORKDIR /go/src/datawriter
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["datawriter"]