FROM golang:1.23.5-alpine AS userreviewbuilder
RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o userReviewServices ./cmd/app


FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary
COPY --from=userreviewbuilder /build/userReviewServices /app/userReviewServices

CMD ["./app/userReviewServices"]