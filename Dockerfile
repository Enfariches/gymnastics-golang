FROM golang:alpine

WORKDIR /rusgidro-gym-backend-go
COPY ./ /rusgidro-gym-backend-go

RUN go mod download
RUN go get .
RUN go build
RUN go install

ENTRYPOINT [ "health", "api" ]