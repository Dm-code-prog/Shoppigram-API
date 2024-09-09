FROM --platform=linux/arm64 golang as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api/main.go ./cmd/api/environment.go ./cmd/api/adapters.go

FROM --platform=linux/arm64 alpine:latest

COPY --from=builder /app/main /app/main

CMD ["/app/main"]