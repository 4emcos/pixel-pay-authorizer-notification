FROM golang:1.22.0 as base

WORKDIR /app
COPY . /app

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build main.go

FROM alpine
COPY go.mod go.sum ./

COPY --from=base /app/main ./

CMD ["./main"]