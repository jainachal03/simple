FROM golang:1.19.2-alpine3.16

RUN mkdir -p /home/grpc

COPY ./ /home/grpc

WORKDIR /home/grpc 

RUN go mod tidy

ENTRYPOINT ["go", "run", "main.go"]