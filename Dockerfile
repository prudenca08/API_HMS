
FROM golang:1.17-alpine AS builder
WORKDIR /finalproject
COPY . .
RUN go mod download
RUN go build -o main main.go

#RUN STAGE
FROM alpine:3.14 
WORKDIR /finalproject

RUN mkdir config
COPY --from=builder /finalproject/config/config.json config
COPY --from=builder /finalproject/main .
EXPOSE 8080

CMD ["/finalproject/main"]

