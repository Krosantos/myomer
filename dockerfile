FROM golang:1.14-alpine as base

WORKDIR /app/

COPY . /app/
RUN go build -o /bin/myomer .

CMD [ "/bin/myomer" ]
