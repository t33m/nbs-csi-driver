FROM golang:1.21 AS builder
ARG VERSION
WORKDIR /build
COPY . .
ENV CGO_ENABLED=0
RUN go build -ldflags "-X main.version=$VERSION" -o nbs-csi-driver cmd/nbs-csi-driver/main.go

FROM scratch
COPY --from=builder /build/nbs-csi-driver /nbs-csi-driver
ENTRYPOINT ["/nbs-csi-driver"]
