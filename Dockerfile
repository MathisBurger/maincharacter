FROM golang:1.21-alpine AS build
WORKDIR /build

# Copy source files
COPY . .
# Get go packages
RUN go mod download
# Build shinpuru backend
RUN go build -o ./bin/main cmd/main.go

FROM alpine:3 AS final

WORKDIR /app

COPY --from=build /build/bin .
ENV MAINCHARACTER_BOTTOKEN=token
ENTRYPOINT ["/app/main"]