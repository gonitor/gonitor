FROM golang:1.10-alpine

RUN apk --update upgrade \
&& apk --no-cache --no-progress add git bash \
&& rm -rf /var/cache/apk/*

RUN mkdir /app
RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

EXPOSE 9000

RUN go build -o /app/main .
CMD ["env GONI_PRODMODE=true","/app/main"]
