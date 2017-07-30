FROM golang:1.9-alpine
COPY . /go/src/github.com/Deseao/anon
WORKDIR /go/src/github.com/Deseao/anon
RUN go build -o api
CMD ./api