FROM golang:1.20.3-alpine AS BUILDER

WORKDIR /usr/src/beat

RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o /beat main.go


FROM scratch

ARG APP_VERSION 0.1.0
ENV APP_VERSION=$APP_VERSION
LABEL app-version=$APP_VERSION

WORKDIR /usr/bin/

COPY --from=BUILDER /beat .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/usr/bin/beat"]