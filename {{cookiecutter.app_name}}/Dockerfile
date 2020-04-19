FROM golang:1.13 as builder
COPY . /opt/base-api-fdn

WORKDIR /opt/base-api-fdn
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.9
RUN apk --no-cache add ca-certificates

WORKDIR /opt/base-api-fdn
COPY --from=builder /opt/base-api-fdn/ .

EXPOSE 8080
CMD ["./app"]
