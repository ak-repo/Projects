FROM golang:1.23-alpine AS build

WORKDIR /src

RUN apk add --no-cache ca-certificates git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG VERSION=dev
ARG COMMIT=none
ARG DATE=unknown

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w -X go-project-skeleton/internal/version.Version=${VERSION} -X go-project-skeleton/internal/version.Commit=${COMMIT} -X go-project-skeleton/internal/version.Date=${DATE}" \
    -o /bin/api ./cmd/api

FROM alpine:3.20

RUN apk add --no-cache ca-certificates && adduser -D -H -u 10001 appuser

USER appuser
WORKDIR /app

COPY --from=build /bin/api /app/api

EXPOSE 8080

ENTRYPOINT ["/app/api"]
