FROM golang:alpine as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o kaiji-admin .

FROM alpine

WORKDIR /app
COPY --from=builder /app/app ./

COPY ./configs/config.yaml ./configs/

# Command to run the executable
CMD ["./kaiji-admin"]
