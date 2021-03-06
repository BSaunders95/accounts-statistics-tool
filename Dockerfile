FROM golang:1.12-alpine

RUN apk add --no-cache git

ARG transactions_mongodb_url
ENV TRANSACTIONS_MONGODB_URL=$transactions_mongodb_url

ARG transactions_mongodb_database
ENV TRANSACTIONS_MONGODB_DATABASE=$transactions_mongodb_database

ARG log_level
ENV LOG_LEVEL=$log_level

WORKDIR /app
COPY . .
RUN GO111MODULE=on go build

CMD ["/app/accounts-statistics-tool"]
