FROM golang:1.12 as builder
WORKDIR /go/src/app
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -v -o app

FROM alpine:3.6
WORKDIR /root
ENV GO_ENV production
ENV GIN_MODE release

COPY --from=builder /go/src/app/app /app

CMD [ "/app" ]