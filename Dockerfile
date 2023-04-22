FROM golang:alpine

ENV GIN_MODE=release

WORKDIR /go/src/geojson-api
COPY . .

RUN apk update && apk add --no-cache git
RUN go get

RUN go build

EXPOSE 3000

ENTRYPOINT ["./geojson-api"]