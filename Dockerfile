# use golang:alpine to get the certificates and timezone data
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache ca-certificates tzdata && update-ca-certificates

# prduction app is from scratch. This minifies the attack vector.
FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Just add the binary and the config directory
ADD flamingo /flamingo
ADD config /config
# don't forget all frontend stuff if needed (non for graphql/api-only projects)
ADD templates /templates
ENTRYPOINT ["/flamingo"]
CMD ["serve"]
