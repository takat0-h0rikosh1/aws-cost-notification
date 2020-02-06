FROM golang:1.15-buster as builder

COPY ./app/ ./aws-cost-notification/app/
COPY ./go.mod ./aws-cost-notification/go.mod
WORKDIR ./aws-cost-notification
RUN go mod tidy
RUN go build -o /main ./app/*.go
RUN cp -r ./app/config/ /config

FROM debian:stretch-slim

COPY --from=builder /main .
COPY --from=builder /config/ ./config/
RUN apt-get update && apt-get install -y ca-certificates

ENTRYPOINT ["./main"]




