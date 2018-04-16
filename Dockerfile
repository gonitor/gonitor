FROM golang:1.10 

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

EXPOSE 9000

RUN go build -o main . 
CMD ["/app/main"]
