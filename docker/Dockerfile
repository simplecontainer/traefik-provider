FROM --platform=${BUILDPLATFORM:-linux/amd64} alpine:latest as builder

ARG BINARY=traefik-provider
ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

RUN adduser -D node && \
    mkdir -p /opt/${BINARY} /home/node/smr /home/node/.ssh && \
    chown -R node /home/node /opt/${BINARY} && \
    touch /home/node/.ssh/known_hosts

COPY ${BINARY}-${TARGETOS}-${TARGETARCH}/${BINARY} /opt/${BINARY}/${BINARY}
RUN chmod +x /opt/${BINARY}/${BINARY} && \
    # Install necessary packages
    apk add --no-cache busybox-extras ca-certificates tzdata

FROM alpine:latest

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /home/node /home/node
COPY --from=builder /opt/${BINARY} /opt/${BINARY}
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /home/node/smr

USER node

EXPOSE 7431

ENTRYPOINT ["/opt/traefik-provider/traefik-provider"]