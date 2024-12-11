FROM golang:1.23.4-bullseye

WORKDIR /area/

COPY . /area/

RUN go mod download

RUN go build -o ./area-server

EXPOSE 8080

CMD [ "./area-server" ]
