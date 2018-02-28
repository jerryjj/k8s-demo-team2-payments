#### build stage ####
FROM golang:1.9 as golang

WORKDIR /build

ENV GOPATH=/build

ADD ./src ./src

RUN cd src/qvik.fi/payments-service && go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/payments-service qvik.fi/payments-service

#### run stage ####
FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=golang /build/bin/payments-service /usr/local/bin/payments-service

EXPOSE 8080

CMD ["/usr/local/bin/payments-service"]
