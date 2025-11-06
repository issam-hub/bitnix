FROM golang:1.23.5-alpine AS gamescatalogbuilder
RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o gameCatalogServices ./cmd/app

# Install migrate tool
RUN apk add --no-cache curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/migrate && \
    chmod +x /usr/local/bin/migrate

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary
COPY --from=gamescatalogbuilder /build/gameCatalogServices /app/gameCatalogServices

# Copy migrate tool
COPY --from=gamescatalogbuilder /usr/local/bin/migrate /usr/local/bin/migrate

# Copy migrations
COPY --from=gamescatalogbuilder /build/migrations /app/migrations

# Create entrypoint script
COPY <<'EOF' /app/entrypoint.sh
#!/bin/sh
set -e

echo "Running database migrations..."
migrate -path=/app/migrations -database="$DSN" up

echo "Starting application..."
exec /app/gameCatalogServices
EOF

RUN chmod +x /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]