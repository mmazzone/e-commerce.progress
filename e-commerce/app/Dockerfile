FROM golang:1.15

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gorilla/mux
EXPOSE 8000

CMD ["app"]