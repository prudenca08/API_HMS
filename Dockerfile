# FROM golang:1.17-alpine3.14

# WORKDIR /hms-app

# COPY go.mod ./
# COPY go.sum ./

# RUN go mod download


# COPY . .

# RUN go build -o mainfile

# EXPOSE 8080

# CMD [ "./mainfile" ]

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

