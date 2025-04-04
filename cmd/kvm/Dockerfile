FROM golang:1.24.2-bookworm AS builder

WORKDIR /app
COPY . /app
RUN make build-kvm

FROM debian:bookworm-slim

LABEL org.opencontainers.image.source="https://github.com/wbpascal/kubevirt-device-plugins"
LABEL org.opencontainers.image.description="KVM Kubernetes Device Plugin"
LABEL org.opencontainers.image.licenses=MIT

COPY --from=builder /app/cmd/kvm/kvm /usr/local/bin/kvm

ENTRYPOINT ["/usr/local/bin/kvm"]
