FROM golang:1.9-alpine
COPY . /go/src/github.com/Deseao/anon/api
WORKDIR /go/src/github.com/Deseao/anon/api
RUN go build
CMD ./api