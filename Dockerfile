FROM golang:latest 

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 

RUN go get github.com/gin-gonic/gin
RUN go get github.com/kaka-crawler/service-reddit

EXPOSE 9000

RUN go build -o main . 
CMD ["/app/main"]
