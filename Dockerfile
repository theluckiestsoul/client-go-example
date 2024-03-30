FROM golang:1.21.8 as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o lister .

FROM ubuntu
COPY --from=builder /app/lister /lister
ENTRYPOINT ["/lister"]
