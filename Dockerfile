FROM golang:1.13.4-alpine3.10
RUN mkdir /dock
ADD . /dock
WORKDIR /dock
RUN go build -o main .
CMD ["/dock/main"]