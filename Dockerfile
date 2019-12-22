FROM golang:alpine as build
RUN apk update && apk add git
RUN mkdir /go/pkg
WORKDIR $GOPATH/src
RUN git clone --single-branch --branch Dev https://github.com/mbedgar/F1.git
WORKDIR $GOPATH/src/F1
RUN go get github.com/influxdata/influxdb1-client/v2
RUN CGO_ENABLED=0 GOOS=linux go build -o .
#ENTRYPOINT ["./F1"]

FROM scratch
COPY --from=build go/src/F1/F1 /
ENTRYPOINT ["./F1"]
