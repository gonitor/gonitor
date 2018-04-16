FROM golang:1.10 

RUN mkdir /app
RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

EXPOSE 9000

RUN go build -o /app/main .
CMD ["/app/main"]
