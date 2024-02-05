FROM golang:1.22rc2

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

#make postgres.sh executable
RUN chmod +x postgres.sh

# build go app
RUN go mod download
RUN go build -o test-crud-api ./cmd/main.go

CMD ["./test-crud-api"]