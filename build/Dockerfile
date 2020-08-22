FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o app .

FROM ubuntu:latest

WORKDIR /app
COPY --from=builder /app/app ./
COPY --from=builder /etc/ssl /etc/ssl

#WORKDIR /app/configs
COPY ./configs/config.yaml ./configs/

# Command to run the executable
CMD ["./app"]
