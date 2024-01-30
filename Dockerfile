FROM golang:1.21-buster

RUN go version
ENV GOPATH=/

COPY ./ ./
RUN apt-get update
RUN apt-get -y install postgresql-clien
RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o golang-test-restapi ./cmd/main.go
CMD ["./golang-test-restapi"]