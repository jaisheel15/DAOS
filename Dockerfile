
# Build stage
FROM golang:latest AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG APP=api

RUN CGO_ENABLED=0 GOOS=linux \
    go build -trimpath -ldflags="-s -w" \
    -o /out/app ./cmd/${APP}

# Runtime stage
FROM gcr.io/distroless/static-debian13:nonroot

COPY --from=builder /out/app /app

EXPOSE 8080

ENTRYPOINT ["/app"]
