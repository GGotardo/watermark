FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" cmd/watermark/watermark.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/watermark /usr/bin/

EXPOSE 8081
EXPOSE 8082

ENTRYPOINT ["watermark"]