FROM golang:alpine3.16

COPY . ./app
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

ENTRYPOINT [ "go" , "run" , "." ]