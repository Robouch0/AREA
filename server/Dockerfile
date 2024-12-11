FROM golang:1.23.4-bullseye

WORKDIR /area

COPY . .

RUN go mod download

RUN go build -o ./area-server

# Modify the env to the prod one
RUN mv .env.prod .env

EXPOSE 8080

CMD [ "./area-server" ]
