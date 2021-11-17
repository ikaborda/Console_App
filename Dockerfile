FROM golang:1.17

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 80/tcp

CMD ["cmd"]