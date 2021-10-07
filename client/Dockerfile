FROM golang:alpine as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o main

FROM alpine

RUN mkdir /app
WORKDIR /app

COPY --from=builder /buil/ /app/

CMD ["./main"]